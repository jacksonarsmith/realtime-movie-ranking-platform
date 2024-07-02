-- name: CreateMovie :one 
INSERT INTO movies (id, title, rank, peak_rank, release_year, duration, audience, rating, votes, image_src, image_alt, movie_url, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14) RETURNING *;

-- name: GetMovies :many
SELECT * FROM movies; 

-- name: GetMoviesUpdatedMoreThanAnHourAgo :many
SELECT * FROM movies 
  WHERE updated_at <= NOW() - INTERVAL '1 hour';

-- name: CheckMovieExists :one
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
) AS exists;

-- name: GetMovieById :one
SELECT *
  FROM movies
  WHERE title = $1
    AND release_year = $2
    AND duration = $3
    AND audience = $4
    AND image_src = $5
    AND image_alt = $6
    AND movie_url = $7
LIMIT 1;

-- name: UpdateMovie :one
UPDATE movies SET rank = $2, peak_rank = $3, rating = $4, updated_at = $5 WHERE id = $1 RETURNING *;