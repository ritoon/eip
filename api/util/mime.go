package util

import (
	"fmt"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

func AuthorizedMimeType(authorizedMimeTypes []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		mtype, _ := mimetype.DetectReader(ctx.Request.Body)
		fmt.Println("mtype :", mtype.String(), mtype.Extension())
	}
}
