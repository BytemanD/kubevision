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

type StatefulSets struct{}

func (c *StatefulSets) Get(ctx context.Context, apiReq *apiv1.StatefulSetsListReq) (res *apiv1.StatefulSetsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := getReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	statefulsets, err := client.ListStatefulSets(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.StatefulSet{"statefulsets": statefulsets})
	req.Response.WriteJson(data)
	return
}
