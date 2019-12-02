package application

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/api"
	"link-art-api/application/middleware/requestid"
	"link-art-api/infrastructure/util/response"
)

func Setup(engine *gin.Engine) {
	engine.Use(requestid.SetUp())

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
}
