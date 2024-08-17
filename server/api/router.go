package api

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/database"
	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/scraper"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type apiConfig struct {
	DB *database.Queries
}

func StartServer() {
	godotenv.Load()

	portStr := os.Getenv("PORT")

	if portStr == "" {
		log.Fatal("PORT is not found in environment variables")
	}

	dbUrl := os.Getenv("DB_URL")

	if dbUrl == "" {
		log.Fatal("DB_URL is not found in environment variables")
	}

	conn, err := sql.Open("postgres", dbUrl)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}

	queries := database.New(conn)

	apiCfg := &apiConfig{
		DB: queries,
	}

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
	})

	router := http.NewServeMux()

	// Define paths for API routes
	apiRouter := http.NewServeMux()
	apiRouter.HandleFunc("GET /health", healthCheckHandler)
	apiRouter.HandleFunc("GET /movies", apiCfg.getMoviesHandler)
	apiRouter.HandleFunc("GET /movies/{id}", apiCfg.getMovieHandler)
	apiRouter.HandleFunc("POST /users", apiCfg.createUserHandler)

	// Define paths for static routes
	router.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))

	server := &http.Server{
		Addr:    ":" + portStr,
		Handler: c.Handler(router),
	}

	data := scraper.Scrape()

	for _, movie := range data {

		// Check if the movie already exists in the database
		exists, err := apiCfg.DB.CheckMovieExists(context.Background(), database.CheckMovieExistsParams{
			Title:       movie.Title,
			ReleaseYear: movie.Year,
			Duration:    movie.Duration,
			Audience:    movie.Audience,
			ImageSrc:    movie.ImageSrc,
			ImageAlt:    movie.ImageAlt,
			MovieUrl:    movie.MovieUrl,
		})

		if err != nil {
			log.Printf("Error checking if movie exists: %v\n", err)
			continue
		}

		if exists {
			log.Printf("Movie already exists: %v\n", movie)

			m, err := apiCfg.DB.GetMovieByFields(context.Background(), database.GetMovieByFieldsParams{
				Title:       movie.Title,
				ReleaseYear: movie.Year,
				Duration:    movie.Duration,
				Audience:    movie.Audience,
				ImageSrc:    movie.ImageSrc,
				ImageAlt:    movie.ImageAlt,
				MovieUrl:    movie.MovieUrl,
			})

			if err != nil {
				log.Printf("Error getting movie: %v\n", err)
			} else {
				log.Printf("Retrieved movie: %s\n", m.ID.String())

				newPeak := m.PeakRank

				if m.PeakRank < movie.Rank {
					newPeak = movie.Rank
				}

				updatedMovie, err := apiCfg.DB.UpdateMovie(context.Background(), database.UpdateMovieParams{
					ID:        m.ID,
					Rank:      m.Rank,
					PeakRank:  newPeak,
					Rating:    m.Rating,
					UpdatedAt: time.Now().UTC(),
				})

				log.Printf("Updated movie: %v\n", updatedMovie)

				if err != nil {
					log.Printf("Error updating movie: %v\n", err)
				} else {
					log.Printf("Updated movie: %v\n", updatedMovie)
				}
			}

			continue
		} else {
			m, err := apiCfg.DB.CreateMovie(context.Background(), database.CreateMovieParams{
				ID:          uuid.New(),
				Title:       movie.Title,
				Rank:        movie.Rank,
				PeakRank:    movie.Rank,
				ReleaseYear: movie.Year,
				Duration:    movie.Duration,
				Audience:    movie.Audience,
				Rating:      movie.Rating,
				Votes:       0,
				ImageSrc:    movie.ImageSrc,
				ImageAlt:    movie.ImageAlt,
				MovieUrl:    movie.MovieUrl,
				CreatedAt:   time.Now().UTC(),
				UpdatedAt:   time.Now().UTC(),
			})

			if err != nil {
				log.Printf("Error creating movie: %v", err)
			} else {
				log.Printf("Created movie: %v", m)
			}
		}
	}

	shiftNonUpdatedMovies(apiCfg)

	log.Printf("Starting server on port %s\n", portStr)
	server.ListenAndServe()
}

func shiftNonUpdatedMovies(apiCfg *apiConfig) {
	movies, err := apiCfg.DB.GetMoviesUpdatedMoreThanAnHourAgo(context.Background())

	if err != nil {
		log.Printf("Error getting movies: %v\n", err)
		return
	}

	for index, movie := range movies {

		updatedMovie, err := apiCfg.DB.UpdateMovie(context.Background(), database.UpdateMovieParams{
			ID:        movie.ID,
			Rank:      101 + int32(index),
			PeakRank:  movie.PeakRank,
			Rating:    movie.Rating,
			UpdatedAt: time.Now().UTC(),
		})

		if err != nil {
			log.Printf("Error updating movie: %v\n", err)
		} else {
			log.Printf("Updated movie: %v\n", updatedMovie)
		}
	}
}
