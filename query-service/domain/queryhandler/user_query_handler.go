package queryhandler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserQueryHandler struct {
	userService *queryservice.UserQueryService
	addrService *queryservice.AddressQueryService
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		userService: queryservice.NewUserQueryService(),
	}
}

func (u *UserQueryHandler) OnAddressCreateEventV1s0(ctx context.Context, event *user_events.AddressCreateEventV1) error {
	return applog.DoEventLog(ctx, u, event, "OnAddressCreateEventV1s0", func() error {
		addr := user_factory.NewAddressViewByUserCreateEventV1(event)
		return u.addrService.Create(ctx, addr)
	})
}

func (u *UserQueryHandler) OnAddressUpdateEventV1s0(ctx context.Context, event *user_events.AddressUpdateEventV1) error {
	return applog.DoEventLog(ctx, u, event, "OnAddressUpdateEventV1s0", func() error {
		addr := user_factory.NewAddressViewByUserUpdateEventV1(event)
		return u.addrService.Create(ctx, addr)
	})
}

func (u *UserQueryHandler) OnAddressDeleteEventV1s0(ctx context.Context, event *user_events.AddressDeleteEventV1) error {
	return applog.DoEventLog(ctx, u, event, "OnAddressDeleteEventV1s0", func() error {
		return u.addrService.DeleteById(ctx, event.TenantId, event.AddressId)
	})
}

func (u *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	return applog.DoEventLog(ctx, u, event, "OnUserCreateEventV1s0", func() error {
		user := user_factory.NewUserViewByUserCreateEventV1(event)
		return u.userService.Create(ctx, user)
	})
}

func (u *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	return applog.DoEventLog(ctx, u, event, "OnUserUpdateEventV1s0", func() error {
		user := user_factory.NewUserViewByUserUpdateEventV1(event)
		return u.userService.Create(ctx, user)
	})
}

func (u *UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *user_events.UserDeleteEventV1) error {
	return applog.DoEventLog(ctx, u, event, "OnUserDeleteEventV1s0", func() error {
		return u.userService.DeleteById(ctx, event.TenantId, event.EventId)
	})

}

func (u *UserQueryHandler) GetClassName() string {
	return "UserQueryHandler"
}
