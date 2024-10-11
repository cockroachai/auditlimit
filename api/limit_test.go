package api_test

import (
	"auditlimit/api"
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func TestGetVisitorWithModel(t *testing.T) {
	ctx := gctx.New()
	limit, per, limiter, err := api.GetVisitorWithModel(ctx, "token", "text-davinci-002-render-sha")
	if err != nil {
		g.Log().Error(ctx, "GetVisitorWithModel", err)
		return
	}
	g.Dump(limiter)
	g.Log().Info(ctx, "limit:", limit, "per:", per, "limiter:", limiter)

}
