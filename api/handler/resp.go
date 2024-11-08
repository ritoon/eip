package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/db"
)

// RespErr respond error
func RespErr(ctx *gin.Context, err error) {
	log.Println(err)
	respErr(ctx, err)
}

// RespErrWithCode respond error with code status
func RespErrWithCode(ctx *gin.Context, code int, err error) {
	log.Println(err)
	ctx.JSON(code, gin.H{"error": err})
	ctx.Abort()
}

func respErr(ctx *gin.Context, err error) {
	// switch type of error and respond with appropriate status code
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
