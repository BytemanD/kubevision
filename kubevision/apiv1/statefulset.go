package apiv1

import "github.com/gogf/gf/v2/frame/g"

type StatefulSetsListReq struct {
	g.Meta `path:"/statefulsets" tags:"StatefulSets" method:"get"`
}
type StatefulSetsListRes struct {
	g.Meta `mime:"application/json" example:"{\"statefulsets\":[]}"`
}
type StatefulSetsPostReq struct {
	g.Meta `path:"/statefulsets" tags:"StatefulSets" method:"post"`
}
type StatefulSetsPostRes struct {
	g.Meta `mime:"application/json"`
}
type StatefulSetsDeleteReq struct {
	g.Meta `path:"/statefulsets" tags:"StatefulSets" method:"delete"`
}
type StatefulSetsDeleteRes struct {
	g.Meta `mime:"application/json"`
}
