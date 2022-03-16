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

func (u *UserQueryHandler) OnUserCreateEventV1_0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	user := factory.NewUserView_UserCreateEventV1(event)
	return u.service.Create(ctx, user)
}

func (u *UserQueryHandler) OnUserCreateEventV2_0(ctx context.Context, event *user_events.UserCreateEventV2) error {
	user := factory.NewUserView_UserCreateEventV2(event)
	return u.service.Create(ctx, user)
}

func (u *UserQueryHandler) OnUserUpdateEventV1_0(ctx context.Context, event *user_events.UserUpdateEvent) error {
	user := factory.NewUserView_UserUpdateEventV1(event)
	return u.service.Update(ctx, user)
}

func (u UserQueryHandler) OnUserDeleteEventV1_0(ctx context.Context, event *user_events.UserDeleteEvent) error {
	return u.service.DeleteById(ctx, event.TenantId, event.EventId)
}
