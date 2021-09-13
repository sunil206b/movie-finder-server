package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := chi.NewRouter()
	router.Use(app.enableCORS)

	router.Get("/api/status", app.statusHandler)
	router.Get("/api/genres", app.getAllGenres)
	router.Get("/api/genres/{id}", app.getAllMoviesByGenre)
	router.Get("/api/movies", app.getAllMovie)
	router.Get("/api/movies/{id}", app.getMovie)
	router.Post("/api/admin/editmovie", app.updateMovie)
	router.Delete("/api/admin", app.deleteMovie)

	return router
}
