package apiv1

import "github.com/gogf/gf/v2/frame/g"

type JobsListReq struct {
	g.Meta `path:"/jobs" tags:"Jobs" method:"get"`
}
type JobsListRes struct {
	g.Meta `mime:"application/json" example:"{\"jobs\":[]}"`
}
type JobsPostReq struct {
	g.Meta `path:"/jobs" tags:"Jobs" method:"post"`
}
type JobsPostRes struct {
	g.Meta `mime:"application/json"`
}
type JobsDeleteReq struct {
	g.Meta `path:"/jobs" tags:"Jobs" method:"delete"`
}
type JobsDeleteRes struct {
	g.Meta `mime:"application/json"`
}
