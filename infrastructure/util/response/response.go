package response

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (g *Gin) Response(code int, message string, data interface{}) {
	g.Ctx.JSON(200, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
	return
}

func (g *Gin) SuccessResponse(data interface{}) {
	g.Response(0, "ok", data)
	return
}

func (g *Gin) ErrorResponse(code int, message string) {
	g.Response(code, message, nil)
	return
}

func (g *Gin) ParamErrorResponse(message string) {
	g.Response(400, message, nil)
	return
}
