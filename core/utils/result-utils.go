package utils

import "github.com/gin-gonic/gin"

type Response struct {
	Code    int32       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(ctx *gin.Context, data ...interface{}) {
	var res Response
	const code = 200
	const message = "Success"
	if len(data) > 0 {
		res = Response{
			Code:    code,
			Message: message,
			Data:    data[0],
		}
	} else {
		res = Response{
			Code:    code,
			Message: message,
		}
	}
	ctx.JSON(200, res)
}

func Fail(ctx *gin.Context, code int32, message string) {
	res := Response{
		Code:    code,
		Message: message,
	}
	ctx.JSON(200, res)
}
