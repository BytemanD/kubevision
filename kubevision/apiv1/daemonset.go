package apiv1

import "github.com/gogf/gf/v2/frame/g"

type DaemonsetsListReq struct {
	g.Meta `path:"/daemonsets" tags:"daemonsets" method:"get"`
}
type DaemonsetsListRes struct {
	g.Meta `mime:"application/json" example:"{\"daemonsets\":[]}"`
}
type DaemonsetsPostReq struct {
	g.Meta `path:"/daemonsets" tags:"daemonsets" method:"post"`
}
type DaemonsetsPostRes struct {
	g.Meta `mime:"application/json"`
}
