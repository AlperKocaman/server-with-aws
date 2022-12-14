package server

import (
	"github.com/AlperKocaman/server-with-aws/cmd/config"
	"github.com/AlperKocaman/server-with-aws/core/app"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/secure"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func InitAndRunServer() error {
	engine := InitializeServer()

	err := engine.Run(getRunAddress())
	if err != nil {
		return err
	}

	return nil
}

func InitializeServer() *gin.Engine {

	engine := gin.Default()

	setupRouter(engine)

	return engine
}

func getRunAddress() string {
	return config.Get().Server.Host + ":" + config.Get().Server.Port
}

func setupRouter(router *gin.Engine) *gin.Engine {

	setupMiddlewares(router)

	// A basic testing endpoint for server status
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	picusGroup := router.Group("/picus")
	app.NewDefaultRouter().Register(picusGroup)

	return router
}

func setupMiddlewares(router *gin.Engine) *gin.Engine {
	router.Use(secure.New(secure.Config{
		STSSeconds:              315360000,
		STSIncludeSubdomains:    true,
		CustomFrameOptionsValue: "SAMEORIGIN",
		ContentTypeNosniff:      true,
		BrowserXssFilter:        true,
		IENoOpen:                true,
		ReferrerPolicy:          "strict-origin-when-cross-origin",
	}))

	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // cors.Default() allows all origins,I kept this for simplicity.
		AllowMethods:    []string{"GET", "POST"},
		MaxAge:          60 * time.Second,
	}))

	return router
}
