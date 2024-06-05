package db

// Column 表列信息
type Column struct {
	Index        int    `db:"Index" json:"index"`               //序号
	Name         string `db:"Name" json:"name"`                 //字段名
	IsPK         bool   `db:"IsPK" json:"isPk"`                 //是否主键
	IsInc        bool   `db:"IsInc" json:"isInc"`               //是否自增
	DbType       string `db:"DbType" json:"dbType"`             //数据库数据类型
	LangDataType string `json:"langDataType"`                   //语言数据类型
	IsNull       bool   `db:"IsNull" json:"isNull"`             //是否可空
	DefaultValue string `db:"DefaultValue" json:"defaultValue"` //默认值
	Comment      string `db:"Comment" json:"comment"`           //注释
}
