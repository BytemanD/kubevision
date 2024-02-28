package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"kubevision/apiv1"
	"kubevision/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type Pods struct{}

func (c *Pods) Get(ctx context.Context, apiReq *apiv1.PodsListReq) (res *apiv1.PodsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := getReqNamespace(req)
	client, err := NewClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	pods, err := client.ListPods(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Pod{"pods": pods})
	req.Response.WriteJson(data)
	return
}
