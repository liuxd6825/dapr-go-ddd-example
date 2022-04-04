package queryhandler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/factory"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserQueryHandler struct {
	service *queryservice.UserQueryService
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		service: queryservice.NewUserQueryService(),
	}
}

func (u *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	return applog.DoEvent(u, event, "OnUserCreateEventV1s0", func() error {
		user := factory.NewUserViewByUserCreateEventV1(event)
		return u.service.Create(ctx, user)
	})
}

func (u *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	return applog.DoEvent(u, event, "OnUserUpdateEventV1s0", func() error {
		user := factory.NewUserViewByUserUpdateEventV1(event)
		return u.service.Create(ctx, user)
	})
}

func (u *UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *user_events.UserDeleteEventV1) error {
	return applog.DoEvent(u, event, "OnUserDeleteEventV1s0", func() error {
		return u.service.DeleteById(ctx, event.TenantId, event.EventId)
	})

}

func (u *UserQueryHandler) GetClassName() string {
	return "UserQueryHandler"
}
