package context

import (
	"github.com/gin-gonic/gin"
	"strings"
)

type Context struct {
	*gin.Context
}

func (c *Context) Domain() string {
	return c.Request.Host[:strings.Index(c.Request.Host, ":")]
}