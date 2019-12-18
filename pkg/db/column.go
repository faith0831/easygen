package db

// Column 表列信息
type Column struct {
	Index        int    `db:"Index"`    //序号
	Name         string `db:"Name"`     //字段名
	IsPK         bool   `db:"IsPK"`     //是否主键
	IsInc        bool   `db:"IsInc"`    //是否自增
	DataType     string `db:"DataType"` //数据类型
	LangDataType string
	IsNull       bool   `db:"IsNull"`       //是否可空
	DefaultValue string `db:"DefaultValue"` //默认值
	Comment      string `db:"Comment"`      //注释
}
