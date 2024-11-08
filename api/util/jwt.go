package util

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// JWTResponse is a struct to define jwt token response with jwt field.
type JWTResponse struct {
	Jwt string `json:"jwt"`
}

// CustomClaims is a struct to define jwt token claims with UUIDUser, Email and AccessLevel fields.
type CustomClaims struct {
	// UUIDUser is a unique identifier for user
	UUIDUser string `json:"uuid_user"`
	// AccessLevel is a level of access for user
	AccessLevel string `json:"access_level"`
	// Email is an email of user
	Email string `json:"email"`
	jwt.StandardClaims
}

// Valid is a function to validate jwt token claims with CustomClaims
func (c CustomClaims) Valid() error {
	return nil
}

var hmacSampleSecret = []byte("secret")

// NewJWT is a function to create new jwt token with uuidUser and email
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

// ValidateJwt is a middleware to validate jwt token from Authorization header
func ValidateJwt() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the Authorization header from the request header
		authValue := ctx.GetHeader("Authorization")
		// Check if the Authorization header is empty or not contains Bearer
		if authValue == "" || !strings.Contains(authValue, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
			ctx.Abort()
			return
		}
		if len(authValue) > 1000 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is too long"})
			ctx.Abort()
			return
		}

		// Get the jwt token value from Authorization header
		jwtValue := strings.ReplaceAll(authValue, "Bearer ", "")
		// Parse the jwt token with CustomClaims
		token, err := jwt.ParseWithClaims(jwtValue, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("util: unexpected signing method: %v", token.Header["alg"])
			}
			return hmacSampleSecret, nil
		})
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			ctx.Abort()
			return
		}
		// Check if the jwt token is valid and set the uuid_user and email to the context
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			// Set the uuid_user and email to the context
			ctx.Set("uuid_user", claims.UUIDUser)
			ctx.Set("email", claims.Email)
		} else {
			// Return an error if the jwt token is invalid
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
