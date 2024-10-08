package api

import (
	"time"

	"github.com/google/uuid"
	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/database"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
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

type LoginUser struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Token   string    `json:"token"`
	Message string    `json:"message"`
}

func databaseUserToAPIUser(user database.User) User {
	return User{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		PasswordHash: user.PasswordHash,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func databaseUserToLoginAPIUser(user database.User, token string, message string) LoginUser {
	return LoginUser{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Token:   token,
		Message: message,
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

func databaseMoviesToAPIMovies(movies []database.Movie) []Movie {
	apiMovies := make([]Movie, 0, len(movies))
	for _, movie := range movies {
		apiMovies = append(apiMovies, databaseMovieToAPIMovie(movie))
	}
	return apiMovies
}
