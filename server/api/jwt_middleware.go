package api

import (
	"context"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ContextKey string

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := godotenv.Load()
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Error loading .env file")
			return
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			respondWithError(w, http.StatusInternalServerError, "JWT_SECRET is not found in environment variables")
			return
		}

		cookie, err := r.Cookie("token")
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), ContextKey("email"), claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
