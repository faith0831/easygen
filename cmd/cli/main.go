package main

import (
	"log"

	"github.com/faith0831/easygen/pkg/app"
)

func main() {
	if err := app.Gen(); err != nil {
		log.Fatal(err)
	}
}
