package builder

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"

	"github.com/iancoleman/strcase"

	"github.com/faith0831/easygen/pkg/config"
	"github.com/faith0831/easygen/pkg/db"
	"github.com/faith0831/easygen/pkg/db/mssql"
	"github.com/faith0831/easygen/pkg/db/mysql"
)

var (
	// ErrNotFoundProvider 没有找到数据源异常
	ErrNotFoundProvider = errors.New("not found provider")
)

// GenerateRequest 生成接收实体
type GenerateRequest struct {
	Lang     string                 `json:"lang"`
	Table    string                 `json:"table"`
	Template string                 `json:"template"`
	ENV      map[string]interface{} `json:"env"`
}

// ENV ENV
type ENV struct {
	Label string `json:"label"`
	Key   string `json:"key"`
	Value string `json:"value"`
}

// Node Node
type Node struct {
	Name     string  `json:"name"`
	Children []*Node `json:"children"`
	Lang     string  `json:"lang"`
	Template string  `json:"template"`
	ENV      []*ENV  `json:"env"`
}

// Builder 生成器
type Builder struct {
	driver   string
	provider db.Provider
	prefixes []string
	config   *config.Config
}

var funcMap = template.FuncMap{
	"lower":      strings.ToLower,
	"upper":      strings.ToUpper,
	"pascal":     pascal,
	"snake":      strcase.ToSnake,
	"camel":      strcase.ToCamel,
	"lowerCamel": strcase.ToLowerCamel,
	"kebab":      strcase.ToKebab,
}

// HasProvider HasProvider
func (b *Builder) HasProvider() (bool, *config.Config) {
	return b.provider != nil, b.config
}

// CreateProvider CreateProvider
func (b *Builder) CreateProvider(c *config.Config) error {
	if c.Driver == mysql.ProviderName {
		conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local", c.Username, c.Password, c.Host, c.Database)
		provider, err := mysql.New(conn)
		if err != nil {
			return fmt.Errorf("连接mysql %w", err)
		}

		b.provider = provider
	} else if c.Driver == mssql.ProviderName {
		conn := fmt.Sprintf("user id=%s;password=%s;server=%s;database=%s", c.Username, c.Password, c.Host, c.Database)
		provider, err := mssql.New(conn)
		if err != nil {
			return fmt.Errorf("连接mssql %w", err)
		}

		b.provider = provider
	} else {
		return fmt.Errorf("不支持数据库%s", c.Driver)
	}

	b.driver = c.Driver
	b.prefixes = strings.Split(c.Prefixes, ",")
	b.config = c

	return nil
}

// Generate 生成代码
func (b *Builder) Generate(r *GenerateRequest) (string, error) {
	if b.provider == nil {
		return "", ErrNotFoundProvider
	}

	s, err := os.ReadFile(r.Template)
	if err != nil {
		return "", err
	}

	if len(r.Lang) == 0 {
		r.Lang = getLang(string(s))
	}

	table, err := b.provider.GetTable(r.Table)
	if err != nil {
		return "", err
	}

	for _, c := range table.Columns {
		c.LangDataType = b.provider.GetMappingType(r.Lang, c.DbType, c.IsNull)
	}

	t, err := template.New("builder").Funcs(funcMap).Parse(string(s))
	if err != nil {
		return "", err
	}

	if r.ENV == nil {
		r.ENV = make(map[string]interface{})
	}

	if len(b.prefixes) > 0 {
		for _, prefix := range b.prefixes {
			table.Name = strings.Replace(table.Name, strings.Trim(prefix, " "), "", 1)
		}
	}

	r.ENV["Table"] = table

	buf := new(bytes.Buffer)
	err = t.Execute(buf, r.ENV)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// GetTemplates 取模板列表
func (b *Builder) GetTemplates() ([]*Node, error) {
	root := &Node{}
	walk("./conf/template", root)
	return root.Children, nil
}

// GetTables 取数据表列表
func (b *Builder) GetTables() ([]string, error) {
	if b.provider == nil {
		return nil, ErrNotFoundProvider
	}

	return b.provider.GetTableNames()
}

func walk(dir string, node *Node) {
	items, err := os.ReadDir(dir)
	if err != nil {
		return
	}

	if node.Children == nil {
		node.Children = []*Node{}
	}

	for _, item := range items {
		tName := strings.TrimSuffix(item.Name(), path.Ext(item.Name()))
		fName := filepath.Join(dir, item.Name())

		child := &Node{
			Name: tName,
		}

		node.Children = append(node.Children, child)

		if item.IsDir() {
			walk(fName, child)
		} else {
			buf, err := os.ReadFile(fName)
			if err == nil {
				content := string(buf)
				child.Template = fName
				child.Lang = getLang(content)
				child.ENV = getEnv(content)
			}
		}
	}
}

func getLang(content string) string {
	exp := regexp.MustCompile(`@lang\s+(\w+)`)
	m := exp.FindStringSubmatch(content)
	if len(m) > 1 {
		return m[1]
	}

	return ""
}

func getEnv(content string) []*ENV {
	var env []*ENV

	exp := regexp.MustCompile(`@env\s+(\w+)\s+([\p{Han}a-zA-Z0-9_-]+)`)
	items := exp.FindAllString(content, 100)

	if len(items) > 0 {
		env = make([]*ENV, len(items))

		for index, item := range items {
			m := exp.FindStringSubmatch(item)
			env[index] = &ENV{
				Key:   m[1],
				Label: m[2],
				Value: "",
			}
		}
	}

	return env
}
