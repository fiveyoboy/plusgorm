package plusgorm

// TableFieldDesc table field description
type TableFieldDesc struct {
	Name          string // column name
	NameUnderline string // column underline name
	NameCamel     string // Column name camel
	Type          string // column go type
	IsPRI         bool   // primary key
	IsUNI         bool   // unique index
	Position      string // column position
	GormExtraTag  string
}

// TableDesc table description
type TableDesc struct {
	Name          string           // table name
	NameUnderline string           // column underline name
	NameCamel     string           // Column name camel
	Fields        []TableFieldDesc // table fields
}

// DBTableStructure table struct
type DBTableStructure struct {
	TableCatalog           string `gorm:"column:TABLE_CATALOG"`            // 数据表登记目录
	TableSchema            string `gorm:"column:TABLE_SCHEMA"`             // 数据库名称
	TabName                string `gorm:"column:TABLE_NAME"`               // 表名
	ColumnName             string `gorm:"column:COLUMN_NAME"`              // 列名称
	OrdinalPosition        string `gorm:"column:ORDINAL_POSITION"`         // 字段在表中的第几列
	ColumnDefault          string `gorm:"column:COLUMN_DEFAULT"`           // 列的默认值
	IsNullable             string `gorm:"column:IS_NULLABLE"`              // 是否为空  YES/NO
	DataType               string `gorm:"column:DATA_TYPE"`                // 数据类型
	CharacterMaximumLength string `gorm:"column:CHARACTER_MAXIMUM_LENGTH"` // 字符最大长度
	CharacterOctetLength   string `gorm:"column:CHARACTER_OCTET_LENGTH"`   // 字节长度
	NumericPrecision       string `gorm:"column:NUMERIC_PRECISION"`        // 数据精度
	NumericScale           string `gorm:"column:NUMERIC_SCALE"`            // 数据规模
	DatetimePrecision      string `gorm:"column:DATETIME_PRECISION"`       //
	CharacterSetName       string `gorm:"column:CHARACTER_SET_NAME"`       // 字符集名称
	CollationName          string `gorm:"column:COLLATION_NAME"`           // 字符集校验名称
	ColumnType             string `gorm:"column:COLUMN_TYPE"`              // 列类型
	ColumnKey              string `gorm:"column:COLUMN_KEY"`               // 主键  PRI
	Extra                  string `gorm:"column:EXTRA"`                    // 额外说明
	Privileges             string `gorm:"column:PRIVILEGES"`               // 字段操作权限
	ColumnComment          string `gorm:"column:COLUMN_COMMENT"`           // 字段备注
	GenerationExpression   string `gorm:"column:GENERATION_EXPRESSION"`
	SrsId                  string `gorm:"column:SRS_ID"` //
	Tag                    string `gorm:"-"`             //
	Type                   string `gorm:"-"`             // Go 数据类型
	ColumnNameCamel        string `gorm:"-"`             // 列名称 驼峰
	ExtraGormTag           string `gorm:"-"`             // 额外的标签，如主键，不能为空等

	//外键信息
	PositionInUniqueConstraint string `json:"POSITION_IN_UNIQUE_CONSTRAINT" gorm:"column:POSITION_IN_UNIQUE_CONSTRAINT"`
	ReferencedTableSchema      string `json:"REFERENCED_TABLE_SCHEMA" gorm:"column:REFERENCED_TABLE_SCHEMA"` //关联数据库
	ReferencedTableName        string `json:"REFERENCED_TABLE_NAME" gorm:"column:REFERENCED_TABLE_NAME"`     //关联表
	ReferencedColumnName       string `json:"REFERENCED_COLUMN_NAME" gorm:"column:REFERENCED_COLUMN_NAME"`   //关联的键
}

// TableToStructConfig table to struct config
type TableToStructConfig struct {
	Stdout         bool              // result write stdout,default
	FilePath       string            // result write to file
	MysqlGoTypeMap map[string]string // replace default MysqlToGoTypeMap
}

// DBConfig db connect config
type DBConfig struct {
	Name         string            // database name
	Host         string            // ip
	Port         uint32            // port
	User         string            // user name
	Password     string            // password
	Parameters   map[string]string // dsn params
	MaxOpenConns uint32            // max open connect
	MaxIdleConns uint32            // max  idle connect
	Debug        bool              // debug mod
}
