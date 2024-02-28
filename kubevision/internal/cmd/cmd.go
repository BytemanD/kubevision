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
				group.Bind(
					controller.Hello,
				)
			})
			// s.SetAddr(":8080")
			// s.BindObjectMethod("/namespaces", new(controller.Namespace), "Index")

			logging.BasicConfig(logging.LogConfig{Level: logging.DEBUG})
			logging.Info("start")
			s.BindObjectRest("/nodes", new(controller.Nodes))
			s.BindObjectRest("/pods", new(controller.Pods))
			s.BindObjectRest("/namespaces", new(controller.Namespaces))
			s.BindObjectRest("/daemonsets", new(controller.Daemonsets))
			s.BindObjectRest("/deployments", new(controller.Deployments))
			s.Run()
			return nil
		},
	}
)
