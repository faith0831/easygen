package db

// Table 表信息
type Table struct {
	Name         string    `db:"name" json:"name"`
	OriginalName string    `json:"originalName"`
	Comment      string    `db:"comment" json:"comment"` //注释
	Columns      []*Column `json:"columns"`
}
