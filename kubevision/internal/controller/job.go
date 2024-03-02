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

type Jobs struct{}

func (c *Jobs) Get(ctx context.Context, apiReq *apiv1.JobsListReq) (res *apiv1.JobsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := getReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	jobs, err := client.ListJobs(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Job{"jobs": jobs})
	req.Response.WriteJson(data)
	return
}
