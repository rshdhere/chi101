package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/rshdhere/chi101/internal/app"
	"github.com/rshdhere/chi101/internal/routes"
)

func main() {
	var PORT int

	flag.IntVar(&PORT, "port", 8080, "port for a go-backend")
	flag.Parse()

	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	r := routes.SetupRoutes(app)

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", PORT),
		Handler:      r,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("your server is listening on http://localhost:%d", PORT)

	err = server.ListenAndServe()
	if err != nil {
		app.Logger.Fatal(err)
	}
}
