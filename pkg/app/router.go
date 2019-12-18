package app

import (
	"easygen/pkg/builder"

	"github.com/gin-gonic/gin"
)

// GetTables 取数据表列表
func (app *Application) GetTables(c *gin.Context) {
	ctx := &GIN{Context: c}

	tables, err := app.b.GetTables()
	if err != nil {
		ctx.Error(err.Error())
		return
	}

	ctx.Ok(tables)
}

// GetTemplates 取模板列表
func (app *Application) GetTemplates(c *gin.Context) {
	ctx := &GIN{Context: c}

	templates, err := app.b.GetTemplates()
	if err != nil {
		ctx.Error(err.Error())
		return
	}

	ctx.Ok(templates)
}

// Generate 生成代码
func (app *Application) Generate(c *gin.Context) {
	ctx := &GIN{Context: c}
	r := &builder.GenerateRequest{}

	if err := ctx.ShouldBindJSON(r); err != nil {
		ctx.Error("参数无效")
		return
	}

	code, err := app.b.Generate(r)
	if err != nil {
		ctx.Error(err.Error())
		return
	}

	ctx.Ok(code)
}
