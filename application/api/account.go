package api

import (
	"link-art-api/application/command"
	"link-art-api/application/middleware"
	"link-art-api/application/representation"
	"link-art-api/domain/model"
	"log"

	"github.com/gin-gonic/gin"
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
		authGroup.POST("/profile", UpdateProfile)
	}
}

func AccountRouterRegister(router *gin.RouterGroup) {

}

func Register(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	s, e := bind.Bind(&command.RegisterCommand{}, c)

	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	registerCommand := s.(*command.RegisterCommand)

	if registerCommand.Sms != "999999" { // FIXME
		utilGin.ErrorResponse(-1, "验证码错误")
		return
	}

	account, err := service.AccountRegister(registerCommand.Phone, registerCommand.Password)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	authMiddleware, err := middleware.NewJWTMiddleware()
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	token, _, err := authMiddleware.TokenGenerator(account)
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

	profile := &representation.AccountProfileRepresentation{}
	profile.Name = account.Name
	profile.Phone = account.Phone
	if account.Birth != nil {
		birthUnix := account.Birth.Unix()
		profile.Birth = &birthUnix
	}
	profile.Gender = account.Gender
	profile.Follow = len(service.ListAccountFollow(account.ID))
	profile.Fans = len(service.ListAccountFans(account.ID))
	profile.Points = 10
	utilGin.SuccessResponse(profile)
}

func UpdateProfile(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	s, e := bind.Bind(&command.UpdateProfileCommand{}, c)
	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	updateCommand := s.(*command.UpdateProfileCommand)
	result, err := service.UpdateProfile(account.ID, updateCommand)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(result)
}
