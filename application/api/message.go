package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/response"
	"strconv"
)

func MessageRouterRegister(group *gin.RouterGroup) {
	messageGroup := group.Group("/message")
	{
		messageGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			messageGroup.DELETE("/:id", DeleteMessage)
			messageGroup.GET("/:id", GetMessage)
			messageGroup.POST("/:id/read", ReadMessage)
		}
	}

	messageListGroup := group.Group("/messages")
	{
		messageListGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			messageListGroup.GET("", ListMessage)
			messageListGroup.GET("/new", CheckNewMessage)
		}
	}
}

func ListMessage(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	messages, err := service.ListMessage(account.ID)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(messages)
}

func DeleteMessage(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	messageId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	r, err := service.DeleteMessage(uint(messageId))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(r)
}

func GetMessage(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	messageId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	message, err := service.GetMessage(uint(messageId))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(message)
}

func CheckNewMessage(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	r, err := service.CheckNewMessage(account.ID)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(r)
}

func ReadMessage(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	messageId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	r, err := service.ReadMessage(uint(messageId))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(r)
}
