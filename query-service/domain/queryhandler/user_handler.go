package queryhandler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type userQueryHandler struct {
	service *queryservice.UserQueryService
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &userQueryHandler{
		service: queryservice_impl.NewUserQueryService(),
	}
}

func (u *userQueryHandler) OnUserCreateEventV1_0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	user := projection.UserView{
		Id:       event.UserId,
		TenantId: event.TenantId,
		UserName: event.UserName,
	}
	return u.service.Create(ctx, &user)
}

func (u *userQueryHandler) OnUserCreateEventV2_0(ctx context.Context, event *user_events.UserCreateEventV2) error {
	user := projection.UserView{
		TenantId: event.TenantId,
		Id:       event.Data.Id,
		UserName: event.Data.UserName,
	}
	return u.service.Create(ctx, &user)
}

func (u *userQueryHandler) OnUserUpdateEventV1_0(ctx context.Context, event *user_events.UserUpdateEvent) error {
	user := projection.UserView{
		TenantId: event.TenantId,
		Id:       event.Id,
	}
	return u.service.Update(ctx, &user)
}

func (u userQueryHandler) OnUserDeleteEventV1_0(ctx context.Context, event *user_events.UserDeleteEvent) error {
	user := projection.UserView{
		TenantId: event.TenantId,
		Id:       event.Id,
	}
	return u.service.Update(ctx, &user)
}
