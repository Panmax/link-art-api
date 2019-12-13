package application

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/api"
	"link-art-api/application/middleware"
	"link-art-api/infrastructure/util/response"
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

	authGroup := apiGroup.Group("/auth")
	{
		authGroup.POST("/register", api.Register)
		authGroup.POST("/login", middleware.JWTMiddleware.LoginHandler)
		authGroup.POST("/refresh_token", middleware.JWTMiddleware.RefreshHandler)

		authGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			authGroup.GET("/logout", api.Logout)
			authGroup.GET("/profile", api.Profile)
			authGroup.POST("/profile", api.UpdateProfile)
			authGroup.PUT("/avatar", api.UpdateAvatar)
		}
	}

	accountGroup := apiGroup.Group("/accounts")
	{
		accountGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			accountGroup.POST("/approval", api.SubmitApproval)
		}
	}

	productGroup := apiGroup.Group("/products")
	{
		productGroup.GET("/categories", api.ListCategoryTree)

		productGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			productGroup.POST("", api.CreateProduct)
			productGroup.PUT("/:id", api.UpdateProduct)
			productGroup.GET("", api.ListAccountProduct)
			productGroup.POST("/:id/shelves", api.ShelvesProduct)
			productGroup.POST("/:id/take-off", api.TakeOffProduct)
		}
	}

	auctionGroup := apiGroup.Group("/auctions")
	{
		auctionGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			auctionGroup.POST("", api.SubmitAuction)
			auctionGroup.GET("", api.ListAuction)
		}
	}

	commonGroup := apiGroup.Group("/common")
	{
		commonGroup.GET("/oss/token", api.GetOssToken)
	}
}
