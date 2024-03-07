package controller

import (
	"context"
	"encoding/json"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"
	"kubevision/utility"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/gogf/gf/v2/frame/g"
)

type Pods struct{}

func (c *Pods) Get(ctx context.Context, apiReq *apiv1.PodsListReq) (res *apiv1.PodsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}

	pods, err := client.ListPods(namespace)
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Pod{"pods": pods})
	req.Response.WriteJson(data)
	return
}
