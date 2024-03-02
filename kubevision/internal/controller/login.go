package controller

import (
	"context"
	"kubevision/apiv1"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/guid"
)

type Login struct{}

type Token struct {
	UUID string
}

var tokens []Token = []Token{}

func (c *Login) Post(ctx context.Context, apiReq *apiv1.LoginPostReq) (res *apiv1.LoginPostRes, err error) {
	req := g.RequestFromCtx(ctx)
	uuid := guid.S()
	req.Response.Header().Set("X-Token", uuid)
	tokens = append(tokens, Token{UUID: uuid})
	req.Response.WriteJson(nil)
	return
}
func (c *Login) Get(ctx context.Context, apiReq *apiv1.LoginGetReq) (res *apiv1.LoginGetRes, err error) {
	req := g.RequestFromCtx(ctx)
	token := req.Header.Get("X-Token")
	if token == "" {
		req.Response.WriteStatusExit(403, "not auth")
	}
	req.Response.WriteJson(nil)
	return
}
