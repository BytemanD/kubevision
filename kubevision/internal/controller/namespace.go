package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type Namespaces struct{}

func (c *Namespaces) Get(ctx context.Context, apiReq *apiv1.NamespacesListReq) (res *apiv1.NamespacesListRes, err error) {
	req := g.RequestFromCtx(ctx)

	client, err := k8s.GetClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	namespaces, err := client.ListNamespaces()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Namespace{"namespaces": namespaces})
	req.Response.WriteJson(data)
	return
}
func (h *Namespaces) Post(ctx context.Context, req *apiv1.NamespacesPostReq) (res *apiv1.NamespacesPostRes, err error) {
	request := g.RequestFromCtx(ctx)
	request.Response.Writeln("Hello World!")
	return
}

func getReqNamespace(req *ghttp.Request) string {
	namespace := req.Header.Get("X-Namespace")
	if namespace != "" {
		return namespace
	} else {
		return "default"
	}
}
