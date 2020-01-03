package mssql

import (
	"fmt"
	"sort"

	"github.com/faith0831/easygen/pkg/db"

	"github.com/jmoiron/sqlx"

	_ "github.com/denisenkom/go-mssqldb"
)

const (
	// ProviderName mssql
	ProviderName = "mssql"

	// TableNamesScript 取表名列表sql语句
	TableNamesScript = `
	SELECT name
	FROM sysobjects
	WHERE xtype='u'
	ORDER BY name
	`
	// TableColumnsScript 取表字段信息sql语句
	TableColumnsScript = `
	SELECT 
	a.colorder N'Index', 
	a.name N'Name', 
	(CASE WHEN COLUMNPROPERTY(a.id,a.name,'IsIdentity') = 1 THEN 1 ELSE 0 END) N'IsInc', 
	(CASE WHEN (
	SELECT COUNT(*)
	FROM sysobjects
	WHERE (name IN (
	SELECT name
	FROM sysindexes
	WHERE (id = a.id) AND (indid IN (
	SELECT indid
	FROM sysindexkeys
	WHERE (id = a.id) AND (colid IN (
	SELECT colid
	FROM syscolumns
	WHERE (id = a.id) AND (name = a.name))))))) AND (xtype = 'PK')) > 0 THEN 1 ELSE 0 END) N'IsPK',
		b.name N'DataType', 
		(CASE WHEN a.isnullable = 1 THEN 1 ELSE 0 END) N'IsNull', ISNULL(e.text,'') N'DefaultValue', ISNULL(g.[value],'') AS N'Comment'
	FROM syscolumns a
	LEFT JOIN systypes b ON a.xtype = b.xusertype
	INNER JOIN sysobjects d ON a.id = d.id AND d.xtype = 'U' AND d.name <> 'dtproperties'
	LEFT JOIN syscomments e ON a.cdefault = e.id
	LEFT JOIN sys.extended_properties g ON a.id = g.major_id AND a.colid = g.minor_id
	WHERE d.name = '%s'
	ORDER BY object_name(a.id),a.colorder
	`
)

// Provider SQLServer数据源
type Provider struct {
	db *sqlx.DB
}

// New 创建数据源实例
func New(connStr string) (db.Provider, error) {
	db, err := sqlx.Open("mssql", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	p := &Provider{
		db: db,
	}

	return p, nil
}

// GetTableNames 取表名列表
func (p *Provider) GetTableNames() ([]string, error) {
	var names []string
	err := p.db.Select(&names, TableNamesScript)
	if err != nil {
		return nil, err
	}

	sort.Strings(names)

	return names, nil
}

// GetTable 根据表名取表信息
func (p *Provider) GetTable(tableName string) (*db.Table, error) {
	var columns []*db.Column
	err := p.db.Select(&columns, fmt.Sprintf(TableColumnsScript, tableName))
	if err != nil {
		return nil, err
	}

	return &db.Table{Name: tableName, Columns: columns}, nil
}
