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
	"time"
)

func ProductRouterRegister(group *gin.RouterGroup) {
	group.GET("/categories", ListCategoryTree)

	productGroup := group.Group("/products")
	{
		productGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			productGroup.POST("", CreateProduct)
			productGroup.PUT("/:id", UpdateProduct)
			productGroup.GET("/:id", GetProduct)
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

func ExhibitionRouterRegister(group *gin.RouterGroup) {
	exhibitionGroup := group.Group("/exhibitions")
	{
		exhibitionGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			exhibitionGroup.POST("", SubmitExhibition)
			exhibitionGroup.GET("", ListExhibition)
			exhibitionGroup.GET("/:id", GetExhibitionInfo)
			exhibitionGroup.GET("/:id/products", ListExhibitionProduct)
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
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	cmd, err := bind.Bind(&command.SubmitAuctionCommand{}, c)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}
	submitCommand := cmd.(*command.SubmitAuctionCommand)
	if submitCommand.Type != model.AuctionLiveType && submitCommand.Type != model.AuctionTextType {
		utilGin.ParamErrorResponse("类型错误")
		return
	}

	err = service.SubmitAuction(account.ID, submitCommand)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func ListAuction(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	accountID, _ := strconv.ParseUint(c.Query("artist_id"), 10, 64)
	auctionType, _ := strconv.ParseUint(c.Query("type"), 10, 8)
	status, _ := strconv.ParseUint(c.Query("status"), 10, 8)
	action, _ := strconv.ParseInt(c.Query("action"), 10, 8)

	auctions, err := service.ListAuction(uint(accountID), model.AuctionType(auctionType), model.AuctionStatus(status), int8(action))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(auctions)
}

func SubmitExhibition(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	cmd, err := bind.Bind(&command.SubmitExhibitionCommand{}, c)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}
	submitCommand := cmd.(*command.SubmitExhibitionCommand)
	if submitCommand.StartTime > submitCommand.EndTime {
		utilGin.ParamErrorResponse("开始时间必须小于结束时间")
		return
	}
	if submitCommand.StartTime < time.Now().Unix() {
		utilGin.ParamErrorResponse("开始时间错误")
		return
	}
	if len(submitCommand.ProductIDs) <= 0 {
		utilGin.ParamErrorResponse("个展商品不能为空")
		return
	}

	err = service.SubmitExhibition(account.ID, submitCommand)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func ListExhibition(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	accountID, _ := strconv.ParseUint(c.Query("artist_id"), 10, 64)
	action, _ := strconv.ParseInt(c.Query("action"), 10, 8)

	exhibition, err := service.ListExhibition(uint(accountID), int8(action))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(exhibition)
}

func GetExhibitionInfo(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	exhibitionId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	exhibitionRepresentation, err := service.GetExhibition(uint(exhibitionId))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(exhibitionRepresentation)
}

func ListExhibitionProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	exhibitionId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	products, err := service.ListExhibitionProduct(uint(exhibitionId))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(products)
}
