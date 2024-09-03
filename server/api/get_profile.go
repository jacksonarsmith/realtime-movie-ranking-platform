package api

import "net/http"

func (apiCfg *apiConfig) getProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Get user email from context
	email := r.Context().Value("email").(string)

	// Get user from database
	user, err := apiCfg.DB.GetUserByEmail(r.Context(), email)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	// Convert database user to API user
	apiUser := databaseUserToAPIUser(user)

	// Respond with user
	respondWithJSON(w, http.StatusOK, apiUser)
}
