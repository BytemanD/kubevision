package controller

import (
	"encoding/json"
	"fmt"
	"kubevision/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Nodes struct{}

func (c *Nodes) Get(req *ghttp.Request) {
	client, err := NewClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	nodes, err := client.ListNodes()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Node{"nodes": nodes})
	req.Response.WriteJson(data)
}
