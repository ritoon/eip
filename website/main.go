package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/website/handler"
)

func main() {

	// create all connections
	dbConn := db.New()
	h := handler.New(dbConn)

	// create router for gin
	router := gin.Default()

	// router.Use(cors.Default())

	// Pages in HTML
	router.LoadHTMLGlob("templates/*")
	router.Static("/public", "./public")
	// router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.GET("login", h.PageLogin)
	router.GET("search", h.PageGameList)

	// run the server
	router.Run(":8880")
}
