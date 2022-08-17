package handler

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type LogStructName interface {
	GetStructName() string
}

type BaseQueryHandler struct {
}

func NewBaseQueryHandler() *BaseQueryHandler {
	return &BaseQueryHandler{}
}

func (h *BaseQueryHandler) DoSession(ctx context.Context, logStruct LogStructName, event ddd.Event, fun func(ctx context.Context) error) error {
	return applog.DoEventLog(ctx, logStruct.GetStructName(), event, applog.RunFuncName(1), func() error {
		return session.StartSession(ctx, func(ctx context.Context) error {
			return fun(ctx)
		})
	})
}

func (h *BaseQueryHandler) Do(ctx context.Context, logStruct LogStructName, event ddd.Event, fun func(ctx context.Context) error) error {
	return applog.DoEventLog(ctx, logStruct.GetStructName(), event, applog.RunFuncName(1), func() error {
		return fun(ctx)
	})
}
