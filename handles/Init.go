package handles

import (
	"github.com/gin-gonic/gin"
)

func Init(port string) {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	rootRouter := app.Group("/")
	{
		rootRouter.GET("/login", nil)
		rootRouter.POST("/login", nil)
		rootRouter.POST("/logout", nil)
		rootRouter.POST("/intro", nil)
		rootRouter.POST("/about", nil)
	}
	exploreRoute := rootRouter.Group("explore")
	{
		exploreRoute.GET("plan", nil)
		exploreRoute.GET("active", nil)
	}

	planRoute := rootRouter.Group("todos")
	planRoute.Use(gin.BasicAuth(gin.Accounts{
		"root": "123456",
	}))
	{
		planRoute.GET("/", nil)
	}

	userRoute := rootRouter.Group("people")
	userRoute.Use(gin.BasicAuth(gin.Accounts{
		"root": "123456",
	}))
	{
		userRoute.GET("/", nil)
	}

	app.Run(":" + port)
}
