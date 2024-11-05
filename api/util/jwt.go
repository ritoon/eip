package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

func ValidateJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authValue := ctx.GetHeader("Authorization")
		if authValue == "" || !strings.Contains(authValue, "Bearer") {
			ctx.JSON(401, gin.H{"error": "authorization header is required"})
			ctx.Abort()
			return
		}

		jwtValue := strings.ReplaceAll(authValue, "Bearer ", "")
		token, err := jwt.Parse(jwtValue, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return hmacSampleSecret, nil
		})
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("uuid_user", claims["uuid_user"])
			ctx.Set("email", claims["email"])
		} else {
			ctx.JSON(401, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
