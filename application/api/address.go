package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/command"
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/bind"
	"link-art-api/infrastructure/util/response"
)

func AddressRouterRegister(group *gin.RouterGroup) {
	addressGroup := group.Group("/address")
	{
		addressGroup.GET("/regions", ListRegion)
		addressGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			addressGroup.POST("", CreateAddress)
		}
	}
}

func ListRegion(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	regions, err := service.ListRegion()
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(regions)
}

func CreateAddress(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	cmd, err := bind.Bind(&command.CreateAddressCommand{}, c)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}
	addressCommand := cmd.(*command.CreateAddressCommand)

	err = service.CreateAddress(account.ID, addressCommand)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}
