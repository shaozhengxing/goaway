package exception

import (
	"fmt"
	"goaway/context"
	"runtime/debug"
)

func Exception(ctx *context.Context) {
	defer func() {
		if r := recover(); r != nil {
			msg := fmt.Sprint(r) + "\n" + string(debug.Stack())
			ctx.String(500, msg)
			ctx.Abort()
		}
	}()

	ctx.Next()
}
