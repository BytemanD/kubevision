package apiv1

import "github.com/gogf/gf/v2/frame/g"

type DeploymentsListReq struct {
	g.Meta `path:"/deployments" tags:"deployments" method:"get"`
}
type DeploymentsListRes struct {
	g.Meta `mime:"application/json" example:"{\"deployments\":[]}"`
}
type DeploymentsPostReq struct {
	g.Meta `path:"/deployments" tags:"deployments" method:"post"`
}
type DeploymentsPostRes struct {
	g.Meta `mime:"application/json"`
}
