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
		authGroup.POST("/send_sms", SendSms)
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
func SendSms(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	cmd, e := bind.Bind(&command.SendSmsCommand{}, c)

	if e != nil {
		utilGin.ParamErrorResponse(e.Error())
		return
	}

	SendSmsCommand := cmd.(*command.SendSmsCommand)
	code := service.SendSms(SendSmsCommand.Phone)

	utilGin.SuccessResponse(code)

}

func UserRouterRegister(group *gin.RouterGroup) {
	userGroup := group.Group("/users")
	{
		userGroup.Use(middleware.JWTMiddleware.MiddlewareFunc())
		{
			userGroup.GET("/:id", GetUser)
			userGroup.POST("/:id/follow", Follow)
			userGroup.DELETE("/:id/follow", UnFollow)
			userGroup.GET("/:id/follows", ListFollow)
			userGroup.GET("/:id/fans", ListFans)
		}
	}

	group.GET("/artists/search", SearchArtist)
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

func GetUser(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	accountID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	user, err := service.GetUser(uint(accountID))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(user)
}

func Follow(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	account := c.MustGet(middleware.IdentityKey).(*model.Account)

	followerID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	if account.ID == uint(followerID) {
		utilGin.ParamErrorResponse("无法关注自己")
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

func ListFollow(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	accountId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	follows, err := service.ListFollow(uint(accountId))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	// 判断当前登录用户是否关注
	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	for _, user := range follows {
		user.Follow = service.CheckFollow(account.ID, user.ID)
	}

	utilGin.SuccessResponse(follows)
}

func ListFans(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	accountId, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utilGin.ParamErrorResponse(err.Error())
		return
	}

	fans, err := service.ListFans(uint(accountId))
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	// 判断当前登录用户是否关注
	account := c.MustGet(middleware.IdentityKey).(*model.Account)
	for _, user := range fans {
		user.Follow = service.CheckFollow(account.ID, user.ID)
	}

	utilGin.SuccessResponse(fans)
}

func SearchArtist(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}
	keyword := c.Query("keyword")
	results, err := service.SearchArtist(keyword)
	if err != nil {
		utilGin.ErrorResponse(-1, err.Error())
		return
	}

	utilGin.SuccessResponse(results)
}
