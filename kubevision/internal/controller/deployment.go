package controller

import (
	"encoding/json"
	"fmt"
	"kubevision/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Deployments struct{}

func (c *Deployments) Get(req *ghttp.Request) {
	namespace := getReqNamespace(req)
	client, err := NewClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	deployments, err := client.ListDeployments(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Deployment{"deployments": deployments})
	req.Response.WriteJson(data)
}
