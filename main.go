package main

import (
	"fmt"
	"goaway/cache"
	"time"
)

func main() {
	//r := gin.Default()
	//
	//// 加载全局变量
	//kernel.Load()
	//
	//routes.Load(r)
	//
	//r.Run()

	c := cache.Cache("file")

	c.Put("peter", "nice", 1*time.Minute)

	value, _ := c.Get("peter")
	fmt.Println(value)
}
