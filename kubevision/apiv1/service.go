package apiv1

import "github.com/gogf/gf/v2/frame/g"

type ServicesListReq struct {
	g.Meta `path:"/services" tags:"Services" method:"get"`
}
type ServicesListRes struct {
	g.Meta `mime:"application/json" example:"{\"services\":[]}"`
}
type ServicesPostReq struct {
	g.Meta `path:"/services" tags:"Services" method:"post"`
}
type ServicesPostRes struct {
	g.Meta `mime:"application/json"`
}
type ServicesDeleteReq struct {
	g.Meta `path:"/services" tags:"Services" method:"delete"`
}
type ServicesDeleteRes struct {
	g.Meta `mime:"application/json"`
}
