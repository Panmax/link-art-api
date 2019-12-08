package bind

import (
	"github.com/gin-gonic/gin"
)

func Bind(command interface{}, c *gin.Context) (interface{}, error) {
	if err := c.ShouldBind(command); err != nil {
		return nil, err
	}
	return command, nil
}
