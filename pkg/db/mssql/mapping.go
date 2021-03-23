package mssql

import (
	"github.com/faith0831/easygen/pkg/db"
)

var csharpMapping = map[string]string{
	"bigint":           "long",
	"binary":           "byte[]",
	"bit":              "bool",
	"char":             "char",
	"date":             "DateTime",
	"datetime":         "DateTime",
	"datetime2":        "DateTime",
	"datetimeoffset":   "DateTimeOffset",
	"decimal":          "decimal",
	"float":            "float",
	"image":            "byte[]",
	"int":              "int",
	"money":            "decimal",
	"nchar":            "char",
	"ntext":            "string",
	"numeric":          "decimal",
	"nvarchar":         "string",
	"real":             "float",
	"smalldatetime":    "DateTime",
	"smallint":         "short",
	"smallmoney":       "decimal",
	"sql_variant":      "byte[]",
	"sysname":          "string",
	"text":             "string",
	"time":             "TimeSpan",
	"timestamp":        "byte[]",
	"tinyint":          "byte",
	"uniqueidentifier": "Guid",
	"varbinary":        "byte[]",
	"varchar":          "string",
	"xml":              "string",
}

var golangMapping = map[string]string{
	"bigint":           "int",
	"binary":           "[]byte",
	"bit":              "bool",
	"char":             "string",
	"date":             "time.Time",
	"datetime":         "time.Time",
	"datetime2":        "time.Time",
	"datetimeoffset":   "time.Time",
	"decimal":          "float",
	"float":            "float",
	"image":            "[]byte",
	"int":              "int",
	"money":            "float",
	"nchar":            "string",
	"ntext":            "string",
	"numeric":          "float",
	"nvarchar":         "string",
	"real":             "float",
	"smalldatetime":    "time.Time",
	"smallint":         "int",
	"smallmoney":       "float",
	"sql_variant":      "float",
	"sysname":          "string",
	"text":             "string",
	"time":             "time.Time",
	"timestamp":        "int",
	"tinyint":          "int",
	"uniqueidentifier": "string",
	"varbinary":        "string",
	"varchar":          "string",
	"xml":              "string",
}

var golangNullMapping = map[string]string{
	"bigint":           "int",
	"binary":           "[]byte",
	"bit":              "bool",
	"char":             "string",
	"date":             "time.Time",
	"datetime":         "time.Time",
	"datetime2":        "time.Time",
	"datetimeoffset":   "time.Time",
	"decimal":          "float",
	"float":            "float",
	"image":            "[]byte",
	"int":              "int",
	"money":            "float",
	"nchar":            "string",
	"ntext":            "string",
	"numeric":          "float",
	"nvarchar":         "string",
	"real":             "float",
	"smalldatetime":    "time.Time",
	"smallint":         "int",
	"smallmoney":       "float",
	"sql_variant":      "float",
	"sysname":          "string",
	"text":             "string",
	"time":             "time.Time",
	"timestamp":        "int",
	"tinyint":          "int",
	"uniqueidentifier": "string",
	"varbinary":        "string",
	"varchar":          "string",
	"xml":              "string",
}

// TypeMapping MsSQL type mapping
var TypeMapping db.TypeMappingFunc = func(lang string, typ string, isNull bool) string {
	if lang == "csharp" {
		if t, ok := csharpMapping[typ]; ok {
			if isNull == false || t == "byte[]" || t == "string" {
				return t
			}

			return t + "?"
		}
	} else if lang == "golang" {
		if isNull == false {
			if t, ok := golangMapping[typ]; ok {
				return t
			}
		} else {
			if t, ok := golangNullMapping[typ]; ok {
				return t
			}
		}
	}

	return "unknown"
}
