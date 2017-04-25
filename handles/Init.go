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
		loginRoute := rootRouter.Group("")
		{
			loginRoute.GET("/login", nil)
			loginRoute.POST("/login", nil)
		}
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
		planRoute.GET("/info", nil)

		detailRoute := planRoute.Group("detail")
		{
			detailRoute.GET("", nil)
			detailRoute.POST("", nil)
			detailRoute.PUT("", nil)
		}

		planRoute.GET("/owner", nil)
		planRoute.GET("/contributer", nil)

		starRoute := planRoute.Group("star")
		{
			starRoute.GET("", nil)
			starRoute.POST("", nil)
			starRoute.PUT("", nil)
			starRoute.DELETE("", nil)
		}

		planRoute.GET("/forker", nil)

		voteRoute := planRoute.Group("vote")
		{
			voteRoute.GET("/up", nil)
			voteRoute.POST("/up", nil)
			voteRoute.DELETE("/up", nil)
			voteRoute.GET("/down", nil)
			voteRoute.POST("/down", nil)
			voteRoute.DELETE("/down", nil)
		}

		donateRoute := planRoute.Group("donate")
		{
			donateRoute.GET("", nil)
			donateRoute.POST("", nil)
			donateRoute.PUT("", nil)
			donateRoute.DELETE("", nil)
		}

		interestedRoute := planRoute.Group("interested")
		{
			interestedRoute.GET("", nil)
			interestedRoute.POST("", nil)
			interestedRoute.PUT("", nil)
			interestedRoute.DELETE("", nil)
		}

		suggestRoute := planRoute.Group("suggest")
		{
			suggestRoute.GET("", nil)
			suggestRoute.POST("", nil)
			suggestRoute.PUT("", nil)
			suggestRoute.DELETE("", nil)
		}

		planRoute.GET("/license", gin.HandlerFunc(func(c *gin.Context) {
			c.Writer.WriteString("this function not support now")
			c.Writer.Flush()
			return
		}))
	}

	userRoute := rootRouter.Group("people")
	userRoute.Use(gin.BasicAuth(gin.Accounts{
		"root": "123456",
	}))
	{
		userRoute.GET("/info", nil)

		detailRoute := userRoute.Group("detail")
		{
			detailRoute.GET("", nil)
			detailRoute.POST("", nil)
			detailRoute.PUT("", nil)
		}

		todosRoute := userRoute.Group("todos")
		{
			todosRoute.GET("", nil)
			todosRoute.POST("", nil)
			todosRoute.PUT("", nil)
			todosRoute.DELETE("", nil)
		}
	}

	groupRoute := rootRouter.Group("group")
	groupRoute.Use(gin.BasicAuth(gin.Accounts{
		"root": "123456",
	}))
	{
		groupRoute.GET("/info", nil)

		detailRoute := groupRoute.Group("detail")
		{
			detailRoute.GET("", nil)
			detailRoute.POST("", nil)
			detailRoute.PUT("", nil)
		}

		todosRoute := groupRoute.Group("todos")
		{
			todosRoute.GET("", nil)
			todosRoute.POST("", nil)
			todosRoute.PUT("", nil)
			todosRoute.DELETE("", nil)
		}
	}

	app.Run(":" + port)
}
