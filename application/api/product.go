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
)

func CreateProduct(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	if !account.IsArtist() {
		utilGin.ErrorResponse(-1, "未通过艺术家认证")
		return
	}

	cmd, e := bind.Bind(&command.CreateProductCommand{}, c)
	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}
	productCommand := cmd.(*command.CreateProductCommand)
	_, err := json.Marshal(productCommand.DetailsPics)
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

}

func ListMyProduct(c *gin.Context) {

}

func ShelvesProduct(c *gin.Context) {

}

func TakeOffProduct(c *gin.Context) {

}
