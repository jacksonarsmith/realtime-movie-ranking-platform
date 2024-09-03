// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: movies.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const checkMovieExists = `-- name: CheckMovieExists :one
SELECT EXISTS (
    SELECT 1
    FROM movies
    WHERE title = $1
      AND release_year = $2
      AND duration = $3
      AND audience = $4
      AND image_src = $5
      AND image_alt = $6
      AND movie_url = $7
) AS exists
`

type CheckMovieExistsParams struct {
	Title       string `json:"title"`
	ReleaseYear int32  `json:"release_year"`
	Duration    int32  `json:"duration"`
	Audience    string `json:"audience"`
	ImageSrc    string `json:"image_src"`
	ImageAlt    string `json:"image_alt"`
	MovieUrl    string `json:"movie_url"`
}

func (q *Queries) CheckMovieExists(ctx context.Context, arg CheckMovieExistsParams) (bool, error) {
	row := q.queryRow(ctx, q.checkMovieExistsStmt, checkMovieExists,
		arg.Title,
		arg.ReleaseYear,
		arg.Duration,
		arg.Audience,
		arg.ImageSrc,
		arg.ImageAlt,
		arg.MovieUrl,
	)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createMovie = `-- name: CreateMovie :one
INSERT INTO movies (id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at
`

type CreateMovieParams struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Rank        int32     `json:"rank"`
	PeakRank    int32     `json:"peak_rank"`
	ReleaseYear int32     `json:"release_year"`
	Duration    int32     `json:"duration"`
	Audience    string    `json:"audience"`
	Rating      float64   `json:"rating"`
	Votes       int32     `json:"votes"`
	ImageSrc    string    `json:"image_src"`
	ImageAlt    string    `json:"image_alt"`
	MovieUrl    string    `json:"movie_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.queryRow(ctx, q.createMovieStmt, createMovie,
		arg.ID,
		arg.Title,
		arg.Rank,
		arg.PeakRank,
		arg.ReleaseYear,
		arg.Duration,
		arg.Audience,
		arg.Rating,
		arg.Votes,
		arg.ImageSrc,
		arg.ImageAlt,
		arg.MovieUrl,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Rank,
		&i.PeakRank,
		&i.ReleaseYear,
		&i.Duration,
		&i.Audience,
		&i.Rating,
		&i.Votes,
		&i.ImageSrc,
		&i.ImageAlt,
		&i.MovieUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getFeaturedMovies = `-- name: GetFeaturedMovies :many
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at FROM movies WHERE rank <= 10
`

func (q *Queries) GetFeaturedMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.query(ctx, q.getFeaturedMoviesStmt, getFeaturedMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Rank,
			&i.PeakRank,
			&i.ReleaseYear,
			&i.Duration,
			&i.Audience,
			&i.Rating,
			&i.Votes,
			&i.ImageSrc,
			&i.ImageAlt,
			&i.MovieUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMovieByFields = `-- name: GetMovieByFields :one
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at
  FROM movies
  WHERE title = $1
    AND release_year = $2
    AND duration = $3
    AND audience = $4
    AND image_src = $5
    AND image_alt = $6
    AND movie_url = $7
LIMIT 1
`

type GetMovieByFieldsParams struct {
	Title       string `json:"title"`
	ReleaseYear int32  `json:"release_year"`
	Duration    int32  `json:"duration"`
	Audience    string `json:"audience"`
	ImageSrc    string `json:"image_src"`
	ImageAlt    string `json:"image_alt"`
	MovieUrl    string `json:"movie_url"`
}

func (q *Queries) GetMovieByFields(ctx context.Context, arg GetMovieByFieldsParams) (Movie, error) {
	row := q.queryRow(ctx, q.getMovieByFieldsStmt, getMovieByFields,
		arg.Title,
		arg.ReleaseYear,
		arg.Duration,
		arg.Audience,
		arg.ImageSrc,
		arg.ImageAlt,
		arg.MovieUrl,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Rank,
		&i.PeakRank,
		&i.ReleaseYear,
		&i.Duration,
		&i.Audience,
		&i.Rating,
		&i.Votes,
		&i.ImageSrc,
		&i.ImageAlt,
		&i.MovieUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMovieById = `-- name: GetMovieById :one
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at FROM movies WHERE id = $1
`

func (q *Queries) GetMovieById(ctx context.Context, id uuid.UUID) (Movie, error) {
	row := q.queryRow(ctx, q.getMovieByIdStmt, getMovieById, id)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Rank,
		&i.PeakRank,
		&i.ReleaseYear,
		&i.Duration,
		&i.Audience,
		&i.Rating,
		&i.Votes,
		&i.ImageSrc,
		&i.ImageAlt,
		&i.MovieUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getMovies = `-- name: GetMovies :many
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at FROM movies
`

func (q *Queries) GetMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.query(ctx, q.getMoviesStmt, getMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Rank,
			&i.PeakRank,
			&i.ReleaseYear,
			&i.Duration,
			&i.Audience,
			&i.Rating,
			&i.Votes,
			&i.ImageSrc,
			&i.ImageAlt,
			&i.MovieUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getMoviesUpdatedMoreThanAnHourAgo = `-- name: GetMoviesUpdatedMoreThanAnHourAgo :many
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at FROM movies 
  WHERE updated_at <= NOW() - INTERVAL '1 hour'
`

func (q *Queries) GetMoviesUpdatedMoreThanAnHourAgo(ctx context.Context) ([]Movie, error) {
	rows, err := q.query(ctx, q.getMoviesUpdatedMoreThanAnHourAgoStmt, getMoviesUpdatedMoreThanAnHourAgo)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Rank,
			&i.PeakRank,
			&i.ReleaseYear,
			&i.Duration,
			&i.Audience,
			&i.Rating,
			&i.Votes,
			&i.ImageSrc,
			&i.ImageAlt,
			&i.MovieUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaginatedMoviesByRank = `-- name: GetPaginatedMoviesByRank :many
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at FROM movies ORDER BY rank ASC LIMIT $1 OFFSET $2
`

type GetPaginatedMoviesByRankParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPaginatedMoviesByRank(ctx context.Context, arg GetPaginatedMoviesByRankParams) ([]Movie, error) {
	rows, err := q.query(ctx, q.getPaginatedMoviesByRankStmt, getPaginatedMoviesByRank, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Rank,
			&i.PeakRank,
			&i.ReleaseYear,
			&i.Duration,
			&i.Audience,
			&i.Rating,
			&i.Votes,
			&i.ImageSrc,
			&i.ImageAlt,
			&i.MovieUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaginatedMoviesByReleaseYear = `-- name: GetPaginatedMoviesByReleaseYear :many
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at FROM movies ORDER BY release_year DESC LIMIT $1 OFFSET $2
`

type GetPaginatedMoviesByReleaseYearParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) GetPaginatedMoviesByReleaseYear(ctx context.Context, arg GetPaginatedMoviesByReleaseYearParams) ([]Movie, error) {
	rows, err := q.query(ctx, q.getPaginatedMoviesByReleaseYearStmt, getPaginatedMoviesByReleaseYear, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Rank,
			&i.PeakRank,
			&i.ReleaseYear,
			&i.Duration,
			&i.Audience,
			&i.Rating,
			&i.Votes,
			&i.ImageSrc,
			&i.ImageAlt,
			&i.MovieUrl,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :one
UPDATE movies SET rank = $2, peak_rank = $3, rating = $4, updated_at = $5 WHERE id = $1 RETURNING id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at
`

type UpdateMovieParams struct {
	ID        uuid.UUID `json:"id"`
	Rank      int32     `json:"rank"`
	PeakRank  int32     `json:"peak_rank"`
	Rating    float64   `json:"rating"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) (Movie, error) {
	row := q.queryRow(ctx, q.updateMovieStmt, updateMovie,
		arg.ID,
		arg.Rank,
		arg.PeakRank,
		arg.Rating,
		arg.UpdatedAt,
	)
	var i Movie
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Rank,
		&i.PeakRank,
		&i.ReleaseYear,
		&i.Duration,
		&i.Audience,
		&i.Rating,
		&i.Votes,
		&i.ImageSrc,
		&i.ImageAlt,
		&i.MovieUrl,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
