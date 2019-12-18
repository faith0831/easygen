package builder

import (
	"bytes"
	"fmt"
	"html"
	"io/ioutil"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"easygen/pkg/config"
	"easygen/pkg/db"
	"easygen/pkg/db/mssql"
	"easygen/pkg/db/mysql"
)

// GenerateRequest 生成接收实体
type GenerateRequest struct {
	Lang     string                 `json:"lang"`
	Table    string                 `json:"table"`
	Template string                 `json:"template"`
	ENV      map[string]interface{} `json:"env"`
}

// Node Node
type Node struct {
	Label    string  `json:"label"`
	Value    string  `json:"value"`
	Children []*Node `json:"children"`
}

// Builder 生成器
type Builder struct {
	config   *config.Conf
	provider db.Provider
	mapping  db.TypeMappingFunc
}

// Create 创建生成器
func Create(c *config.Conf) (*Builder, error) {
	var p db.Provider
	var m db.TypeMappingFunc

	if c.Driver == "mysql" {
		conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", c.Username, c.Password, c.Host, c.Database)
		p1, err := mysql.New(conn)
		if err != nil {
			return nil, fmt.Errorf("连接mysql %w", err)
		}

		p = p1
		m = mysql.TypeMapping
	} else if c.Driver == "mssql" {
		conn := fmt.Sprintf("user id=%s;password=%s;server=%s;database=%s", c.Username, c.Password, c.Host, c.Database)
		p1, err := mssql.New(conn)
		if err != nil {
			return nil, fmt.Errorf("连接mssql %w", err)
		}

		p = p1
		m = mssql.TypeMapping
	} else {
		return nil, fmt.Errorf("不支持数据库%s", c.Driver)
	}

	b := &Builder{
		config:   c,
		provider: p,
		mapping:  m,
	}

	return b, nil
}

var funcMap = template.FuncMap{
	"lower": strings.ToLower,
}

// Generate 生成代码
func (b *Builder) Generate(r *GenerateRequest) (string, error) {
	s, err := ioutil.ReadFile(fmt.Sprintf("./tpl/%s/%s.tpl", r.Lang, r.Template))
	if err != nil {
		return "", err
	}

	table, err := b.provider.GetTable(r.Table)
	if err != nil {
		return "", err
	}

	for _, c := range table.Columns {
		c.LangDataType = b.mapping(r.Lang, c.DataType, c.IsNull)
	}

	t, err := template.New("builder").Funcs(funcMap).Parse(string(s))
	if err != nil {
		return "", err
	}

	if r.ENV == nil {
		r.ENV = make(map[string]interface{})
	}

	r.ENV["Table"] = table

	buf := new(bytes.Buffer)
	err = t.Execute(buf, r.ENV)
	if err != nil {
		return "", err
	}

	return html.EscapeString(buf.String()), nil
}

// GetTemplates 取模板列表
func (b *Builder) GetTemplates() ([]*Node, error) {
	basePath := "./tpl"
	items, err := ioutil.ReadDir(basePath)
	if err != nil {
		return nil, err
	}

	nodes := []*Node{}

	for _, item := range items {
		name := strings.TrimSuffix(item.Name(), path.Ext(item.Name()))
		n := Node{
			Label: name,
			Value: name,
		}

		if item.IsDir() {
			n.Children = []*Node{}
			items, err := ioutil.ReadDir(filepath.Join(basePath, item.Name()))
			if err != nil {
				continue
			}

			for _, item := range items {
				if item.IsDir() {
					continue
				}

				name := strings.TrimSuffix(item.Name(), path.Ext(item.Name()))
				n.Children = append(n.Children, &Node{
					Label: name,
					Value: name,
				})
			}
		}

		nodes = append(nodes, &n)
	}

	return nodes, nil
}

// GetTables 取数据表列表
func (b *Builder) GetTables() ([]string, error) {
	return b.provider.GetTableNames()
}
