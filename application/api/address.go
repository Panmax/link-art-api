package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/response"
)

func AddressRouterRegister(group *gin.RouterGroup) {
	addressGroup := group.Group("/address")
	{
		addressGroup.GET("/regions", ListRegion)
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
