package util

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var hmacSampleSecret = []byte("secret")

func NewJWT(uuidUser, email string) (string, error) {
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"uuid_user": uuidUser,
		"email":     email,
		"exp":       time.Now().Add(1 * time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
