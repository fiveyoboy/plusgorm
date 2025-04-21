package plusgorm

import (
	"bytes"
	"fmt"
	"go/format"
	"html/template"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GormPlus gorm plus
type GormPlus struct {
	Config    *DBConfig
	Logger    Logger
	db        *gorm.DB
	tableDesc TableDesc
}

// NewPlusGorm new plus gorm
func NewPlusGorm(conf *DBConfig) *GormPlus {
	return &GormPlus{
		Config: conf,
		Logger: defLog,
	}
}

// TableToStruct mysql table covert to go struct
func (g *GormPlus) TableToStruct(table string, conf *TableToStructConfig) error {
	var err error
	if conf == nil {
		conf = &TableToStructConfig{}
	}
	err = g.connect()
	if err != nil {
		return err
	}
	defer func() {
		_ = g.close()
	}()
	err = g.loadTableStruct(table)
	if err != nil {
		return err
	}
	err = g.writeToStruct(conf)
	if err != nil {
		return err
	}
	return nil
}
func (g *GormPlus) writeToStruct(conf *TableToStructConfig) error {
	var err error
	var buffer = new(bytes.Buffer)
	var tpl = "table.tpl"
	t := template.Must(template.New(tpl).ParseFiles(tpl))
	err = t.Execute(buffer, g.tableDesc)
	if err != nil {
		return err
	}
	data, err := format.Source(buffer.Bytes())
	if err != nil {
		data = buffer.Bytes()
	}
	switch true {
	case conf.FilePath != "":
		err = writeToFile(conf.FilePath, data, false)
	default:
		// stdout
		g.Logger.Infof("%s", data)
	}
	return nil
}
func (g *GormPlus) loadTableStruct(table string) error {
	var err error
	var tables []DBTableStructure
	if table == "" {
		return ErrTableNameInvalid
	}
	var sqlTpl = "SELECT * FROM information_schema.COLUMNS where TABLE_SCHEMA = '%s'   and information_schema.COLUMNS.TABLE_NAME = '%s'"
	var sql = fmt.Sprintf(sqlTpl, g.Config.Name, table)
	err = g.db.Raw(sql).Scan(&tables).Error
	if err != nil {
		return err
	}
	for i, col := range tables {
		tables[i].ColumnNameCamel = StrToCamel(col.ColumnName)
		tables[i].ColumnComment = strings.ReplaceAll(tables[i].ColumnComment, "\n", "")
		tables[i].ColumnComment = strings.ReplaceAll(tables[i].ColumnComment, "\r", "")
		tables[i].ColumnComment = strings.Trim(tables[i].ColumnComment, "\r")
		tables[i].ColumnComment = strings.Trim(tables[i].ColumnComment, "\n")
		tables[i].Type = MysqlToGoTypeMap[col.DataType]
		var extraTags []string
		if col.Type == "string" {
			extraTags = append(extraTags, fmt.Sprintf("size:%s", col.CharacterMaximumLength))
		}
		//不能为空 添加标签
		if col.IsNullable == "NO" {
			extraTags = append(extraTags, fmt.Sprintf("%s", "not null"))
		}
		//添加 主键
		if col.ColumnKey == "PRI" {
			extraTags = append(extraTags, fmt.Sprintf("%s", "primaryKey"))
		}
		//添加 唯一索引
		if col.ColumnKey == "UNI" {
			extraTags = append(extraTags, fmt.Sprintf("%s", "unique"))
		}
		//自增长 标签
		if col.Extra == "auto_increment" {
			extraTags = append(extraTags, fmt.Sprintf("%s", "autoIncrement"))
		}
		//添加默认值 标签
		if col.ColumnDefault != "" {
			extraTags = append(extraTags, fmt.Sprintf("default:%s", col.ColumnDefault))
		}
		// comment
		if col.ColumnComment != "" {
			extraTags = append(extraTags, fmt.Sprintf("comment:%s", col.ColumnComment))
		}
		if len(extraTags) != 0 {
			tables[i].ExtraGormTag = strings.Join(extraTags, ";")
		}
	}

	g.dbTableStructureToTableDesc(tables)
	return nil
}
func (g *GormPlus) connect() error {
	if g.Config == nil {
		panic("config is nil")
	}
	var parameters = "parseTime=True&loc=Local&charset=utf8mb4"
	for name, value := range g.Config.Parameters {
		if parameters == "" {
			parameters = fmt.Sprintf("%s=%s", name, value)
		} else {
			parameters += fmt.Sprintf("&%s=%s", name, value)
		}
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", g.Config.User, g.Config.Password, g.Config.Host, g.Config.Port, g.Config.Name, parameters)

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         255,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{})

	if err != nil {
		return err
	}

	db = db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci")

	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	if g.Config.MaxOpenConns != 0 {
		sqlDB.SetMaxOpenConns(int(g.Config.MaxOpenConns))
	}
	if g.Config.MaxIdleConns != 0 {
		sqlDB.SetMaxIdleConns(int(g.Config.MaxIdleConns))
	}

	sqlDB.SetConnMaxIdleTime(2 * time.Minute)
	sqlDB.SetConnMaxLifetime(5 * time.Minute)

	if g.Config.Debug {
		db = db.Debug()
	}
	g.db = db
	return nil
}
func (g *GormPlus) close() error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}
	err = sqlDB.Close()
	return err
}

func (g *GormPlus) dbTableStructureToTableDesc(tables []DBTableStructure) {
	var tableDesc = TableDesc{}
	sort.Slice(tables, func(i, j int) bool {
		return cast.ToInt(tables[i].OrdinalPosition) < cast.ToInt(tables[j].OrdinalPosition)
	})
	var fields = make([]TableFieldDesc, 0, len(tables))
	for _, tb := range tables {
		if tableDesc.Name == "" {
			tableDesc.Name = tb.TabName
			tableDesc.NameCamel = StrToCamel(tb.TabName)
			tableDesc.NameUnderline = StrToUnderline(tb.TabName)
		}
		fields = append(fields, TableFieldDesc{
			Name:          tb.ColumnName,
			NameUnderline: StrToUnderline(tb.ColumnName),
			NameCamel:     StrToCamel(tb.ColumnName),
			Type:          tb.Type,
			IsPRI:         tb.ColumnKey == ColumnKeyPRI,
			IsUNI:         tb.ColumnKey == ColumnKeyUNI,
			Position:      tb.OrdinalPosition,
			GormExtraTag:  tb.ExtraGormTag,
		})
	}
	tableDesc.Fields = fields
	g.tableDesc = tableDesc
	return
}

// SetLogger set custom logger
func (g *GormPlus) SetLogger(logger Logger) *GormPlus {
	g.Logger = logger
	return g
}
