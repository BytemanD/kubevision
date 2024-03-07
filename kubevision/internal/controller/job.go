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

type Jobs struct{}

func (c *Jobs) Get(ctx context.Context, apiReq *apiv1.JobsListReq) (res *apiv1.JobsListRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}

	jobs, err := client.ListJobs(namespace)
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Job{"jobs": jobs})
	req.Response.WriteJson(data)
	return
}
