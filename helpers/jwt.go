package helpers

import (
	"time"

	"go_practice/structs"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("my_secret_key")

func GenerateToken(username string) (string, error) {
	// Set expiration time to 30 minutes from now
	expirationTime := time.Now().Add(30 * time.Minute)

	// Create claims
	claims := &structs.Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
