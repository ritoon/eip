package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/api/model"
	"github.com/ritoon/eip/api/util"
)

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

	if u.Pass == nil || payload.Pass == nil || *u.Pass != *payload.Pass {
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
	// u.Pass = nil
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
