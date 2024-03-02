package cmd

import (
	"context"
	"fmt"

	"kubevision/internal/controller"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

func MiddlewareCORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	// r.Response.Header().Set("Access-Control-Expose-Headers", "X-Token")
	// r.Response.Header().Set("Access-Control-Allow-Headers", "*")
	r.Middleware.Next()
}
func MiddlewareLogResponse(r *ghttp.Request) {
	r.Middleware.Next()
	logging.Info("%s %s -> %d", r.Method, r.URL, r.Response.Status)
}
func MiddlewareAuth(req *ghttp.Request) {
	token := req.Header.Get("X-Token")
	if token == "" {
		logging.Error("no auth")
		req.Response.WriteStatusExit(403, "not auth")
		return
	}
	req.Middleware.Next()
}

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					new(controller.Login),
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareAuth)
				group.Middleware(MiddlewareLogResponse)
				group.Bind(
					new(controller.Cluster),
					new(controller.Version),
					new(controller.Nodes),
					new(controller.Namespaces),
					new(controller.Daemonsets),
					new(controller.Deployments),
					new(controller.Pods),
					new(controller.Jobs),
					new(controller.CronJobs),
					new(controller.Services),
					new(controller.StatefulSets),
					new(controller.Events),
					new(controller.ConfigMaps),
					new(controller.Secrets),
				)
			})
			port := parser.GetOpt("port", "8091").String()
			if port != "" {
				s.SetAddr(fmt.Sprintf(":%s", port))
			}
			// s.BindObjectMethod("/namespaces", new(controller.Namespace), "Index")

			logging.BasicConfig(logging.LogConfig{Level: logging.DEBUG})
			logging.Info("start")
			s.Run()
			return nil
		},
	}
)

func init() {
	Main.Arguments = append(Main.Arguments, gcmd.Argument{
		Name: "port", Short: "p", Brief: "The port of server",
	})
}
