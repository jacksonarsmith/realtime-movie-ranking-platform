package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type ContextKey string

const ContextKeyEmail ContextKey = "email"

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func jwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			respondWithError(w, http.StatusInternalServerError, "JWT_SECRET is not found in environment variables")
			return
		}

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized, no token provided")
			return
		}

		// Strip "Bearer " prefix if present
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized, invalid token format")
			return
		}

		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != "HS256" {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			respondWithError(w, http.StatusUnauthorized, "Unauthorized, token not valid")
			return
		}

		ctx := context.WithValue(r.Context(), ContextKeyEmail, claims.Subject)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
