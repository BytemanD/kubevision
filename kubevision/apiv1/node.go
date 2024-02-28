package apiv1

import "github.com/gogf/gf/v2/frame/g"

type NodesListReq struct {
	g.Meta `path:"/nodes" tags:"Nodes" method:"get"`
}
type NodesListRes struct {
	g.Meta `mime:"application/json" example:"{\"nodes\":[]}"`
}
type NodesDeleteReq struct {
	g.Meta `path:"/nodes" tags:"Nodes" method:"delete"`
}
type NodesDeleteRes struct {
	g.Meta `mime:"application/json"`
}
