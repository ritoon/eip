package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func main() {
	router := gin.Default()
	router.GET("/search", func(ctx *gin.Context) {
		query := ctx.Query("q")
		log.Println("query:", query)
		// simulate long process

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		ctx.JSON(http.StatusOK, gin.H{"lat": 123, "lng": 456})
	})
	router.Run(":8080")
}
