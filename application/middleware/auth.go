package middleware

import (
	"github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"link-art-api/application/param_bind"
	"link-art-api/domain/model"
	"link-art-api/infrastructure/config"
	"time"
)

const IdentityKey = "_account"

func NewJWTMiddleware() (*jwt.GinJWTMiddleware, error) {
	middleware := jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte(config.AppConfig.JwtSecret),
		Timeout:     8 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: IdentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.Account); ok {
				return jwt.MapClaims{
					IdentityKey: v.Phone,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)

			account, err := model.FindAccountByPhone(claims[IdentityKey].(string))
			if err != nil {
				return nil
			}
			return &account
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVal param_bind.Login
			if err := c.ShouldBind(&loginVal); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			phone := loginVal.Phone
			password := loginVal.Password

			account, err := model.FindAccountByPhone(phone)
			if err == nil && account.CheckPassword(password) {
				return &account, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.Account); ok {
				return true
			}

			return false
		},
	}
	return jwt.New(&middleware)
}
