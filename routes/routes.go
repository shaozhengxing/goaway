package routes

import (
	"github.com/gin-gonic/gin"
	"goaway/controller"
	"goaway/middleware/kernel"
)

func config(router group) {
	//router.Group("/api", func(api group) {
	//	api.Group("/user", func(user group) {
	//		user.Registered(GET, "/info", controller.Index)
	//		user.Registered(GET, "/order", controller.Index)
	//		user.Registered(GET, "/money", controller.Index)
	//	})
	//})

	//router.Registered(GET, "/", controller.Index)
	//router.Registered(GET, "/", controller.Index2)
	//router.Registered(GET, "/", controller.Index3)
	router.Registered(GET, "/", controller.Index4)
}

func Load(r *gin.Engine) {
	router := newRouter(r)

	router.Group("", func(g group) {
		config(g)
	}, kernel.Middleware...) // 加载全局中间件
}
