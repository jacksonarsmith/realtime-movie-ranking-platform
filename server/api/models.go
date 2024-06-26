package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Movie struct {
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

func databaseUserToAPIUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func databaseMovieToAPIMovie(movie database.Movie) Movie {
	return Movie{
		ID:          movie.ID,
		Title:       movie.Title,
		Rank:        movie.Rank,
		PeakRank:    movie.PeakRank,
		ReleaseYear: movie.ReleaseYear,
		Duration:    movie.Duration,
		Audience:    movie.Audience,
		Rating:      movie.Rating,
		Votes:       movie.Votes,
		ImageSrc:    movie.ImageSrc,
		ImageAlt:    movie.ImageAlt,
		MovieUrl:    movie.MovieUrl,
		CreatedAt:   movie.CreatedAt,
		UpdatedAt:   movie.UpdatedAt,
	}
}
