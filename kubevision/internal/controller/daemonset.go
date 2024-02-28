package controller

import (
	"encoding/json"
	"fmt"
	"kubevision/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Daemonsets struct{}

func (c *Daemonsets) Get(req *ghttp.Request) {
	namespace := getReqNamespace(req)
	client, err := NewClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	daemonsets, err := client.ListDaemonsets(namespace)
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Daemonset{"daemonsets": daemonsets})
	req.Response.WriteJson(data)
}
