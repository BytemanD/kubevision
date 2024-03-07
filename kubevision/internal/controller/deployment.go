package controller

import (
	"context"
	"encoding/json"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"
	"kubevision/utility"

	"github.com/gogf/gf/v2/frame/g"
)

type Deployments struct{}

func (c *Deployments) Get(ctx context.Context, apiReq *apiv1.DeploymentsListReq) (res *apiv1.DeploymentsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}

	deployments, err := client.ListDeployments(namespace)
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}
	data, _ := json.Marshal(map[string][]model.Deployment{"deployments": deployments})
	req.Response.WriteJson(data)
	return
}
