package session

import (
	uuid "github.com/satori/go.uuid"
	"goaway/context"
)

var cookedName = "my_gin"

func Session(c *context.Context) {
	sessionKey := uuid.NewV4().String()

	c.SetCookie(cookedName, sessionKey, 3600, "/", c.Domain(), false, true)
}
