package application

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/api"
	"link-art-api/application/middleware"
	"link-art-api/infrastructure/util/response"
)

func SetupRoute(engine *gin.Engine) {
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
	api.AuthRouterRegister(apiGroup)
	api.AccountRouterRegister(apiGroup)
	api.UserRouterRegister(apiGroup)

	api.ProductRouterRegister(apiGroup)
	api.AuctionRouterRegister(apiGroup)
	api.ExhibitionRouterRegister(apiGroup)

	api.IndexRouterRegister(apiGroup)

	api.CommonRouterRegister(apiGroup)

}
