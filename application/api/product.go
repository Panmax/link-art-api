package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"link-art-api/application/command"
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/bind"
	"link-art-api/infrastructure/util/response"
	"strconv"
)

func ProductRouterRegister(group *gin.RouterGroup) {
	productGroup := group.Group("/products")
	{
		productGroup.GET("/categories", ListCategoryTree)

		productGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			productGroup.POST("", CreateProduct)
			productGroup.PUT("/:id", UpdateProduct)
			productGroup.GET("", ListAccountProduct)
			productGroup.POST("/:id/shelves", ShelvesProduct)
			productGroup.POST("/:id/take-off", TakeOffProduct)
		}
	}
}

func AuctionRouterRegister(group *gin.RouterGroup) {
	auctionGroup := group.Group("/auctions")
	{
		auctionGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			auctionGroup.POST("", SubmitAuction)
			auctionGroup.GET("", ListAuction)
		}
	}
}

func CreateProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	if !account.IsArtist() {
		utilGin.ErrorResponse(-1, "未通过艺术家认证")
		return
	}

	cmd, err := bind.Bind(&command.CreateProductCommand{}, c)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}
	productCommand := cmd.(*command.CreateProductCommand)
	_, err = json.Marshal(productCommand.DetailPics)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	err = service.CreateProduct(account.ID, productCommand)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func UpdateProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	productId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	cmd, err := bind.Bind(&command.CreateProductCommand{}, c)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}
	productCommand := cmd.(*command.CreateProductCommand)
	_, err = productCommand.GetDetailPicsJson()
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	err = service.UpdateProduct(uint(productId), &account.ID, productCommand)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func ListAccountProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	products, err := service.ListProductByAccount(account.ID)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(products)
}

func GetProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	product, err := service.GetProduct(uint(id))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}
	utilGin.SuccessResponse(product)
}

func ShelvesProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	err = service.ShelvesProduct(uint(id), &account.ID)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func TakeOffProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	err = service.TakeOffProduct(uint(id), &account.ID)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func ListCategoryTree(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	categories, err := service.ListAllCategory()
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(categories)
}

func SubmitAuction(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	// TODO parse SubmitAuctionCommand
	utilGin.SuccessResponse(true)
}

func ListAuction(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	// TODO return []AuctionRepresentation
	// TODO filter by type
	utilGin.SuccessResponse(true)
}
