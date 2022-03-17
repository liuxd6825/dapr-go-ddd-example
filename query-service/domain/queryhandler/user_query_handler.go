package queryhandler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/factory"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserQueryHandler struct {
	service *queryservice.UserQueryService
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		service: queryservice_impl.NewUserQueryService(),
	}
}

func (u *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	user := factory.NewUserView_UserCreateEventV1(event)
	return u.service.Create(ctx, user)
}

func (u *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	user := factory.NewUserView_UserUpdateEventV1(event)
	return u.service.Update(ctx, user)
}

func (u UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *user_events.UserDeleteEventV1) error {
	return u.service.DeleteById(ctx, event.TenantId, event.EventId)
}
