package cmd

import (
	"context"

	"kubevision/internal/controller"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Bind(
					new(controller.Nodes),
					new(controller.Namespaces),
					new(controller.Daemonsets),
					new(controller.Deployments),
					new(controller.Pods),
				)

			})
			// s.SetAddr(":8080")
			// s.BindObjectMethod("/namespaces", new(controller.Namespace), "Index")

			logging.BasicConfig(logging.LogConfig{Level: logging.DEBUG})
			logging.Info("start")
			s.Run()
			return nil
		},
	}
)
