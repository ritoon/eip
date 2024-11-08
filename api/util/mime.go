package util

import (
	"fmt"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

// AuthorizedMimeType check if the mimetype is authorized for the request body
func AuthorizedMimeType(authorizedMimeTypes []string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get the mimetype from the request body
		mtype, _ := mimetype.DetectReader(ctx.Request.Body)
		fmt.Println("mtype :", mtype.String(), mtype.Extension())
	}
}
