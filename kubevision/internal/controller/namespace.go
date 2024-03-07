package controller

import (
	"context"
	"encoding/json"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/gogf/gf/v2/frame/g"
)

type Namespaces struct{}

func (c *Namespaces) Get(ctx context.Context, apiReq *apiv1.NamespacesListReq) (res *apiv1.NamespacesListRes, err error) {
	req := g.RequestFromCtx(ctx)

	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}

	namespaces, err := client.ListNamespaces()
	if err != nil {
		logging.Error("%v", err)
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
