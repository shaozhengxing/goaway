package kernel

import (
	"goaway/context"
	"goaway/exception"
)

// Middleware 全局中间件
var Middleware []context.HandlerFunc

func Load() {
	Middleware = []context.HandlerFunc{
		exception.Exception,
	}
}
