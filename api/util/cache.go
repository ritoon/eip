package util

import (
	"bytes"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/cache"
)

type CacheContext struct {
	body *bytes.Buffer
	gin.ResponseWriter
}

func (c *CacheContext) Write(b []byte) (int, error) {
	c.body.Write(b)
	return c.ResponseWriter.Write(b)
}

func GetCache(cache *cache.Redis, duration time.Duration, keyName, queryName string, funcErr func(ctx *gin.Context, err error)) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		log.Println("GetCache before")
		query := ctx.Query(queryName)
		res, err := cache.Get(ctx, keyName+"-"+query)
		if err == nil {
			log.Println("GetCache response from cache")
			ctx.Writer.WriteHeader(http.StatusOK)
			ctx.Header("Content-Type", "application/json")
			ctx.Writer.Write(res)
			ctx.Abort()
			return
		}
		cacheContext := CacheContext{
			body:           bytes.NewBuffer([]byte{}),
			ResponseWriter: ctx.Writer,
		}

		ctx.Writer = &cacheContext

		ctx.Next()
		log.Println("GetCache after")

		if ctx.Writer.Status() != http.StatusOK {
			return
		}

		log.Println("GetCache set cache")

		err = cache.Set(ctx, keyName+"-"+query, cacheContext.body.Bytes(), duration)
		if err != nil {
			funcErr(ctx, err)
			return
		}
		// ctx.Writer = originWriter
	}
}
