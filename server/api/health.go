package api

import (
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, "Server is up and running!")
}
