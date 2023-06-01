package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func TestHLog(ctx context.Context, c *app.RequestContext) {
	// it will be output
	hlog.Info("Hello, hertz Info")
	hlog.Debug("Hello, hertz Debug")
	hlog.Warn("Hello, hertz Warn")
	// it will not be output
	hlog.Trace("Hello, hertz")
	c.JSON(
		consts.StatusOK,
		utils.H{
			"msg":     "test_hlog ok",
			"boolean": true,
		},
	)
}
