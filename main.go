package main

import (
	"github.com/gin-gonic/gin"
	"goaway/kernel"
	"goaway/routes"
)

func main() {
	r := gin.Default()

	// 加载全局变量
	kernel.Load()

	routes.Load(r)

	r.Run()
}
