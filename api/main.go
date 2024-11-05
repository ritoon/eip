package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/db"
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
