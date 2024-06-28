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
	Title       string
	ReleaseYear int32
	Duration    int32
	Audience    string
	ImageSrc    string
	ImageAlt    string
	MovieUrl    string
}

func (q *Queries) CheckMovieExists(ctx context.Context, arg CheckMovieExistsParams) (bool, error) {
	row := q.db.QueryRowContext(ctx, checkMovieExists,
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
	ID          uuid.UUID
	Title       string
	Rank        int32
	PeakRank    int32
	ReleaseYear int32
	Duration    int32
	Audience    string
	Rating      float64
	Votes       int32
	ImageSrc    string
	ImageAlt    string
	MovieUrl    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, createMovie,
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

const getMovieById = `-- name: GetMovieById :one
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

type GetMovieByIdParams struct {
	Title       string
	ReleaseYear int32
	Duration    int32
	Audience    string
	ImageSrc    string
	ImageAlt    string
	MovieUrl    string
}

func (q *Queries) GetMovieById(ctx context.Context, arg GetMovieByIdParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, getMovieById,
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

const getMovies = `-- name: GetMovies :many
SELECT id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at FROM movies
`

func (q *Queries) GetMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.QueryContext(ctx, getMovies)
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
	ID        uuid.UUID
	Rank      int32
	PeakRank  int32
	Rating    float64
	UpdatedAt time.Time
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) (Movie, error) {
	row := q.db.QueryRowContext(ctx, updateMovie,
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
