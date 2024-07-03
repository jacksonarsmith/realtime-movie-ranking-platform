package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/jacksonarsmith/realtime-movie-ranking-platform/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func (apiCfg *apiConfig) getUserHandler(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	params := parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Error parsing request body")
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(params.Email) {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Invalid email: %v", params.Email))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(params.Password), bcrypt.DefaultCost)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error hashing password: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUser(r.Context(), database.GetUserParams{
		Email:        params.Email,
		PasswordHash: string(hashedPassword),
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToAPIUser(user))
}
