package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (apiCfg *apiConfig) getMovieHandler(w http.ResponseWriter, r *http.Request) {

	id := r.PathValue("id")

	newId, err := uuid.Parse(id)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid movie ID: %v", id))
		return
	}

	movie, err := apiCfg.DB.GetMovieById(r.Context(), newId)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting movie: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseMovieToAPIMovie(movie))
}
