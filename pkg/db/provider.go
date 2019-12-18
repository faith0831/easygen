package db

// Provider 数据源接口
type Provider interface {
	GetTableNames() ([]string, error)
	GetTable(tableName string) (*Table, error)
}

// TypeMappingFunc 数据type转换函数
type TypeMappingFunc func(lang string, typ string, isNull bool) string
