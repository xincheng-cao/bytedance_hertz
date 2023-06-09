// Code generated by hertz generator.

package main

import (
	handler "bytedance_hertz_mod/biz/handler"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", handler.Ping)

	// your code ...
	r.POST(
		"/post",
		func(c context.Context, ctx *app.RequestContext) {
			ctx.String(consts.StatusOK, `{"res":"ok"}`)
			resp_body, _ := ctx.Body()
			fmt.Println(string(resp_body))
		},
	)

	r.GET(
		"/test_hlog",
		handler.TestHLog,
	)
}
