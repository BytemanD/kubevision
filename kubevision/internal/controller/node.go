package controller

import (
	"context"
	"encoding/json"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"

	"github.com/gogf/gf/v2/frame/g"
)

type Nodes struct{}

func (h *Nodes) List(ctx context.Context, apiReq *apiv1.NodesListReq) (res *apiv1.NodesListRes, err error) {
	req := g.RequestFromCtx(ctx)
	client, err := k8s.GetClient()

	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}

	nodes, err := client.ListNodes()
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}
	data, _ := json.Marshal(map[string][]model.Node{"nodes": nodes})
	req.Response.WriteJson(data)
	return
}
