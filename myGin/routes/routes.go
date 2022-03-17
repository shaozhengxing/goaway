package routes

import (
	"github.com/gin-gonic/gin"
)

func Load(r *gin.Engine) {
	//r.GET("/", convert(controller.Index))
	router := newRouter(r)

	router.Group("/api", func(api group) {
		api.Group("user", func(user group) {

		})
	})
}

//func convert(f func(ctx *context.Context) *response.Response) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		resp := f(&context.Context{Context: c})
//		data := resp.GetData()
//		switch item := data.(type) {
//		case string:
//			c.String(200, item)
//		case gin.H:
//			c.JSON(200, item)
//		}
//	}
//}
