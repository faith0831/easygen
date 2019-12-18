package main

import (
	"log"

	"easygen/pkg/app"
	"easygen/pkg/builder"
	"easygen/pkg/config"
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
