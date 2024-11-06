package util

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWTResponse struct {
	Jwt string `json:"jwt"`
}

type CustomClaims struct {
	UUIDUser    string `json:"uuid_user"`
	AccessLevel string `json:"access_level"`
	Email       string `json:"email"`
	jwt.StandardClaims
}

func (c CustomClaims) Valid() error {
	return nil
}

var hmacSampleSecret = []byte("secret")

func NewJWT(uuidUser, email string) (string, error) {
	var c CustomClaims
	c.UUIDUser = uuidUser
	c.Email = email
	c.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSampleSecret)
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}

func ValidateJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authValue := ctx.GetHeader("Authorization")
		if authValue == "" || !strings.Contains(authValue, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			ctx.Abort()
			return
		}

		jwtValue := strings.ReplaceAll(authValue, "Bearer ", "")
		token, err := jwt.ParseWithClaims(jwtValue, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			ctx.Set("uuid_user", claims.UUIDUser)
			ctx.Set("email", claims.Email)
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
