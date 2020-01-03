package app

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/faith0831/easygen/pkg/builder"

	"github.com/zserge/lorca"
)

// Run Run
func Run() {
	a := Application{}
	a.run()
}

// Application 应用结构体
type Application struct {
	b  *builder.Builder
	ui lorca.UI
}

func (app *Application) run() {
	ui, err := lorca.New("", "", 900, 700)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	app.b = &builder.Builder{}
	app.ui = ui

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	ui.Bind("api_hasProvider", app.HasProvider)
	ui.Bind("api_createProvider", app.CreateProvider)
	ui.Bind("api_generate", app.Generate)
	ui.Bind("api_getTables", app.GetTables)
	ui.Bind("api_getTemplates", app.GetTemplates)

	go http.Serve(ln, http.FileServer(http.Dir("./ui")))
	ui.Load(fmt.Sprintf("http://%s", ln.Addr()))

	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}
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
