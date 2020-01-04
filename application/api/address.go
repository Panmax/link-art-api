package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/command"
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/bind"
	"link-art-api/infrastructure/util/response"
	"strconv"
)

func AddressRouterRegister(group *gin.RouterGroup) {
	addressGroup := group.Group("/address")
	{
		addressGroup.GET("/regions", ListRegion)
		addressGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			addressGroup.POST("", CreateAddress)
			addressGroup.DELETE("/:id", DeleteAddress)
			addressGroup.POST("/:id/default", SetDefaultAddress)
		}
	}

	addressesGroup := group.Group("/addresses")
	addressesGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
	{
		addressesGroup.GET("", ListAddress)
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

func DeleteAddress(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}
	address, err := service.GetAddress(uint(id))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}
	if address.AccountId != account.ID {
		utilGin.ErrorResponse(-1, "fuck u")
		return
	}

	err = service.DeleteAddress(uint(id))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func ListAddress(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	addressRepresentations, err := service.ListAddress(account.ID)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(addressRepresentations)
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

func SetDefaultAddress(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	err = service.SetDefaultAddress(account.ID, uint(id))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}
