package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type returnData struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, returnData{
		Code: 200,
		Msg:  "success",
		Data: data,
	})
}

func Fail(ctx *gin.Context, format string, a ...any) {
	ctx.JSON(200, returnData{
		Code: 500,
		Msg:  fmt.Sprintf(format, a...),
		Data: nil,
	})
}
