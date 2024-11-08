package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	ginratelimit "github.com/ljahier/gin-ratelimit"
	"github.com/penglongli/gin-metrics/ginmetrics"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/ritoon/eip/api/cache/mokecache"
	"github.com/ritoon/eip/api/db"
	"github.com/ritoon/eip/api/docs"
	"github.com/ritoon/eip/api/handler"
	"github.com/ritoon/eip/api/util"
)

func main() {
	// pprof for debug and profiling
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	// create router for gin
	router := gin.Default()

	// grafana tracing
	m := ginmetrics.GetMonitor()
	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(router)

	// add cors definition
	// router.Use(cors.Default())

	allowedOrigins := []string{"http://localhost:8880", "http://localhost:8888"}

	router.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "HEAD", "UPDATE", "PATCH", "TRACE", "CONNECT"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept", "Accept-Encoding", "Accept-Language", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Authorization", "Baggage", "Content-Length", "Content-Type", "Origin", "Referer", "Sec-Ch-Ua", "Sec-Ch-Ua-Mobile", "Sec-Ch-Ua-Platform", "Sentry-Trace", "User-Agent", "X-CSRF-Token", "X-Forwarded-Host", "X-Forwarded-Port", "X-Forwarded-Proto", "X-Forwarded-Server", "X-Max", "X-Requested-With", "X-XSRF-Token", "XSRF-Token", "Application_kind", "Application_version", "Application_device_id"},
		AllowCredentials: true,
		MaxAge:           5 * time.Minute,
	}))

	// create all connections
	dbConn := db.New()
	cacheConn := mokecache.New()
	h := handler.New(dbConn)

	// create all middlewares
	account := gin.Accounts{"admin": "admin"}
	basicAuth := gin.BasicAuth(account)
	jwtValidation := util.ValidateJwt()
	cache2minForGames := util.GetCache(cacheConn, 2*time.Second, "searchgames", "name", handler.RespErr)
	rateLimitByIPForLogin := ginratelimit.RateLimitByIP(ginratelimit.NewTokenBucket(5, 1*time.Minute))
	rateLimitByIP := ginratelimit.RateLimitByIP(ginratelimit.NewTokenBucket(1000, 1*time.Minute))
	rateLimitUser := util.RateLimitUser(ginratelimit.NewTokenBucket(1000, 1*time.Minute))

	// add the login route
	router.POST("login", rateLimitByIPForLogin, h.LoginUser)
	router.POST("register", basicAuth, h.RegisterUser)

	// add the api routes
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

	// Swagger
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// run the server
	router.Run(":8888")
}
