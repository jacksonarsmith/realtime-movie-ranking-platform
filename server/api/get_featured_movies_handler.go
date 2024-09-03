package api

import (
	"fmt"
	"net/http"
)

func (apiCfg *apiConfig) getFeaturedMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies, err := apiCfg.DB.GetFeaturedMovies(r.Context())

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting movies: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseMoviesToAPIMovies(movies))
}
