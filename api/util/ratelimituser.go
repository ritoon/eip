package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginratelimit "github.com/ljahier/gin-ratelimit"
)

// RateLimitUser rate limit by user id
func RateLimitUser(bucket *ginratelimit.TokenBucket) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get user id from context
		uuidUser := ctx.GetString("uuid_user")
		// if user id is empty, return unauthorized
		if uuidUser == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		// rate limit by user id
		handlInternal := ginratelimit.RateLimitByUserId(bucket, uuidUser)
		handlInternal(ctx)
	}
}
