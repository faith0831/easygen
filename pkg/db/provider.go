package db

// Provider 数据源接口
type Provider interface {
	GetTableNames() ([]string, error)
	GetTable(tableName string) (*Table, error)
	GetMappingType(lang string, typ string, isNull bool) string
}
