-- +goose Up

CREATE table movies {
    id UUID PRIMARY KEY,
    title TEXT NOT NULL,
    release_year INT NOT NULL,
    duration INT NOT NULL,
    audience TEXT NOT NULL, 
    rating FLOAT NOT NULL, 
    votes FLOAT NOT NULL,
    image_url TEXT NOT NULL, 
    image_alt TEXT NOT NULL, 
    movie_url TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
};

-- +goose Down

DROP TABLE movies;
