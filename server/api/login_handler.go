package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (apiCfg *apiConfig) loginHandler(w http.ResponseWriter, r *http.Request) {

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

	user, err := apiCfg.DB.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(params.Password))
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	token, err := generateJWT(user.Email)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error generating JWT: %v", err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HttpOnly: true,
	})

	respondWithJSON(w, http.StatusOK, databaseUserToLoginAPIUser(user, token, "Login successful"))
}
