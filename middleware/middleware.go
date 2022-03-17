package middleware

import (
	"fmt"
	"goaway/context"
)

func M1(c *context.Context) {
	fmt.Println("我的1")
}

func M2(c *context.Context) {
	fmt.Println("我的2")
}

func M3(c *context.Context) {
	fmt.Println("我的3")
}
