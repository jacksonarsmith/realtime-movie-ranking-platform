package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getMoviesHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing request body: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}
