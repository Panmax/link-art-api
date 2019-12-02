package requestid

import (
	"github.com/gin-gonic/gin"
	"link-art-api/infrastructure/util/uuid"
)

const header = "X-Request-Id"

func SetUp() gin.HandlerFunc {

	return func(c *gin.Context) {
		requestId := c.Request.Header.Get(header)
		if requestId == "" {
			requestId = uuid.GenUUID()
		}
		c.Set(header, requestId)
		c.Writer.Header().Set(header, requestId)
		c.Next()
	}
}
