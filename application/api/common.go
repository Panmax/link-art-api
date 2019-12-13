package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"link-art-api/infrastructure/config"
	"link-art-api/infrastructure/util/response"
)

func CommonRouterRegister(group *gin.RouterGroup) {
	commonGroup := group.Group("/common")
	{
		commonGroup.GET("/oss/token", GetOssToken)
	}
}

func GetOssToken(c *gin.Context) {
	utilGin := response.Gin{Ctx: c}

	putPolicy := storage.PutPolicy{Scope: config.QiniuConfig.Bucket}
	mac := qbox.NewMac(config.QiniuConfig.AccessKey, config.QiniuConfig.SecretKey)
	upToken := putPolicy.UploadToken(mac)

	utilGin.SuccessResponse(upToken)
}
