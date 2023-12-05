package api

import (
	"auditlimit/config"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var visitors = make(map[string]*visitor)
var mtx sync.Mutex

func GetVisitor(token string, limit int, per time.Duration) *rate.Limiter {
	mtx.Lock()
	defer mtx.Unlock()

	v, exists := visitors[token]
	if !exists {
		limiter := rate.NewLimiter(rate.Every(per/time.Duration(limit)), limit)
		visitors[token] = &visitor{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

func CleanupVisitors() {
	mtx.Lock()
	defer mtx.Unlock()

	for token, v := range visitors {
		if time.Since(v.lastSeen) > config.PER {
			delete(visitors, token)
		}
	}
}

func init() {
	// 每小时清理一次
	go func() {
		for {
			time.Sleep(1 * time.Hour)
			CleanupVisitors()
		}
	}()
}
