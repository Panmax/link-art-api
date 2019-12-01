package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/infrastructure/util/bind"
	"link-art-api/infrastructure/util/response"
	"link-art-api/route/param_bind"
)

func AuthRouterRegister(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	auth.POST("/login", Login)
	auth.GET("/logout", Logout)
}

func AccountRouterRegister(router *gin.RouterGroup) {

}

func Login(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	s, e := bind.Bind(&param_bind.Login{}, c)

	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	login := s.(*param_bind.Login)
	utilGin.SuccessResponse(login.Phone + "-" + login.Password)
}

func Logout(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	utilGin.SuccessResponse(nil)
}
