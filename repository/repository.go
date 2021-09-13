package repository

import (
	"context"
	"database/sql"
	"fmt"
	"movie-finder/models"
	"time"
)

const (
	GetMovieByID = `select id, title, description, year, release_date, runtime, rating, mpaa_rating, 
       				created_at, updated_at from movies where id = $1`
	GetMovieGenre = `select mg.id, mg.movie_id, mg.genre_id, mg.created_at, mg.updated_at, g.id, g.genre_name, g.created_at, 
       				g.updated_at from movies_genres mg inner join genres g on g.id = mg.genre_id 
       				where mg.movie_id = $1`
	GetAllMovies = `select id, title, description, year, release_date, runtime, rating, mpaa_rating, 
       				created_at, updated_at from movies %s order by title`
	GetAllGenres = `select id, genre_name, created_at, updated_at from genres order by genre_name`
	CreateMovie  = `insert into movies(title, description, year, release_date, runtime, rating, mpaa_rating, created_at, updated_at)
 					values($1, $2, $3, $4, $5, $6, $7, $8, $9) returning id`
	UpdateMovie = `update movies set title = $1, description = $2, year = $3, release_date = $4, runtime = $5, 
                  rating = $6, mpaa_rating = $7, updated_at = $8 where id = $9`
)

type Repo struct {
	DB *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		DB: db,
	}
}

// GetMovie will return single movie
func (r *Repo) GetMovie(id int) (*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.DB.Prepare(GetMovieByID)
	if err != nil {
		return nil, err
	}
	var movie models.Movie
	err = stmt.QueryRowContext(ctx, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Year,
		&movie.ReleaseDate,
		&movie.RunTime,
		&movie.Rating,
		&movie.MPAARating,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	stmt, err = r.DB.Prepare(GetMovieGenre)
	if err != nil {
		return nil, err
	}
	var genres []models.MovieGenre
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var genre models.MovieGenre
		rows.Scan(
			&genre.ID,
			&genre.MovieID,
			&genre.GenreID,
			&genre.CreatedAt,
			&genre.UpdatedAt,
			&genre.Genre.ID,
			&genre.Genre.GenreName,
			&genre.Genre.CreatedAt,
			&genre.Genre.UpdatedAt,
		)
		genres = append(genres, genre)
	}
	movie.MovieGenre = genres
	return &movie, nil
}

// GetAllMovie will return all movies
func (r *Repo) GetAllMovie(genreID ...int) ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	where := ""
	if len(genreID) > 0 {
		where = fmt.Sprintf("where id in (select movie_id from movies_genres where genre_id = %d)", genreID[0])
	}

	query := fmt.Sprintf(GetAllMovies, where)

	stmt, err := r.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []*models.Movie
	for rows.Next() {
		var movie models.Movie
		rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Year,
			&movie.ReleaseDate,
			&movie.RunTime,
			&movie.Rating,
			&movie.MPAARating,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		movies = append(movies, &movie)
	}
	return movies, nil
}

func (r *Repo) DeleteMovie(id int) error {
	return nil
}

func (r *Repo) CreateMovie(movie *models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.DB.Prepare(CreateMovie)
	if err != nil {
		return err
	}
	id := 0
	err = stmt.QueryRowContext(ctx,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.RunTime,
		movie.Rating,
		movie.MPAARating,
		movie.CreatedAt,
		movie.UpdatedAt,
	).Scan(&id)
	if err != nil {
		return err
	}
	movie.ID = id
	return nil
}

func (r *Repo) UpdateMovie(movie *models.Movie) error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.DB.Prepare(UpdateMovie)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx,
		movie.Title,
		movie.Description,
		movie.Year,
		movie.ReleaseDate,
		movie.RunTime,
		movie.Rating,
		movie.MPAARating,
		movie.UpdatedAt,
		movie.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) SearchMovie(searchTerm string) ([]*models.Movie, error) {
	return nil, nil
}

func (r Repo) GenresAll() ([]*models.Genre, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	stmt, err := r.DB.Prepare(GetAllGenres)
	if err != nil {
		return nil, err
	}
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var genres []*models.Genre
	for rows.Next() {
		var g models.Genre
		rows.Scan(
			&g.ID,
			&g.GenreName,
			&g.CreatedAt,
			&g.UpdatedAt,
		)
		genres = append(genres, &g)
	}
	return genres, nil
}
