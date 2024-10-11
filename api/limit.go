package api

import (
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/time/rate"
)

type visitor struct {
	limiter  *rate.Limiter
	lastSeen time.Time
	Per      time.Duration
}

var visitors = make(map[string]*visitor)
var mtx sync.Mutex

func GetVisitor(key string, limit int, per time.Duration) *rate.Limiter {
	mtx.Lock()
	defer mtx.Unlock()

	v, exists := visitors[key]
	if !exists {
		limiter := rate.NewLimiter(rate.Every(per/time.Duration(limit)), limit)
		visitors[key] = &visitor{limiter, time.Now(), per}
		return limiter
	}

	v.lastSeen = time.Now()
	return v.limiter
}

func GetVisitorWithModel(ctx g.Ctx, token, model string) (limit int, per time.Duration, limiter *rate.Limiter, err error) {
	model = gstr.ToUpper(model)
	modelrate := g.Cfg().MustGetWithEnv(ctx, model, "40/3h").String()
	modelratearr := strings.Split(modelrate, "/")
	// g.Dump(modelratearr)
	if len(modelratearr) != 2 {
		modelratearr = []string{"40", "3h"}
	}
	limit = gconv.Int(modelratearr[0])
	// per = gconv.Duration(modelratearr[1])
	per, err = time.ParseDuration(modelratearr[1])
	if err != nil {
		return 0, 0, nil, err
	}
	return limit, per, GetVisitor(token+"|"+model, limit, per), nil

}

func CleanupVisitors() {
	mtx.Lock()
	defer mtx.Unlock()

	for token, v := range visitors {
		if time.Since(v.lastSeen) > v.Per {
			delete(visitors, token)
		}
	}
}

func init() {
	// 每星期清理一次
	go func() {
		for {
			time.Sleep(time.Hour * 24 * 7)
			CleanupVisitors()
		}
	}()
}
