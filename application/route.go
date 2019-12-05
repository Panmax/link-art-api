package application

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/api"
	"link-art-api/application/middleware"
	"link-art-api/infrastructure/util/response"
	"log"
)

func Setup(engine *gin.Engine) {
	engine.Use(middleware.NewRequestIdMiddleware())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "404 你懂吗？", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.SuccessResponse("pong")
	})

	apiGroup := engine.Group("/api")

	authMiddleware, err := middleware.NewJWTMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authGroup := apiGroup.Group("/auth")
	{
		authGroup.POST("/register", api.Register)
		authGroup.POST("/login", authMiddleware.LoginHandler)
		authGroup.POST("/refresh_token", authMiddleware.RefreshHandler)

		authGroup.Use(authMiddleware.MiddlewareFunc())
		{
			authGroup.GET("/logout", api.Logout)
			authGroup.GET("/profile", api.Profile)
			authGroup.POST("/profile", api.UpdateProfile)
			authGroup.PUT("/avatar", api.UpdateAvatar)
		}
	}

	accountGroup := apiGroup.Group("/accounts")
	{
		accountGroup.Use(authMiddleware.MiddlewareFunc())
		{
			accountGroup.POST("/approval", api.SubmitApproval)
		}
	}

}
