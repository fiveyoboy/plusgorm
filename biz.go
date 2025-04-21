package plusgorm

import (
	"errors"
)

var (
	ErrConnectErr       = errors.New("connect err")
	ErrTableNameInvalid = errors.New("table_name is must")
)

const (
	ColumnKeyNormal = ""    // normal
	ColumnKeyPRI    = "PRI" // primary key
	ColumnKeyMUL    = "MUL" // secondary index
	ColumnKeyUNI    = "UNI" // unique index
)

// MysqlToGoTypeMap map for converting mysql type to golang types
var MysqlToGoTypeMap = map[string]string{
	"int":                "int",
	"integer":            "int",
	"tinyint":            "int",
	"smallint":           "int",
	"mediumint":          "int",
	"bigint":             "int64",
	"int unsigned":       "int",
	"integer unsigned":   "int",
	"tinyint unsigned":   "int",
	"smallint unsigned":  "int",
	"mediumint unsigned": "int",
	"bigint unsigned":    "int64",
	"bit":                "int",
	"bool":               "bool",
	"enum":               "string",
	"set":                "string",
	"varchar":            "string",
	"char":               "string",
	"tinytext":           "string",
	"mediumtext":         "string",
	"text":               "string",
	"longtext":           "string",
	"blob":               "string",
	"tinyblob":           "string",
	"mediumblob":         "string",
	"longblob":           "string",
	"date":               "time.Time", // time.Time or string
	"datetime":           "time.Time", // time.Time or string
	"timestamp":          "time.Time", // time.Time or string
	"time":               "time.Time", // time.Time or string
	"float":              "float64",
	"double":             "float64",
	"decimal":            "float64",
	"binary":             "string",
	"varbinary":          "string",
}

var GoToProtoTypeMap = map[string]string{
	"uint":    "uint64",
	"uint8":   "uint64",
	"uint16":  "uint64",
	"uint32":  "uint32",
	"uint64":  "uint64",
	"int":     "int64",
	"int8":    "int64",
	"int16":   "int64",
	"int32":   "int32",
	"int64":   "int64",
	"float32": "float",
	"float64": "double",
	"string":  "string",
	"bool":    "bool",
}

type DatabaseConfig struct {
	Name         string            // 数据库名
	Driver       string            // 驱动，未启用，默认mysql
	Host         string            // ip
	Port         uint32            // 端口
	User         string            // 用户名
	Password     string            // 密码
	Parameters   map[string]string // dsn参数
	MaxOpenConns uint32            // 最大连接数
	MaxIdleConns uint32            // 最大空闲数
	Debug        bool              // 是否打印sql语句
}
