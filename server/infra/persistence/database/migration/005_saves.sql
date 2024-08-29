-- +goose Up

CREATE TABLE saves (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    movie_id UUID NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_saves_user_id ON saves(user_id);
CREATE INDEX idx_saves_movie_id ON saves(movie_id);

-- +goose Down

DROP TABLE saves;