package api

import (
	"net/http"
)

func errorHandler(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, http.StatusNotFound, "Something went wrong. Please try again later.")
}
