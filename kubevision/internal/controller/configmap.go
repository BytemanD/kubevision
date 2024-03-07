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

type ConfigMaps struct{}

func (h *ConfigMaps) List(ctx context.Context, apiReq *apiv1.ConfigMapsListReq) (res *apiv1.ConfigMapsListRes, err error) {
	req := g.RequestFromCtx(ctx)
	client, err := k8s.GetClient()

	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}

	cms, err := client.ListConfigMaps(utility.GetReqNamespace(req))
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}
	data, _ := json.Marshal(map[string][]model.ConfigMap{"configmaps": cms})
	req.Response.WriteJson(data)
	return
}
