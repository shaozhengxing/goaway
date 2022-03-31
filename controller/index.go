package controller

import (
	"goaway/context"
	"goaway/response"
	"strconv"
)

func Index(context *context.Context) *response.Response {
	//panic("假装错误")
	//return response.Resp().Json(gin.H{"msg": "hello world"})

	context.Session().Set("msg", "golang 是世界上最好的语言")

	return response.Resp().String(context.Domain())
}

func Index2(context *context.Context) *response.Response {
	msg, _ := context.Session().Get("msg")

	return response.Resp().String(msg.(string))
}

func Index3(context *context.Context) *response.Response {
	context.Session().Remove("msg")

	return response.Resp().String("")
}

func Index4(context *context.Context) *response.Response {
	session := context.Session()

	for i := 0; i < 100; i++ {
		go func(index int) {
			session.Set("msg"+strconv.Itoa(index), index)
		}(i)
	}

	return response.Resp().String("")
}
