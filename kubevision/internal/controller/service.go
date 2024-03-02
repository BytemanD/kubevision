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

type Services struct{}

func (c *Services) Get(ctx context.Context, apiReq *apiv1.ServicesListReq) (res *apiv1.ServicesListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := getReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	services, err := client.ListServices(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Service{"services": services})
	req.Response.WriteJson(data)
	return
}
