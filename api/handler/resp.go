package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/db"
)

func RespErr(ctx *gin.Context, err error) {
	log.Println(err)
	respErr(ctx, err)
}

func respErr(ctx *gin.Context, err error) {

	switch e := err.(type) {
	case *db.Error:
		switch e.Code {
		case int(db.ErrCodeNotFound):
			ctx.JSON(http.StatusNotFound, gin.H{"error": e.Message})
			return
		case int(db.ErrCodeInternal):
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Message})
			return
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e.Message})
			return
		}
	default:
		eOrigin := errors.Unwrap(e)
		if eOrigin == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": e})
			return
		}
		respErr(ctx, eOrigin)
	}
}
