package util

import (
	"bytes"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/cache"
)

// CacheContext is a struct to store cache context
// body is a buffer to store response body
// ResponseWriter is a gin.ResponseWriter from gin.Context
type CacheContext struct {
	body *bytes.Buffer
	gin.ResponseWriter
}

// Write is a function to write response body to buffer and gin.ResponseWriter
func (c *CacheContext) Write(b []byte) (int, error) {
	c.body.Write(b)
	return c.ResponseWriter.Write(b)
}

// GetCache is a function to get cache from redis and set cache to redis if not found.
func GetCache(cache *cache.Redis, duration time.Duration, keyName, queryName string, funcErr func(ctx *gin.Context, err error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the query from the request
		query := ctx.Query(queryName)
		// get the cache from redis with keyName and query
		res, err := cache.Get(ctx, keyName+"-"+query)
		// if a cache is found, return the cache
		if err == nil {
			// write the cache to the response writer
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Header("Content-Type", "application/json")
			ctx.Writer.Write(res)
			ctx.Abort()
			return
		}
		// if no cache is not found, create a new cache
		// create a new cache context
		cacheContext := CacheContext{
			body:           bytes.NewBuffer([]byte{}),
			ResponseWriter: ctx.Writer,
		}
		// set the original response writer to cacheContext
		ctx.Writer = &cacheContext

		// call the next middleware
		ctx.Next()

		// if the status is not OK, don't set the cache
		if ctx.Writer.Status() != http.StatusOK {
			return
		}

		// set the cache to redis with keyName and query
		err = cache.Set(ctx, keyName+"-"+query, cacheContext.body.Bytes(), duration)
		if err != nil {
			funcErr(ctx, err)
			return
		}
	}
}
