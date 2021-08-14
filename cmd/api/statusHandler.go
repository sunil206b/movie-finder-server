package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"movie-finder/utilities"
	"net/http"
	"strconv"
)

func (app *application) statusHandler(w http.ResponseWriter, r *http.Request) {
	currentStatus := AppStatus{
		Status:      "Available",
		Environment: app.config.env,
		Version:     version,
	}

	err := json.NewEncoder(w).Encode(&currentStatus)
	if err != nil {
		app.logger.Println(err)
		return
	}
	w.Header().Set("Content_type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (app *application) getMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter" + err.Error()))
		utilities.WriteError(w, http.StatusBadRequest, "Not a valid movie id")
		return
	}

	movie, err := app.repo.GetMovie(id)
	if err != nil {
		app.logger.Println(err)
		utilities.WriteError(w, http.StatusBadRequest, "Movie not found with given id")
		return
	}
	utilities.WriteJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovie(w http.ResponseWriter, r *http.Request) {
	movies, err := app.repo.GetAllMovie()
	if err != nil {
		app.logger.Println(err)
		utilities.WriteError(w, http.StatusBadRequest, "Movies not found")
		return
	}
	utilities.WriteJSON(w, http.StatusOK, movies, "movies")
}

func (app application) getAllGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := app.repo.GenresAll()
	if err != nil {
		app.logger.Println(err)
		utilities.WriteError(w, http.StatusBadRequest, "Genres not found")
		return
	}
	utilities.WriteJSON(w, http.StatusOK, genres, "genres")
}
func (app application) getAllMoviesByGenre(w http.ResponseWriter, r *http.Request) {
	genreID, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		app.logger.Println(errors.New("invalid id parameter" + err.Error()))
		utilities.WriteError(w, http.StatusBadRequest, "Not a valid genre id")
		return
	}
	movies, err := app.repo.GetAllMovie(genreID)
	if err != nil {
		app.logger.Println(err)
		utilities.WriteError(w, http.StatusBadRequest, "Movies not found for the given genre")
		return
	}
	utilities.WriteJSON(w, http.StatusOK, movies, "movies")
}

func (app application) deleteMovie(w http.ResponseWriter, r *http.Request) {

}

func (app application) createMovie(w http.ResponseWriter, r *http.Request) {

}

func (app application) updateMovie(w http.ResponseWriter, r *http.Request) {

}

func (app application) searchMovie(w http.ResponseWriter, r *http.Request) {

}
