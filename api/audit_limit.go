package api

import (
	"auditlimit/config"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func AuditLimit(r *ghttp.Request) {
	ctx := r.Context()
	// 获取Bearer Token 用来判断用户身份
	token := r.Header.Get("Authorization")
	// 移除Bearer
	if token != "" {
		token = token[7:]
	}
	g.Log().Debug(ctx, "token", token)
	// 获取gfsessionid 可以用来分析用户是否多设备登录
	gfsessionid := r.Cookie.Get("gfsessionid").String()
	g.Log().Debug(ctx, "gfsessionid", gfsessionid)
	// 获取referer 可以用来判断用户请求来源
	referer := r.Header.Get("referer")
	g.Log().Debug(ctx, "referer", referer)
	// 获取请求内容
	reqJson, err := r.GetJson()
	if err != nil {
		g.Log().Error(ctx, "GetJson", err)
		r.Response.Status = 400
		r.Response.WriteJson(g.Map{
			"detail": "请求格式错误.",
		})
	}
	action := reqJson.Get("action").String() // action为 next时才是真正的请求，否则可能是继续上次请求 action 为 variant 时为重新生成
	g.Log().Debug(ctx, "action", action)

	model := reqJson.Get("model").String() // 模型名称
	g.Log().Debug(ctx, "model", model)
	prompt := reqJson.Get("messages.0.content.parts.0").String() // 输入内容
	g.Log().Debug(ctx, "prompt", prompt)

	// 判断提问内容是否包含禁止词
	if containsAny(prompt, config.ForbiddenWords) {
		r.Response.Status = 400
		r.Response.WriteJson(g.Map{
			"detail": "请珍惜账号,不要提问违禁内容.",
		})
		return
	}

	// 判断模型是否为plus模型 如果是则使用plus模型的限制
	if config.PlusModels.Contains(model) {
		limiter := GetVisitor(token, config.LIMIT, config.PER)
		// 获取剩余次数
		remain := limiter.TokensAt(time.Now())
		g.Log().Debug(ctx, "remain", remain)
		if remain < 1 {
			r.Response.Status = 429
			resMsg := gjson.New(MsgPlus429)
			// 根据remain计算需要等待的时间
			// 生产间隔
			creatInterval := config.PER / time.Duration(config.LIMIT)
			// 转换为秒
			creatIntervalSec := float64(creatInterval.Seconds())
			// 等待时间
			wait := (1 - remain) * creatIntervalSec
			g.Log().Debug(ctx, "wait", wait, "creatIntervalSec", creatIntervalSec)
			resMsg.Set("detail.clears_in", int(wait))
			r.Response.WriteJson(resMsg)
			return
		} else {
			// 消耗一个令牌
			limiter.Allow()
			r.Response.Status = 200
			return
		}

	}

	r.Response.Status = 200

}

// 判断字符串是否包含数组中的任意一个元素
func containsAny(text string, array []string) bool {
	for _, item := range array {
		if strings.Contains(text, item) {
			return true
		}
	}
	return false
}
