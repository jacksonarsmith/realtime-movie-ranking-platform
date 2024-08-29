-- +goose Up

CREATE TABLE likes (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    movie_id UUID NOT NULL REFERENCES movies(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_likes_user_id ON likes(user_id);
CREATE INDEX idx_likes_movie_id ON likes(movie_id);

-- +goose Down

DROP TABLE likes;