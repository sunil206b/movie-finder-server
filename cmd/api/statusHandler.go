package main

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"movie-finder/models"
	"movie-finder/utilities"
	"net/http"
	"strconv"
	"time"
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
	//movieID, err := strconv.Atoi(chi.URLParam(r, "id"))
	//if err != nil {
	//	app.logger.Println(errors.New("invalid id parameter" + err.Error()))
	//	utilities.WriteError(w, http.StatusBadRequest, "Not a valid movie id")
	//	return
	//}
}

type MoviePayload struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Year        string `json:"year"`
	ReleaseDate string `json:"release_date"`
	RunTime     string `json:"run_time"`
	Rating      string `json:"rating"`
	MPAARating  string `json:"mpaa_rating"`
}

func (app application) updateMovie(w http.ResponseWriter, r *http.Request) {
	var payload MoviePayload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		app.logger.Println(err)
		utilities.WriteError(w, http.StatusBadRequest, "Movies not found for the given genre")
		return
	}

	var movie models.Movie
	movie.ID, _ = strconv.Atoi(payload.ID)
	movie.Title = payload.Title
	movie.Description = payload.Description
	movie.ReleaseDate, _ = time.Parse("2006-01-02", payload.ReleaseDate)
	movie.Year = movie.ReleaseDate.Year()
	movie.RunTime, _ = strconv.Atoi(payload.RunTime)
	movie.Rating, _ = strconv.Atoi(payload.Rating)
	movie.MPAARating = payload.MPAARating
	movie.UpdatedAt = time.Now()
	if movie.ID <= 0 {
		movie.CreatedAt = time.Now()
		err = app.repo.CreateMovie(&movie)
		if err != nil {
			app.logger.Println(err)
			utilities.WriteError(w, http.StatusInternalServerError, "Failed to create a movie")
			return
		}
	} else {
		err = app.repo.UpdateMovie(&movie)
		if err != nil {
			app.logger.Println(err)
			utilities.WriteError(w, http.StatusInternalServerError, "Failed to update a movie")
			return
		}
	}
	utilities.WriteJSON(w, http.StatusOK, movie, "movie")
}

func (app application) searchMovie(w http.ResponseWriter, r *http.Request) {

}
