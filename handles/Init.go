package handles

import (
	"github.com/gin-gonic/gin"
)

const (
	forkcookie = "forkcookie"
)

func Init(port string) {
	app := gin.New()
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	rootRouter := app.Group("/")
	{
		loginRoute := rootRouter.Group("")
		{
			loginRoute.GET("/login", HandleLoginGet)
			loginRoute.POST("/login", HandleLoginPost)
			loginRoute.GET("/register", HandleRegisterGet)
			loginRoute.POST("/register", HandleRegisterPost)
		}
		rootRouter.POST("/logout", HandleLogoutGet)

		rootRouter.GET("/intro", HandleIntro)
		rootRouter.GET("/about", HandleAbout)
	}
	exploreRoute := rootRouter.Group("explore")
	{
		exploreRoute.GET("/plan", HandleExplorePlan)
		exploreRoute.GET("/active", HandleExploreActive)
	}

	planRoute := rootRouter.Group("todos/:id")
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

	userRoute := rootRouter.Group("people/:id")
	{
		userRoute.GET("", HandleUserGet)
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

	groupRoute := rootRouter.Group("group/:id")
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
