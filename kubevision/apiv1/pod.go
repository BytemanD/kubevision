package apiv1

import "github.com/gogf/gf/v2/frame/g"

type PodsListReq struct {
	g.Meta `path:"/pods" tags:"Pods" method:"get"`
}
type PodsListRes struct {
	g.Meta `mime:"application/json" example:"{\"pods\":[]}"`
}
type PodsPostReq struct {
	g.Meta `path:"/pods" tags:"Pods" method:"post"`
}
type PodsPostRes struct {
	g.Meta `mime:"application/json"`
}
type PodsDeleteReq struct {
	g.Meta `path:"/pods" tags:"Pods" method:"delete"`
}
type PodsDeleteRes struct {
	g.Meta `mime:"application/json"`
}
