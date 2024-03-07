package utility

import "github.com/gogf/gf/v2/net/ghttp"

func GetReqNamespace(req *ghttp.Request) string {
	namespace := req.URL.Query().Get("namespace")
	if namespace != "" {
		return namespace
	} else {
		return "default"
	}
}
