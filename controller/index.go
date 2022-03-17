package controller

import (
	"myGin/context"
	"myGin/response"
)

func Index(ctx *context.Context) *response.Response {
	//return response.Resp().Json(gin.H{"msg": "hello world"})
	return response.Resp().String(ctx.Domain())
}
