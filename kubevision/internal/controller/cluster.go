package controller

import (
	"context"
	"encoding/json"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/gogf/gf/v2/frame/g"
)

type Cluster struct{}

func (c *Cluster) Get(ctx context.Context, apiReq *apiv1.ClusterGetReq) (res *apiv1.ClusterGetRes, err error) {
	req := g.RequestFromCtx(ctx)
	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("load client failed: %s", err)
		req.Response.WriteStatusExit(400, "load client failed")
	}
	clusterInfo, err := client.GetClusterInfo()
	if err != nil {
		logging.Error("get cluster info failed: %s", err)
		req.Response.WriteStatusExit(400, "get cluster info failed")
	}
	data, _ := json.Marshal(map[string]model.Cluster{"cluster": *clusterInfo})
	req.Response.WriteJson(data)
	return
}
