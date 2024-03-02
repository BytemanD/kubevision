package apiv1

import "github.com/gogf/gf/v2/frame/g"

type SecretsListReq struct {
	g.Meta `path:"/secrets" tags:"Secrets" method:"get"`
}
type SecretsListRes struct {
	g.Meta `mime:"application/json" example:"{\"secrets\":[]}"`
}
type SecretsPostReq struct {
	g.Meta `path:"/secrets" tags:"Secrets" method:"post"`
}
type SecretsPostRes struct {
	g.Meta `mime:"application/json"`
}
type SecretsDeleteReq struct {
	g.Meta `path:"/secrets" tags:"Secrets" method:"delete"`
}
type SecretsDeleteRes struct {
	g.Meta `mime:"application/json"`
}
