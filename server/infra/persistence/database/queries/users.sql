-- name: CreateUser :one
INSERT INTO users (id, name, email, password_hash, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users WHERE email = $1;