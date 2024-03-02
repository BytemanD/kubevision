package apiv1

import "github.com/gogf/gf/v2/frame/g"

type EventsListReq struct {
	g.Meta `path:"/events" tags:"Events" method:"get"`
}
type EventsListRes struct {
	g.Meta `mime:"application/json" example:"{\"events\":[]}"`
}
type EventsPostReq struct {
	g.Meta `path:"/events" tags:"Events" method:"post"`
}
type EventsPostRes struct {
	g.Meta `mime:"application/json"`
}
type EventsDeleteReq struct {
	g.Meta `path:"/events" tags:"Events" method:"delete"`
}
type EventsDeleteRes struct {
	g.Meta `mime:"application/json"`
}
