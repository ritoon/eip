package util

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
)

func TestNewJWT(t *testing.T) {
	uuidUser := "12345"
	email := "test@example.com"

	tokenString, err := NewJWT(uuidUser, email)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSampleSecret, nil
	})
	assert.NoError(t, err)
	assert.True(t, token.Valid)

	claims, ok := token.Claims.(*CustomClaims)
	assert.True(t, ok)
	assert.Equal(t, uuidUser, claims.UUIDUser)
	assert.Equal(t, email, claims.Email)
	assert.WithinDuration(t, time.Now().Add(1*time.Hour), time.Unix(claims.ExpiresAt, 0), 10*time.Second)
}
