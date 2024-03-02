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

type CronJobs struct{}

func (c *CronJobs) Get(ctx context.Context, apiReq *apiv1.CronJobsListReq) (res *apiv1.CronJobsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := getReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	cronjobs, err := client.ListCronJobs(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.CronJob{"cronjobs": cronjobs})
	req.Response.WriteJson(data)
	return
}
