package main

import (
	"time"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	ginratelimit "github.com/ljahier/gin-ratelimit"

	"github.com/ritoon/eip/api/cache"
	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/api/docs"
	"github.com/ritoon/eip/api/handler"
	"github.com/ritoon/eip/api/util"
)

func main() {
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	jwtValidation := util.ValidateJwt()
	account := gin.Accounts{"admin": "admin"}
	dbConn := db.New()
	cacheConn := cache.New("localhost:6379", "", 0)
	h := handler.New(cacheConn, dbConn)

	cache2minForGames := util.GetCache(cacheConn, 2*time.Second, "searchgames", "name", handler.RespErr)
	rateLimit := ginratelimit.NewTokenBucket(5, 1*time.Minute)

	router.POST("login", ginratelimit.RateLimitByIP(rateLimit), h.LoginUser)
	// Users
	router.POST("users", gin.BasicAuth(account), h.CreateUser)
	router.GET("users/:uuid", jwtValidation, h.GetUser)
	router.DELETE("users/:uuid", jwtValidation, h.DeleteUser)

	// mimeImagesOk := util.AuthorizedMimeType([]string{"image/jpeg", "image/png"})

	// Games
	router.GET("games", jwtValidation, cache2minForGames, h.SearchGames)
	router.POST("games", jwtValidation, h.CreateGame)
	router.GET("games/:uuid", jwtValidation, h.GetGame)
	router.DELETE("games/:uuid", jwtValidation, h.DeleteGame)
	router.POST("games/:uuid/images", jwtValidation, h.AddImageToGame)

	// Addresses
	router.POST("addresses", jwtValidation, h.CreateAddress)
	router.GET("addresses/:uuid", jwtValidation, h.GetAddress)
	router.DELETE("addresses/:uuid", jwtValidation, h.DeleteAddress)

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8888")
}
