package api

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/database"
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
	router.HandleFunc("POST /data", apiCfg.createDataHandler)

	router.Handle("/api/v1", http.StripPrefix("/api/v1", router))

	server := &http.Server{
		Addr:    ":" + portStr,
		Handler: router,
	}

	log.Printf("Starting server on port %s\n", portStr)
	server.ListenAndServe()
}
