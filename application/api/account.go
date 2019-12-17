package api

import (
	"github.com/gin-gonic/gin"
	"link-art-api/application/command"
	"link-art-api/application/middleware"
	"link-art-api/domain/model"
	"link-art-api/domain/service"
	"link-art-api/infrastructure/util/bind"
	"link-art-api/infrastructure/util/response"
	"strconv"
)

func AuthRouterRegister(group *gin.RouterGroup) {
	authGroup := group.Group("/auth")
	{
		authGroup.POST("/register", Register)
		authGroup.POST("/login", middleware.JWTMiddleware.LoginHandler)
		authGroup.POST("/refresh_token", middleware.JWTMiddleware.RefreshHandler)

		authGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			authGroup.GET("/logout", Logout)
			authGroup.GET("/profile", Profile)
			authGroup.POST("/profile", UpdateProfile)
			authGroup.PUT("/avatar", UpdateAvatar)
		}
	}
}

func AccountRouterRegister(group *gin.RouterGroup) {
	accountGroup := group.Group("/accounts")
	{
		accountGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			accountGroup.POST("/approval", SubmitApproval)
		}
	}
}

func ArtistRouterRegister(group *gin.RouterGroup) {
	artistGroup := group.Group("/artists")
	{
		artistGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			artistGroup.GET("/:id", GetArtist)
			artistGroup.POST("/:id/follow", Follow)
			artistGroup.DELETE("/:id/follow", UnFollow)
		}
	}
}

func Register(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	cmd, e := bind.Bind(&command.RegisterCommand{}, c)

	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	registerCommand := cmd.(*command.RegisterCommand)

	if registerCommand.Sms != "999999" { // FIXME
		utilGin.ErrorResponse(-1, "验证码错误")
		return
	}

	account, err := service.AccountRegister(registerCommand.Phone, registerCommand.Password)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	token, _, err := middleware.JWTMiddleware.TokenGenerator(account)
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

	profile, err := service.GetProfile(account.ID)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(profile)
}

func UpdateProfile(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	cmd, e := bind.Bind(&command.UpdateProfileCommand{}, c)
	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	updateCommand := cmd.(*command.UpdateProfileCommand)
	result, err := service.UpdateProfile(account.ID, updateCommand)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(result)
}

func UpdateAvatar(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	cmd, e := bind.Bind(&command.UpdateAvatarCommand{}, c)
	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}
	updateCommand := cmd.(*command.UpdateAvatarCommand)
	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	account.UpdateAvatar(&updateCommand.Url)
	if err := model.SaveOne(account); err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func SubmitApproval(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	cmd, err := bind.Bind(&command.SubmitApprovalCommand{}, c)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}
	submitCommand := cmd.(*command.SubmitApprovalCommand)
	if submitCommand.Type != model.ApprovalPersonalType && submitCommand.Type != model.ApprovalCompanyType {
		utilGin.ParamErrorResponse("申请类型错误")
		return
	}

	if err := service.SubmitApproval(account.ID, submitCommand); err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(true)
}

func GetArtist(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	artist, err := service.GetArtist(uint(accountID))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(artist)
}

func Follow(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	followerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	if err := service.Follow(account.ID, uint(followerID)); err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}
	utilGin.SuccessResponse(true)
}

func UnFollow(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	followerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	if err := service.UnFollow(account.ID, uint(followerID)); err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}
	utilGin.SuccessResponse(true)
}
