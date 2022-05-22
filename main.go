package main

import (
	// "fmt"

	"fmt"
	"github.com/pilinux/gorest/config"
	"github.com/pilinux/gorest/core/xlogger"
	"github.com/pilinux/gorest/database"
	"github.com/pilinux/gorest/lib/middleware"
	"github.com/pilinux/gorest/routers"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

var configure = config.Config()

func main() {
	xlogger.DefaultTag = "Tootier"
	xlogger.LogLevel = xlogger.LEVEL_VERBOSE

	if configure.Database.RDBMS.Activate == "yes" {
		// Initialize RDBMS client
		if err := database.InitDB().Error; err != nil {
			fmt.Println(err)
			return
		}
	}

	if configure.Database.REDIS.Activate == "yes" {
		// Initialize REDIS client
		if _, err := database.InitRedis(); err != nil {
			fmt.Println(err)
			return
		}
	}

	if configure.Database.MongoDB.Activate == "yes" {
		// Initialize MONGO client
		if _, err := database.InitMongo(); err != nil {
			fmt.Println(err)
			return
		}
	}

	// JWT
	middleware.AccessKey = []byte(configure.Security.JWT.AccessKey)
	middleware.AccessKeyTTL = configure.Security.JWT.AccessKeyTTL
	middleware.RefreshKey = []byte(configure.Security.JWT.RefreshKey)
	middleware.RefreshKeyTTL = configure.Security.JWT.RefreshKeyTTL

	router, err := SetupRouter()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = router.Run(":" + configure.Server.ServerPort)
	if err != nil {
		fmt.Println(err)
		return
	}
}

// SetupRouter ...
func SetupRouter() (*gin.Engine, error) {
	if configure.Server.ServerEnv == "production" {
		gin.SetMode(gin.ReleaseMode) // Omit this line to enable debug mode
	}

	file, err := os.Create("./logs/start.txt")
	if err != nil {
		return nil, err
	}

	// If it is required to write the logs to the file and the console
	// at the same time
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	router := gin.Default()

	// Which proxy to trust
	if configure.Security.TrustedIP == "nil" {
		err := router.SetTrustedProxies(nil)
		if err != nil {
			return router, err
		}
	} else {
		if configure.Security.TrustedIP != "" {
			err := router.SetTrustedProxies([]string{configure.Security.TrustedIP})
			if err != nil {
				return router, err
			}
		}
	}

	router.Use(middleware.CORS())
	router.Use(middleware.SentryCapture(configure.Logger.SentryDsn))

	// API:v1.0
	v1 := router.Group("/api/v1/") // localhost:3000//api/v1/
	{
		// RDBMS
		if configure.Database.RDBMS.Activate == "yes" {
			// USER
			routers.AuthRouter(v1)
			routers.PreAuthRouter(v1)
			routers.SearchRoute(v1 , middleware.JWT())
			routers.UserRouter(v1 , middleware.JWT())
		}

		// REDIS Playground
		/*if configure.Database.REDIS.Activate == "yes" {
			rPlayground := v1.Group("playground")
			rPlayground.GET("/redis_read", controller.RedisRead)        // Non-protected
			rPlayground.POST("/redis_create", controller.RedisCreate)   // Non-protected
			rPlayground.DELETE("/redis_delete", controller.RedisDelete) // Non-protected
			rPlayground.GET("/redis_read_hash", controller.RedisReadHash)        // Non-protected
			rPlayground.POST("/redis_create_hash", controller.RedisCreateHash)   // Non-protected
			rPlayground.DELETE("/redis_delete_hash", controller.RedisDeleteHash) // Non-protected
		}

		// Mongo Playground
		if configure.Database.MongoDB.Activate == "yes" {
			rPlaygroundMongo := v1.Group("playground-mongo")
			rPlaygroundMongo.POST("/mongo_create_one", controller.MongoCreateOne)                 // Non-protected
			rPlaygroundMongo.GET("/mongo_get_all", controller.MongoGetAll)                        // Non-protected
			rPlaygroundMongo.GET("/mongo_get_by_id/:id", controller.MongoGetByID)                 // Non-protected
			rPlaygroundMongo.POST("/mongo_get_by_filter", controller.MongoGetByFilter)            // Non-protected
			rPlaygroundMongo.PUT("/mongo_update_by_id", controller.MongoUpdateByID)               // Non-protected
			rPlaygroundMongo.DELETE("/mongo_delete_field_by_id", controller.MongoDeleteFieldByID) // Non-protected
			rPlaygroundMongo.DELETE("/mongo_delete_doc_by_id/:id", controller.MongoDeleteByID)    // Non-protected
		}*/
	}

	return router, nil
}
