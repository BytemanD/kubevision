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

type PodGetReq struct {
	g.Meta `path:"/pods/:name" tags:"Pods" method:"get"`
}
type PodGetRes struct {
	g.Meta `mime:"application/json"`
}
type PodDeleteReq struct {
	g.Meta `path:"/pods/:name" tags:"Pods" method:"delete"`
}
type PodDeleteRes struct {
	g.Meta `mime:"application/json"`
}
type PodDescribeReq struct {
	g.Meta `path:"/pods/:name/describe" tags:"Pods" method:"get"`
}
type PodDescribeRes struct {
	g.Meta `mime:"application/text"`
}
