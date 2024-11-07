package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ginratelimit "github.com/ljahier/gin-ratelimit"
)

func RateLimitUser(bucket *ginratelimit.TokenBucket) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		uuidUser := ctx.GetString("uuid_user")
		if uuidUser == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}
		handlInternal := ginratelimit.RateLimitByUserId(bucket, uuidUser)
		handlInternal(ctx)
	}
}
