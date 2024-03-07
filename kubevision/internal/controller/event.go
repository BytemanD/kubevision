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

type Events struct{}

func (c *Events) Get(ctx context.Context, apiReq *apiv1.EventsListReq) (res *apiv1.EventsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}

	events, err := client.ListEvents(namespace)
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Event{"events": events})
	req.Response.WriteJson(data)
	return
}
