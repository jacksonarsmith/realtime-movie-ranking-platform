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

	router := http.NewServeMux()

	router.HandleFunc("GET /health", healthCheckHandler)
	router.HandleFunc("GET /movies", getMoviesHandler)
	router.HandleFunc("POST /users", apiCfg.createUserHandler)

	router.Handle("/api/v1", http.StripPrefix("/api/v1", router))

	server := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
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
			log.Printf("Error checking if movie exists: %v", err)
			continue
		}

		if exists {
			log.Printf("Movie already exists: %v", movie)
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

	log.Printf("Starting server on port %s\n", portStr)
	server.ListenAndServe()
}
