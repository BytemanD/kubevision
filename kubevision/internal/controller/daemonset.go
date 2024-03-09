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

type Daemonsets struct{}

func (c *Daemonsets) Get(ctx context.Context, apiReq *apiv1.DaemonsetsListReq) (res *apiv1.DaemonsetsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}

	daemonsets, err := client.ListDaemonsets(namespace)
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}
	data, _ := json.Marshal(map[string][]model.Daemonset{"daemonsets": daemonsets})
	req.Response.WriteJson(data)
	return
}

func (c *Daemonsets) Post(ctx context.Context, apiReq *apiv1.DaemonsetsPostReq) (res *apiv1.DaemonsetsPostRes, err error) {
	return

}
func (c *Daemonsets) Delete(ctx context.Context, apiReq *apiv1.DaemonsetsDelteReq) (res *apiv1.DaemonsetsDeleteRes, err error) {
	return
}
