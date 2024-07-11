package app

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/faith0831/easygen/pkg/builder"
	"github.com/faith0831/easygen/pkg/config"
	"github.com/jessevdk/go-flags"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// Run Run
func Run(assets *embed.FS) error {
	a := Application{
		assets: assets,
		b:      &builder.Builder{},
	}

	return a.run()
}

// Gen Gen
func Gen() error {
	a := Application{
		b: &builder.Builder{},
	}

	return a.gen()
}

// Application 应用结构体
type Application struct {
	assets *embed.FS
	b      *builder.Builder
}

// GenOptions 生成的参数
type GenOptions struct {
	Template string            `long:"template" description:"模板名称" required:"true"`
	Table    string            `long:"table" description:"表名" required:"true"`
	Env      map[string]string `long:"env" description:"自定义变量"`
}

func (app *Application) gen() error {
	var options GenOptions
	var parser = flags.NewParser(&options, flags.Default)
	if _, err := parser.ParseArgs(os.Args); err != nil {
		log.Fatal(err)
	}

	c, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = app.b.CreateProvider(c)
	if err != nil {
		log.Fatal(err)
	}

	req := builder.GenerateRequest{
		Table:    options.Table,
		Template: options.Template,
		ENV:      map[string]interface{}{},
	}

	if len(options.Env) > 0 {
		for k, v := range options.Env {
			req.ENV[k] = v
		}
	}

	s, err := app.b.Generate(&req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(s)

	return nil
}

func (app *Application) startup(ctx context.Context) {
}

func (app *Application) run() error {
	err := wails.Run(&options.App{
		Title:  "easygen ver2.0.1 - 简单易用的代码生成器",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: app.assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	return nil
}

// Ok Ok
func (app *Application) Ok(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": 1,
		"msg":  "ok",
		"data": data,
	}
}

// Error Error
func (app *Application) Error(msg string) map[string]interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg":  msg,
	}
}

// Custom Custom
func (app *Application) Custom(code int, msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
