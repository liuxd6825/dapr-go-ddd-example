package controller

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_context"
)

type serverContext struct {
	ctx iris.Context
}

func NewContext(irisCtx iris.Context) context.Context {
	metadata := make(map[string]string, 0)
	header := irisCtx.Request().Header
	for k, v := range header {
		metadata[k] = v[0]
	}
	serverCtx := newServerContext(irisCtx)
	return ddd_context.NewContext(context.Background(), metadata, serverCtx)
}

func newServerContext(ctx iris.Context) ddd_context.ServerContext {
	return &serverContext{
		ctx: ctx,
	}
}

func (s *serverContext) SetResponseHeader(name string, value string) {
	s.ctx.Header(name, value)
}
