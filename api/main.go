package main

import (
	"log"
	"time"

	"net/http"
	_ "net/http/pprof"

	"github.com/gin-gonic/gin"
	ginratelimit "github.com/ljahier/gin-ratelimit"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ritoon/eip/api/cache"
	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/api/docs"
	"github.com/ritoon/eip/api/handler"
	"github.com/ritoon/eip/api/util"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	jwtValidation := util.ValidateJwt()
	// account := gin.Accounts{"admin": "admin"}
	dbConn := db.New()
	cacheConn := cache.New("localhost:6379", "", 0)
	h := handler.New(cacheConn, dbConn)

	cache2minForGames := util.GetCache(cacheConn, 2*time.Second, "searchgames", "name", handler.RespErr)
	rateLimitByIPForLogin := ginratelimit.RateLimitByIP(ginratelimit.NewTokenBucket(5, 1*time.Minute))
	rateLimitByIP := ginratelimit.RateLimitByIP(ginratelimit.NewTokenBucket(1000, 1*time.Minute))
	rateLimitUser := util.RateLimitUser(ginratelimit.NewTokenBucket(1000, 1*time.Minute))

	router.POST("login", rateLimitByIPForLogin, h.LoginUser)

	var v1 = router.Group("/api/v1")
	v1.Use(rateLimitByIP, jwtValidation, rateLimitUser)
	{
		// Users
		v1.POST("users", h.CreateUser)
		v1.GET("users/:uuid", h.GetUser)
		v1.DELETE("users/:uuid", jwtValidation, h.DeleteUser)

		// Games
		v1.GET("games", jwtValidation, cache2minForGames, h.SearchGames)
		v1.POST("games", jwtValidation, h.CreateGame)
		v1.GET("games/:uuid", jwtValidation, h.GetGame)
		v1.DELETE("games/:uuid", jwtValidation, h.DeleteGame)
		v1.POST("games/:uuid/images", jwtValidation, h.AddImageToGame)

		// Addresses
		v1.POST("addresses", jwtValidation, h.CreateAddress)
		v1.GET("addresses/:uuid", jwtValidation, h.GetAddress)
		v1.DELETE("addresses/:uuid", jwtValidation, h.DeleteAddress)
	}

	router.LoadHTMLGlob("templates/*")

	router.Static("/public", "./public")
	// router.StaticFile("/favicon.ico", "./public/favicon.ico")
	var pages = router.Group("/pages")
	{
		pages.GET("login", h.PageLogin)
	}

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run(":8888")
}
