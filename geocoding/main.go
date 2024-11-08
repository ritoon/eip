package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/exp/rand"
)

func main() {
	// create a new gin router
	router := gin.Default()

	// create a new route with GET method
	router.GET("/search", func(ctx *gin.Context) {

		// get query parameter from the request url
		query := ctx.Query("q")
		log.Println("query:", query)

		// simulate long process
		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
		ctx.JSON(http.StatusOK, gin.H{"lat": 123, "lng": 456})
	})

	// run the server on port 8080
	router.Run(":8080")
}
