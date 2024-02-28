package controller

import (
	"encoding/json"
	"fmt"
	"kubevision/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Namespaces struct{}

func (c *Namespaces) Get(req *ghttp.Request) {
	client, err := NewClient()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}

	namespaces, err := client.ListNamespaces()
	if err != nil {
		fmt.Println(err)
		req.Response.WriteStatusExit(400)
	}
	data, _ := json.Marshal(map[string][]model.Namespace{"namespaces": namespaces})
	req.Response.WriteJson(data)
}

func getReqNamespace(req *ghttp.Request) string {
	namespace := req.Header.Get("X-Namespace")
	if namespace != "" {
		return namespace
	} else {
		return "default"
	}
}
