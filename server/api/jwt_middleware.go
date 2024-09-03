package api

import (
	"context"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		godotenv.Load()
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

		claims := &Claims{}

		token, err := jwt.ParseWithClaims(cookie.Value, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		ctx := context.WithValue(r.Context(), "email", claims.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
