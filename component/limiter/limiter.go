package limiter

import (
	"golang.org/x/time/rate"
	"sync"
	"time"
)

type Limiters struct {
	limiters *sync.Map
}

type Limiter struct {
	limiter *rate.Limiter
	lastGet time.Time // 上一次获取token的时间
	key     string
}

var GlobalLimiters = &Limiters{
	limiters: &sync.Map{},
}

func NewLimiter(l rate.Limit, b int, key string) *Limiter {
	keyLimiter := GlobalLimiters.getLimiter(l, b, key)

	return keyLimiter
}

func (l *Limiter) Allow() bool {
	l.lastGet = time.Now()

	return l.limiter.Allow()
}

func (ls *Limiters) getLimiter(r rate.Limit, b int, key string) *Limiter {
	limiter, ok := ls.limiters.Load(key)
	if ok {
		return limiter.(*Limiter)
	}

	l := &Limiter{
		limiter: rate.NewLimiter(r, b),
		lastGet: time.Now(),
		key:     key,
	}

	ls.limiters.Store(key, l)

	return l
}
