package controller

import (
	"encoding/json"
	"fmt"
	"kubevision/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Pods struct{}

func (c *Pods) Get(req *ghttp.Request) {
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
}
