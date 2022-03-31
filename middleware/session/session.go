package session

import (
	context2 "context"
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"goaway/context"
	"goaway/redis"
	"time"
)

var cookieName = "my_gin"

var leftTime = 3600

func Session(c *context.Context) {
	cookie, err := c.Cookie(cookieName)
	if err == nil {
		sessionString, err := redis.Client().Get(context2.TODO(), cookie).Result()
		if err == nil {
			var session context.Session
			err := json.Unmarshal([]byte(sessionString), &session)
			if err == nil {
				c.Set("_session", session)
			}
			return
		}
	}

	sessionKey := uuid.NewV4().String()

	c.SetCookie(cookieName, sessionKey, leftTime, "/", c.Domain(), false, true)

	session := context.Session{
		Cookie:      sessionKey,
		ExpireTime:  time.Now().Unix() + int64(leftTime),
		SessionList: make(map[string]interface{}),
	}
	jsonString, _ := json.Marshal(session)

	redis.Client().Set(context2.TODO(), sessionKey, jsonString, time.Second*time.Duration(leftTime))
}
