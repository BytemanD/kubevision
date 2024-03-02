package apiv1

import "github.com/gogf/gf/v2/frame/g"

type ConfigMapsListReq struct {
	g.Meta `path:"/configmaps" tags:"configmaps" method:"get"`
}
type ConfigMapsListRes struct {
	g.Meta `mime:"application/json" example:"{\"configmaps\":[]}"`
}
