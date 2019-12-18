package mysql

import (
	"fmt"
	"sort"

	"easygen/pkg/db"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

const (
	// TableNamesScript 取表名列表sql语句
	TableNamesScript = `
	SELECT
		TABLE_NAME
	FROM
		information_schema.TABLES
	WHERE TABLE_SCHEMA = DATABASE()
	`
	// TableColumnsScript 取表字段信息sql语句
	TableColumnsScript = `
	SELECT
		ORDINAL_POSITION AS 'Index',
		COLUMN_NAME AS 'Name',
		(CASE EXTRA WHEN 'auto_increment' THEN 1 ELSE 0 END) AS 'IsInc',
		(CASE COLUMN_KEY WHEN 'PRI' THEN 1 ELSE 0 END) AS 'IsPK',
		DATA_TYPE AS 'DataType',
		(CASE IS_NULLABLE WHEN 'YES' THEN 1 ELSE 0 END) AS 'IsNull', 
		IFNULL(COLUMN_DEFAULT, '') AS 'DefaultValue', 
		IFNULL(COLUMN_COMMENT, '') AS 'Comment'
	FROM
		(
	SELECT *
	FROM information_schema.COLUMNS
	WHERE table_schema = DATABASE() AND table_name = '%s') t
	`
)

// Provider MySQL数据源
type Provider struct {
	db *sqlx.DB
}

// New 创建数据源实例
func New(connStr string) (db.Provider, error) {
	db, err := sqlx.Open("mysql", connStr)
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