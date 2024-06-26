package api

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/database"
	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/scraper"
)

func (apiCfg *apiConfig) createDataHandler(w http.ResponseWriter, r *http.Request) {
	var data []scraper.Movie
	data = scraper.Scrape()

	for _, movie := range data {

		m, err := apiCfg.DB.CreateMovie(r.Context(), database.CreateMovieParams{
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
			respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating movie: %v", err))
			return
		}

		log.Printf("Created movie: %v\n", m)
	}

	respondWithJSON(w, http.StatusCreated, "Data created successfully")
}
