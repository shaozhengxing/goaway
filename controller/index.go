package controller

import (
	"goaway/component/limiter"
	"goaway/context"
	"goaway/response"
	"golang.org/x/time/rate"
	"strconv"
	"time"
)

func Index(context *context.Context) *response.Response {
	//return response.Resp().Json(gin.H{"msg": "hello world"})
	//panic("假装错误")
	//context.Session().Set("msg", "golang 是世界上最好的语言")

	l := limiter.NewLimiter(rate.Every(1*time.Second), 5, context.ClientIP())
	if !l.Allow() {
		return response.Resp().String("您的访问过于频繁")
	}

	return response.Resp().String(time.Now().String())
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
