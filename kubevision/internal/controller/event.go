package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"

	"github.com/gogf/gf/v2/frame/g"
)

type Events struct{}

func (c *Events) Get(ctx context.Context, apiReq *apiv1.EventsListReq) (res *apiv1.EventsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := getReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	events, err := client.ListEvents(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Event{"events": events})
	req.Response.WriteJson(data)
	return
}
