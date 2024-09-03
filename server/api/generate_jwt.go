package api

import (
	"log"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
)

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func generateJWT(email string) (string, error) {
	godotenv.Load()

	secret := os.Getenv("JWT_SECRET")

	if secret == "" {
		log.Fatal("JWT_SECRET is not found in environment variables")
	}

	claims := &Claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
