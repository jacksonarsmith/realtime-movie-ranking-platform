-- +goose Up

CREATE TABLE movies (
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    rank INT NOT NULL,
    peak_rank INT NOT NULL, 
    release_year INT NOT NULL,
    duration INT NOT NULL,
    audience TEXT NOT NULL, 
    rating FLOAT NOT NULL, 
    votes INT NOT NULL,
    image_src TEXT NOT NULL, 
    image_alt TEXT NOT NULL, 
    movie_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- +goose Down

DROP TABLE movies;
