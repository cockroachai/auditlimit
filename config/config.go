package config

import (
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	PORT           = 8080
	PlusModels     = garray.NewStrArrayFrom([]string{"gpt-4", "gpt-4o", "gpt-4-browsing", "gpt-4-plugins", "gpt-4-mobile", "gpt-4-code-interpreter", "gpt-4-dalle", "gpt-4-gizmo", "gpt-4-magic-create"})
	O1preModels    = garray.NewStrArrayFrom([]string{"o1-preview"})
	O1miniModels   = garray.NewStrArrayFrom([]string{"o1-mini"})
	ForbiddenWords = []string{}         // 禁止词
	LIMIT          = 40                 // 限制次数
	PER            = time.Hour * 3      // 限制时间
	O1PRELIMIT     = 5                  // o1pre限制次数
	O1PREPER       = time.Hour * 24 * 7 // o1pre限制时间
	O1MINILIMIT    = 10                 // o1mini限制次数
	O1MINIPER      = time.Hour * 24     // o1mini限制时间
	OAIKEY         = ""                 // OAIKEY
	OAIKEYLOG      = ""                 // OAIKEYLOG 隐藏
	// MODERATION     = "https://api.openai.com/v1/moderations" // OPENAI Moderation 检测
	MODERATION = "https://gateway.ai.cloudflare.com/v1/a8cace244ffbc233655fefeaca37d515/xyhelper/openai/moderations"
)

func init() {
	ctx := gctx.GetInitCtx()
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)
	limit := g.Cfg().MustGetWithEnv(ctx, "LIMIT").Int()
	if limit > 0 {
		LIMIT = limit
	}
	g.Log().Info(ctx, "LIMIT:", LIMIT)
	per := g.Cfg().MustGetWithEnv(ctx, "PER").Duration()
	if per > 0 {
		PER = per
	}
	g.Log().Info(ctx, "PER:", PER)
	o1prelimit := g.Cfg().MustGetWithEnv(ctx, "O1PRELIMIT").Int()
	if o1prelimit > 0 {
		O1PRELIMIT = o1prelimit
	}
	g.Log().Info(ctx, "O1PRELIMIT:", O1PRELIMIT)
	o1preper := g.Cfg().MustGetWithEnv(ctx, "O1PER").Duration()
	if o1preper > 0 {
		O1PREPER = o1preper
	}
	g.Log().Info(ctx, "O1PREPER:", O1PREPER)
	o1minilimit := g.Cfg().MustGetWithEnv(ctx, "O1MINILIMIT").Int()
	if o1minilimit > 0 {
		O1MINILIMIT = o1minilimit
	}
	g.Log().Info(ctx, "O1MINILIMIT:", O1MINILIMIT)
	o1miniper := g.Cfg().MustGetWithEnv(ctx, "O1MINIPER").Duration()
	if o1miniper > 0 {
		O1MINIPER = o1miniper
	}
	g.Log().Info(ctx, "O1MINIPER:", O1MINIPER)
	oaikey := g.Cfg().MustGetWithEnv(ctx, "OAIKEY").String()
	// oaikey 不为空
	if oaikey != "" {
		OAIKEY = oaikey
		// 日志隐藏 oaikey，有 * 代表有值
		OAIKEYLOG = "******"
	}
	g.Log().Info(ctx, "OAIKEY:", OAIKEYLOG)
	moderation := g.Cfg().MustGetWithEnv(ctx, "MODERATION").String()
	if moderation != "" {
		MODERATION = moderation
	}
	g.Log().Info(ctx, "MODERATION:", MODERATION)
}
