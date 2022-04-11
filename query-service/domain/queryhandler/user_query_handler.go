package queryhandler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserQueryHandler struct {
	userService *queryservice.UserQueryService
	addrService *queryservice.AddressQueryService
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		userService: queryservice.NewUserQueryService(),
		addrService: queryservice.NewAddressQueryService(),
	}
}

func (u *UserQueryHandler) OnUserAddressCreateEventV1s0(ctx context.Context, event *user_events.AddressCreateEventV1) error {
	return u.start(ctx, event, "OnUserAddressCreateEventV1s0", func(ctx context.Context) error {
		addr := user_factory.NewAddressViewByUserCreateEventV1(event)
		return u.addrService.Create(ctx, addr)
	})
}

func (u *UserQueryHandler) OnUserAddressUpdateEventV1s0(ctx context.Context, event *user_events.AddressUpdateEventV1) error {
	return u.start(ctx, event, "OnUserAddressUpdateEventV1s0", func(ctx context.Context) error {
		addr := user_factory.NewAddressViewByUserUpdateEventV1(event)
		return u.addrService.Create(ctx, addr)
	})
}

func (u *UserQueryHandler) OnUserAddressDeleteEventV1s0(ctx context.Context, event *user_events.AddressDeleteEventV1) error {
	return u.start(ctx, event, "OnUserAddressDeleteEventV1s0", func(ctx context.Context) error {
		return u.addrService.DeleteById(ctx, event.TenantId, event.AddressId)
	})
}

func (u *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	return u.start(ctx, event, "OnUserCreateEventV1s0", func(ctx context.Context) error {
		user := user_factory.NewUserViewByUserCreateEventV1(event)
		return u.userService.Create(ctx, user)
	})
}

func (u *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	return u.start(ctx, event, "OnUserUpdateEventV1s0", func(ctx context.Context) error {
		user := user_factory.NewUserViewByUserUpdateEventV1(event)
		return u.userService.Create(ctx, user)
	})
}

func (u *UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *user_events.UserDeleteEventV1) error {
	return u.start(ctx, event, "OnUserDeleteEventV1s0", func(ctx context.Context) error {
		return u.userService.DeleteById(ctx, event.TenantId, event.EventId)
	})
}

func (u *UserQueryHandler) start(ctx context.Context, event ddd.Event, funcName string, fun func(ctx context.Context) error) error {
	return applog.DoEventLog(ctx, u, event, funcName, func() error {
		return ddd_repository.StartSession(ctx, nil, func(ctx context.Context) error {
			return fun(ctx)
		})
	})
}

func (u *UserQueryHandler) GetClassName() string {
	return "UserQueryHandler"
}
