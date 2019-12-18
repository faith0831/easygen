package db

// Table 表信息
type Table struct {
	Name    string
	Columns []*Column
}
