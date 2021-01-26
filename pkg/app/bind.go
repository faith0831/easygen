package app

import (
	"errors"

	"github.com/faith0831/easygen/pkg/builder"
	"github.com/faith0831/easygen/pkg/config"
)

// GetConfig 取配置信息
func (app *Application) GetConfig() map[string]interface{} {
	c, err := config.LoadConfig()
	if err != nil {
		return app.Error(err.Error())
	}

	return app.Ok(c)
}

// HasProvider 是否已创建数据源
func (app *Application) HasProvider() map[string]interface{} {
	return app.Ok(app.b.HasProvider())
}

// CreateProvider 创建数据源
func (app *Application) CreateProvider(o *builder.Options) map[string]interface{} {
	err := app.b.CreateProvider(o)
	if err != nil {
		return app.Error(err.Error())
	}

	return app.Ok(nil)
}

// GetTables 取数据表列表
func (app *Application) GetTables() map[string]interface{} {
	tables, err := app.b.GetTables()
	if err != nil {
		if errors.Is(err, builder.ErrNotFoundProvider) {
			return app.Custom(400, err.Error(), nil)
		}

		return app.Error(err.Error())
	}

	return app.Ok(tables)
}

// GetTemplates 取模板列表
func (app *Application) GetTemplates() map[string]interface{} {
	templates, err := app.b.GetTemplates()
	if err != nil {
		return app.Error(err.Error())
	}

	return app.Ok(templates)
}

// Generate 生成代码
func (app *Application) Generate(r *builder.GenerateRequest) map[string]interface{} {
	code, err := app.b.Generate(r)
	if err != nil {
		if errors.Is(err, builder.ErrNotFoundProvider) {
			return app.Custom(400, err.Error(), nil)
		}

		return app.Error(err.Error())
	}

	return app.Ok(code)
}
