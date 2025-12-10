package main

import (
	"flag"

	"github.com/rshdhere/chi101/internal/app"
)

func main() {
	var PORT int

	flag.IntVar(&PORT, "port", 8080, "port for a go-backend")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}
}
