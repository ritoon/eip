package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ritoon/eip/api/handler"
	"github.com/ritoon/eip/api/util"
)

func main() {
	router := gin.Default()
	jwtValidation := util.ValidateJwt()
	account := gin.Accounts{"admin": "admin"}

	router.POST("login", handler.LoginUser)

	// Users
	router.POST("users", gin.BasicAuth(account), handler.CreateUser)
	router.GET("users/:uuid", jwtValidation, handler.GetUser)
	router.DELETE("users/:uuid", jwtValidation, handler.DeleteUser)

	// Games
	router.GET("games", jwtValidation, handler.SearchGames)
	router.POST("games", jwtValidation, handler.CreateGame)
	router.GET("games/:uuid", jwtValidation, handler.GetGame)
	router.DELETE("games/:uuid", jwtValidation, handler.DeleteGame)

	// Addresses
	router.POST("addresses", jwtValidation, handler.CreateAddress)
	router.GET("addresses/:uuid", jwtValidation, handler.GetAddress)
	router.DELETE("addresses/:uuid", jwtValidation, handler.DeleteAddress)
	router.Run(":8888")
}
