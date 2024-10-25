package api

import (
	"auditlimit/config"
	"strings"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
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
			"error": err.Error(),
		})
	}
	action := reqJson.Get("action").String() // action为 next时才是真正的请求，否则可能是继续上次请求 action 为 variant 时为重新生成
	g.Log().Debug(ctx, "action", action)

	model := reqJson.Get("model").String() // 模型名称
	g.Log().Debug(ctx, "model", model)
	prompt := reqJson.Get("messages.0.content.parts.0").String() // 输入内容
	g.Log().Debug(ctx, "prompt", prompt)

	// 判断提问内容是否包含禁止词
	if containsAny(ctx, prompt, config.ForbiddenWords) {
		r.Response.Status = 400
		r.Response.WriteJson(g.Map{
			"error": "请珍惜账号,不要提问违禁内容.",
		})
		return
	}

	// OPENAI Moderation 检测
	if config.OAIKEY != "" && prompt != "" {
		// 检测是否包含违规内容
		respVar := g.Client().SetHeaderMap(g.MapStrStr{
			"Authorization": "Bearer " + config.OAIKEY,
			"Content-Type":  "application/json",
		}).PostVar(ctx, config.MODERATION, g.Map{
			"input": prompt,
		})

		// 返回的 json 中 results.flagged 为 true 时为违规内容
		// respBody := resp.ReadAllString()
		//g.Log().Debug(ctx, "resp:", respBody)
		g.Dump(respVar)
		respJson := gjson.New(respVar)
		isFlagged := respJson.Get("results.0.flagged").Bool()
		g.Log().Debug(ctx, "flagged", isFlagged)
		if isFlagged {
			r.Response.Status = 400
			r.Response.WriteJson(MsgMod400)
			return
		}
	}
	limit, per, limiter, err := GetVisitorWithModel(ctx, token, model)
	if err != nil {
		g.Log().Error(ctx, "GetVisitorWithModel", err)
		r.Response.Status = 500
		r.Response.WriteJson(g.Map{
			"error": err.Error(),
		})
		return
	}
	// 获取剩余次数
	remain := limiter.TokensAt(time.Now())
	g.Log().Debug(ctx, token, model, "remain", remain, "limit", limit, "per", per)
	if remain < 1 {
		r.Response.Status = 429
		reservation := limiter.ReserveN(time.Now(), 1)
		if !reservation.OK() {
			// 处理预留失败的情况，例如返回错误
			r.Response.WriteJson(g.Map{
				"error": "You have triggered the usage frequency limit of " + model + ", the current limit is " + gconv.String(limit) + " times/" + gconv.String(per) + ", please wait a moment before trying again.\n" + "您已经触发 " + model + " 使用频率限制,当前限制为 " + gconv.String(limit) + " 次/" + gconv.String(per) + ",请稍后再试.",
			})
			reservation.Cancel() // 取消预留，不消耗令牌
			return
		}
		delayFrom := reservation.Delay()
		reservation.Cancel() // 取消预留，不消耗令牌

		g.Log().Debug(ctx, "delayFrom", delayFrom)
		r.Response.WriteJson(g.Map{
			"error": "You have triggered the usage frequency limit of " + model + ", the current limit is " + gconv.String(limit) + " times/" + gconv.String(per) + ", please wait " + gconv.String(int(delayFrom.Seconds())) + " seconds before trying again.\n" + "您已经触发 " + model + " 使用频率限制,当前限制为 " + gconv.String(limit) + " 次/" + gconv.String(per) + ",请等待 " + gconv.String(int(delayFrom.Seconds())) + " 秒后再试.",
		})
		return
	}
	// 消耗一个令牌
	limiter.Allow()

	r.Response.Status = 200

}

// 判断字符串是否包含数组中的任意一个元素
func containsAny(ctx g.Ctx, text string, array []string) bool {
	for _, item := range array {
		if strings.Contains(text, item) {
			g.Log().Debug(ctx, "containsAny", text, item)
			return true
		}
	}
	return false
}
