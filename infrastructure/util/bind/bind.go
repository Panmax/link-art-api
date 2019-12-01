package bind

import (
	"github.com/gin-gonic/gin"
)

func Bind(s interface{}, c *gin.Context) (interface{}, error) {
	if err := c.ShouldBind(s); err != nil {
		return nil, err
	}
	return s, nil
}
