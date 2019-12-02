package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/param_bind"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/bind"
	"link-art-api/infrastructure/util/response"
)

func AuthRouterRegister(router *gin.RouterGroup) {
	auth := router.Group("/auth")
	auth.POST("/register", Register)
	auth.POST("/login", Login)
	auth.GET("/logout", Logout)
}

func AccountRouterRegister(router *gin.RouterGroup) {

}

func Register(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	s, e := bind.Bind(&param_bind.Register{}, c)

	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	register := s.(*param_bind.Register)

	if register.Sms != "999999" { // FIXME
		utilGin.ErrorResponse(-1, "验证码错误")
		return
	}

	token, err := service.AccountRegister(register.Phone, register.Password)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(token)
}

func Login(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	s, e := bind.Bind(&param_bind.Login{}, c)

	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	login := s.(*param_bind.Login)
	token, err := service.GetLoginToken(login.Phone, login.Password)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}
	utilGin.SuccessResponse(token)
}

func Logout(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	utilGin.SuccessResponse(nil)
}
