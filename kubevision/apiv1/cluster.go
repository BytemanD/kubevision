package apiv1

import "github.com/gogf/gf/v2/frame/g"

type ClusterGetReq struct {
	g.Meta `path:"/cluster" tags:"configmaps" method:"get"`
}
type ClusterGetRes struct {
	g.Meta `mime:"application/json" example:"{}"`
}
