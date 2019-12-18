package main

import (
	"log"

	"github.com/faith0831/easygen/pkg/app"
	"github.com/faith0831/easygen/pkg/builder"
	"github.com/faith0831/easygen/pkg/config"
)

func main() {
	c, err := config.GetConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	b, err := builder.Create(c)
	if err != nil {
		log.Fatal(err.Error())
	}

	a := app.Create(b)
	if err := a.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
