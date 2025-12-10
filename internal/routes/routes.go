// Package routes deals with routing logic using chi router
package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/rshdhere/chi101/internal/app"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", app.HealthCheck)
	r.Get("/workout/{id}", app.WorkoutHandler.GetWorkoutByID)

	r.Post("/workout", app.WorkoutHandler.CreateWorkout)

	return r
}
