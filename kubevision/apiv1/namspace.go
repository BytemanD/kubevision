package apiv1

import "github.com/gogf/gf/v2/frame/g"

type NamespacesListReq struct {
	g.Meta `path:"/namespaces" tags:"Namespaces" method:"get"`
}
type NamespacesListRes struct {
	g.Meta `mime:"application/json" example:"{\"namespaces\":[]}"`
}
type NamespacesPostReq struct {
	g.Meta `path:"/namespaces" tags:"Namespaces" method:"post"`
}
type NamespacesPostRes struct {
	g.Meta `mime:"application/json"`
}
