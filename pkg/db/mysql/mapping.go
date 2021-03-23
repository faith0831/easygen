package mysql

import (
	"github.com/faith0831/easygen/pkg/db"
)

var csharpMapping = map[string]string{
	"tinyint":    "bool",
	"bit":        "bool",
	"smallint":   "short",
	"mediumint":  "int",
	"int":        "int",
	"integer":    "int",
	"bigint":     "long",
	"float":      "float",
	"double":     "double",
	"decimal":    "decimal",
	"date":       "DateTime",
	"time":       "TimeSpan",
	"year":       "TimeSpan",
	"datetime":   "DateTime",
	"timestamp":  "byte[]",
	"char":       "char",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}

var golangMapping = map[string]string{
	"tinyint":    "bool",
	"bit":        "bool",
	"smallint":   "int",
	"mediumint":  "int",
	"int":        "int",
	"integer":    "int",
	"bigint":     "int",
	"float":      "float",
	"double":     "float",
	"decimal":    "float",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}

var golangNullMapping = map[string]string{
	"tinyint":    "bool",
	"bit":        "bool",
	"smallint":   "int",
	"mediumint":  "int",
	"int":        "int",
	"integer":    "int",
	"bigint":     "int",
	"float":      "float",
	"double":     "float",
	"decimal":    "float",
	"date":       "string",
	"time":       "string",
	"year":       "string",
	"datetime":   "time.Time",
	"timestamp":  "time.Time",
	"char":       "string",
	"varchar":    "string",
	"tinyblob":   "string",
	"tinytext":   "string",
	"blob":       "string",
	"text":       "string",
	"mediumblob": "string",
	"mediumtext": "string",
	"longblob":   "string",
	"longtext":   "string",
}

// TypeMapping MySQL type mapping
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
