//go:generate go run -tags generate gen.go
package main

import (
	"embed"
	"log"

	"github.com/faith0831/easygen/pkg/app"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	if err := app.Run(&assets); err != nil {
		log.Fatal(err)
	}
}
