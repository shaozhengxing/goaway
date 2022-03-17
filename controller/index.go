package controller

import (
	"goaway/context"
	"goaway/response"
)

func Index(ctx *context.Context) *response.Response {
	panic("假装错误")

	//return response.Resp().Json(gin.H{"msg": "hello world"})
	return response.Resp().String(ctx.Domain())
}
