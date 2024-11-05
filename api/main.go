package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/api/model"
)

func main() {
	router := gin.Default()
	router.GET("users/:uuid", GetUser)
	router.Run(":8888")
}

func GetUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := db.GetUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func DeleteUser(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	err := db.DeleteUser(uuid)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, nil)
}

func CreateUser(ctx *gin.Context) {
	var u model.User
	err := ctx.Bind(&u)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.CreateUser(&u)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, u)
}
