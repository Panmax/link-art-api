package api

import (
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"log"

	"github.com/gin-gonic/gin"
	"link-art-api/application/param_bind"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/bind"
	"link-art-api/infrastructure/util/response"
)

func AuthRouterRegister(router *gin.RouterGroup) {
	authMiddleware, err := middleware.NewJWTMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	authGroup := router.Group("/auth")
	authGroup.POST("/register", Register)
	authGroup.POST("/login", authMiddleware.LoginHandler)
	authGroup.POST("/refresh_token", authMiddleware.RefreshHandler)

	authGroup.Use(authMiddleware.MiddlewareFunc())
	{
		authGroup.GET("/logout", Logout)
		authGroup.GET("/profile", Profile)
	}
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

func Profile(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	utilGin.SuccessResponse(account.Phone)
}
