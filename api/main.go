package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/api/model"
	"github.com/ritoon/eip/api/util"
)

func main() {
	router := gin.Default()
	router.POST("login", LoginUser)
	router.POST("users", CreateUser)
	router.GET("users/:uuid", GetUser)
	router.DELETE("users/:uuid", DeleteUser)
	router.Run(":8888")
}

var dbConn = db.New()

func LoginUser(ctx *gin.Context) {
	var payload model.UserLogin
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := dbConn.GetUserByEmail(payload.Email)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	if u.Pass != payload.Pass {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}
	jwtValue, err := util.NewJWT(u.UUID, u.Email)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"jwt": jwtValue})
}

func CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = dbConn.CreateUser(&u)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, u)
}

func GetUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := dbConn.GetUser(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func DeleteUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := dbConn.DeleteUser(uuid)
	if err != nil {
		RespErr(ctx, err)
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

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
