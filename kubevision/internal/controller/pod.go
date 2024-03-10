package controller

import (
	"context"
	"encoding/json"
	"kubevision/apiv1"
	"kubevision/internal/model"
	"kubevision/internal/service/k8s"
	"kubevision/utility"
	"strings"

	"github.com/BytemanD/easygo/pkg/global/logging"
	"github.com/gogf/gf/v2/frame/g"
)

type Pods struct{}

type Pod struct{}

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

func (c *Pod) Delete(ctx context.Context, apiReq *apiv1.PodDeleteReq) (res *apiv1.PodDeleteRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	err = client.DeletePod(namespace, req.Get("name").String(), k8s.NewDeleteOption())
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	req.Response.WriteStatusExit(204)
	return
}

func (c *Pod) Describe(ctx context.Context, apiReq *apiv1.PodDescribeReq) (res *apiv1.PodDescribeRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	data, err := client.DescribePod(namespace, req.Get("name").String())
	if err != nil {
		logging.Error("describe pod failed: %v", err)
		req.Response.WriteStatusExit(400)
	}
	req.Response.WriteStatusExit(200, data)
	return
}
func (c *Pod) Execute(ctx context.Context, apiReq *apiv1.PodExecReq) (res *apiv1.PodExecRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)

	type Params struct {
		Container string `json:"container"`
		Command   string `json:"command"`
	}
	body := map[string]Params{"exec": {}}
	err = utility.GetReqBody(req, &body)
	if err != nil {
		req.Response.WriteStatusExit(400, "read request body failed")
	}

	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	stdout, stderr, err := client.ExecCommand(
		namespace, req.Get("name").String(),
		body["exec"].Container, strings.Fields(body["exec"].Command))
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}
	respBody := map[string]string{"stdout": stdout, "stderr": stderr}
	req.Response.WriteStatusExit(200, respBody)
	return
}

func (c *Pod) GetLogs(ctx context.Context, apiReq *apiv1.PodLogsReq) (res *apiv1.PodLogsRes, err error) {
	req := g.RequestFromCtx(ctx)

	namespace := utility.GetReqNamespace(req)
	container := utility.GetReqParamString(req, "container")
	lines := utility.GetReqParamInt64(req, "lines")

	client, err := k8s.GetClient()
	if err != nil {
		logging.Error("%v", err)
		req.Response.WriteStatusExit(400)
	}
	var containerName string
	if container != nil {
		containerName = *container
	}

	logs, err := client.GetLogs(
		namespace, req.Get("name").String(),
		containerName, lines,
	)
	if err != nil {
		req.Response.WriteStatusExit(400, err)
	}
	respBody := map[string]string{"logs": logs}
	req.Response.WriteStatusExit(200, respBody)
	return
}
