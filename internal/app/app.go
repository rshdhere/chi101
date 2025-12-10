// Package app deals with all the app logic
package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/rshdhere/chi101/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
	handler := api.NewWorkoutHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: handler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "status avilable\n")
}
