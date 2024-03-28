package cmd

import (
	"context"
	"fmt"
	"time"

	"kubevision/internal/controller"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/glog"
)

func MiddlewareCORS(req *ghttp.Request) {
	req.Response.CORSDefault()
	req.Response.Header().Set("Access-Control-Expose-Headers", "X-Auth-Token")
	req.Middleware.Next()
}
func MiddlewareResponseStatus(r *ghttp.Request) {
	startTime := time.Now()
	r.Middleware.Next()
	spentTime := time.Since(startTime).Seconds()
	if r.Response.Status < 400 {
		logging.Info("%s %s -> %d (%fs)", r.Method, r.URL, r.Response.Status, spentTime)
	} else {
		logging.Error("%s %s -> %d (%fs)\n    Resp: %s",
			r.Method, r.URL, r.Response.Status, spentTime,
			r.Response.BufferString())
	}
}
func MiddlewareAuth(req *ghttp.Request) {
	token := req.Header.Get("X-Auth-Token")
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
				group.Middleware(MiddlewareResponseStatus)
				group.Bind(
					new(controller.Login),
				)
			})
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(MiddlewareCORS)
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Middleware(MiddlewareAuth)
				group.Middleware(MiddlewareResponseStatus)
				group.Bind(
					new(controller.Cluster),
					new(controller.Version),
					new(controller.Nodes),
					new(controller.Namespaces),
					new(controller.Daemonsets),
					new(controller.Deployments),
					new(controller.Pods), new(controller.Pod),
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
			debug := parser.ContainsOpt("debug")

			if port != "" {
				s.SetAddr(fmt.Sprintf(":%s", port))
			}
			level := logging.INFO

			if debug {
				level = logging.DEBUG
			}

			// s.BindObjectMethod("/namespaces", new(controller.Namespace), "Index")

			logging.BasicConfig(logging.LogConfig{Level: level})
			glog.Info(context.TODO(), "start server")
			s.Run()
			return nil
		},
	}
)

func init() {
	Main.Arguments = append(
		Main.Arguments,
		gcmd.Argument{Name: "port", Short: "p", Brief: "The port of server"},
		gcmd.Argument{Name: "debug", Short: "d", Orphan: true, Brief: "Show DEBUG message"},
	)
}
