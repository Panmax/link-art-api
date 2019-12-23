package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/response"
)

func IndexRouterRegister(group *gin.RouterGroup) {
	indexGroup := group.Group("/index")
	{
		indexGroup.GET("/discovery/products", ListDiscoveryProduct)
		indexGroup.GET("/discovery/artists", ListDiscoveryArtist)

		indexGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			indexGroup.GET("/follow/products", ListFollowArtistProduct)
		}
	}
}

func ListDiscoveryProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	products, err := service.ListDiscoveryProduct()
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	utilGin.SuccessResponse(products)
}

func ListDiscoveryArtist(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	// 可登录也可以不登录
	var accountId uint
	claims, _ := middleware.JWTMiddleware.GetClaimsFromJWT(c)
	if claims != nil {
		id := claims[middleware.IdentityKey].(float64)
		accountId = uint(id)
	}

	artists, err := service.ListDiscoveryArtist(accountId)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	utilGin.SuccessResponse(artists)
}

func ListFollowArtistProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	products, err := service.ListFollowArtistProduct(account.ID)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	utilGin.SuccessResponse(products)
}
