package config

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

var (
	PORT = 8080
	// PlusModels     = garray.NewStrArrayFrom([]string{"gpt-4", "gpt-4o", "gpt-4-browsing", "gpt-4-plugins", "gpt-4-mobile", "gpt-4-code-interpreter", "gpt-4-dalle", "gpt-4-gizmo", "gpt-4-magic-create", "gpt-4o-canmore"})
	// O1Models       = garray.NewStrArrayFrom([]string{"o1-preview", "o1-mini"})
	ForbiddenWords = []string{} // 禁止词
	// LIMIT          = 40                 // 限制次数
	// PER            = time.Hour * 3      // 限制时间
	// O1LIMIT        = 5                  // 限制次数
	// O1PER          = time.Hour * 24 * 7 // 限制时间
	OAIKEY    = "" // OAIKEY
	OAIKEYLOG = "" // OAIKEYLOG 隐藏
	// MODERATION     = "https://api.openai.com/v1/moderations" // OPENAI Moderation 检测
	MODERATION = "https://gateway.ai.cloudflare.com/v1/040ac2002b4dd67637e97c628feb3484/xyhelper/openai/moderations"
)

func init() {
	ctx := gctx.GetInitCtx()
	port := g.Cfg().MustGetWithEnv(ctx, "PORT").Int()
	if port > 0 {
		PORT = port
	}
	g.Log().Info(ctx, "PORT:", PORT)
	// limit := g.Cfg().MustGetWithEnv(ctx, "LIMIT").Int()
	// if limit > 0 {
	// 	LIMIT = limit
	// }
	// g.Log().Info(ctx, "LIMIT:", LIMIT)
	// per := g.Cfg().MustGetWithEnv(ctx, "PER").Duration()
	// if per > 0 {
	// 	PER = per
	// }
	// g.Log().Info(ctx, "PER:", PER)
	// o1limit := g.Cfg().MustGetWithEnv(ctx, "O1LIMIT").Int()
	// if o1limit > 0 {
	// 	O1LIMIT = o1limit
	// }
	// g.Log().Info(ctx, "O1LIMIT:", O1LIMIT)
	// o1per := g.Cfg().MustGetWithEnv(ctx, "O1PER").Duration()
	// if o1per > 0 {
	// 	O1PER = o1per
	// }
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
