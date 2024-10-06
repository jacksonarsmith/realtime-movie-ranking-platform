package api

import (
	"log"
	"net/http"
)

func (apiCfg *apiConfig) getProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user email from context
	email, ok := r.Context().Value(ContextKeyEmail).(string)
	if !ok || email == "" {
		log.Println("Error: email not found in context or is not a string")
		respondWithError(w, http.StatusUnauthorized, "Unauthorized, email not found in context")
		return
	}

	// Get user from database
	user, err := apiCfg.DB.GetUserByEmail(r.Context(), email)
	if err != nil {
		log.Printf("Error fetching user for email %s: %v", email, err)
		respondWithError(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	// Convert database user to API user
	apiUser := databaseUserToAPIUser(user)

	// Respond with user
	respondWithJSON(w, http.StatusOK, apiUser)
}
