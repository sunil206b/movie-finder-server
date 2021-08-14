package models

import "time"

type Movie struct {
	ID          int          `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Year        int          `json:"year"`
	ReleaseDate time.Time    `json:"release_date"`
	RunTime     int          `json:"run_time"`
	Rating      int          `json:"rating"`
	MPAARating  string       `json:"mpaa_rating"`
	CreatedAt   time.Time    `json:"-"`
	UpdatedAt   time.Time    `json:"-"`
	MovieGenre  []MovieGenre `json:"movie_genre"`
}

type Genre struct {
	ID        int       `json:"id"`
	GenreName string    `json:"genre_name"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type MovieGenre struct {
	ID        int       `json:"-"`
	MovieID   int       `json:"-"`
	GenreID   int       `json:"-"`
	Genre     Genre     `json:"genre"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
