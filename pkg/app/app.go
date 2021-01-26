package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/faith0831/easygen/pkg/builder"
	"github.com/faith0831/easygen/pkg/config"

	"github.com/jessevdk/go-flags"
	"github.com/zserge/lorca"
)

// Run Run
func Run() error {
	a := Application{
		b: &builder.Builder{},
	}

	if len(os.Args) > 0 {
		return a.gen()
	} else {
		return a.run()
	}
}

// Application 应用结构体
type Application struct {
	b  *builder.Builder
	ui lorca.UI
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

func (app *Application) run() error {
	args := []string{"--disable-features=TranslateUI"}
	ui, err := lorca.New("", "", 900, 700, args...)
	if err != nil {
		return err
	}
	defer ui.Close()

	app.ui = ui

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}
	defer ln.Close()

	ui.Bind("api_getConfig", app.GetConfig)
	ui.Bind("api_hasProvider", app.HasProvider)
	ui.Bind("api_createProvider", app.CreateProvider)
	ui.Bind("api_generate", app.Generate)
	ui.Bind("api_getTables", app.GetTables)
	ui.Bind("api_getTemplates", app.GetTemplates)

	go http.Serve(ln, http.FileServer(FS))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	select {
	case <-ch:
	case <-ui.Done():
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
