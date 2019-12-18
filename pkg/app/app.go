package app

import (
	"easygen/pkg/builder"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

// Application 应用结构体
type Application struct {
	b *builder.Builder
}

// Create 创建应用实例
func Create(b *builder.Builder) *Application {
	return &Application{
		b: b,
	}
}

// Run Run
func (app *Application) Run() error {
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./ui", false)))
	r.GET("api/tables", app.GetTables)
	r.GET("api/templates", app.GetTemplates)
	r.POST("api/generate", app.Generate)

	return r.Run(":1234")
}
