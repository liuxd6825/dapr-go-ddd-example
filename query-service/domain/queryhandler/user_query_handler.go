package queryhandler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/db"
	"github.com/liuxd6825/dapr-go-ddd-sdk/applog"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/rest"
)

type UserQueryHandler struct {
	userService *queryservice.UserQueryService
	addrService *queryservice.AddressQueryService
}

func NewUserSubscribeHandler() *rest.RegisterSubscribe {
	subscribes := []ddd.Subscribe{
		{PubsubName: "pubsub", Topic: user_events.UserCreateEventType.String(), Route: "/users/user-create-event"},
		{PubsubName: "pubsub", Topic: user_events.UserUpdateEventType.String(), Route: "/users/user-update-event"},
		{PubsubName: "pubsub", Topic: user_events.UserDeleteEventType.String(), Route: "/users/user-delete-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressCreateEventType.String(), Route: "/users/user-address-create-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressUpdateEventType.String(), Route: "/users/user-address-update-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressDeleteEventType.String(), Route: "/users/user-address-delete-event"},
	}
	return &rest.RegisterSubscribe{
		Subscribes: subscribes,
		Handler:    NewUserQueryHandler(),
	}
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		userService: queryservice.NewUserQueryService(),
		addrService: queryservice.NewAddressQueryService(),
	}
}

func (h *UserQueryHandler) OnUserAddressCreateEventV1s0(ctx context.Context, event *user_events.AddressCreateEventV1) error {
	return h.doSession(ctx, event, func(ctx context.Context) error {
		addr := user_factory.NewAddressViewByUserCreateEventV1(event)
		return h.addrService.Create(ctx, addr)
	})
}

func (h *UserQueryHandler) OnUserAddressUpdateEventV1s0(ctx context.Context, event *user_events.AddressUpdateEventV1) error {
	return h.doSession(ctx, event, func(ctx context.Context) error {
		addr := user_factory.NewAddressViewByUserUpdateEventV1(event)
		return h.addrService.Create(ctx, addr)
	})
}

func (h *UserQueryHandler) OnUserAddressDeleteEventV1s0(ctx context.Context, event *user_events.AddressDeleteEventV1) error {
	return h.doSession(ctx, event, func(ctx context.Context) error {
		return h.addrService.DeleteById(ctx, event.TenantId, event.AddressId)
	})
}

func (h *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	return h.doSession(ctx, event, func(ctx context.Context) error {
		user := user_factory.NewUserViewByUserCreateEventV1(event)
		return h.userService.Create(ctx, user)
	})
}

func (h *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	return h.doSession(ctx, event, func(ctx context.Context) error {
		user := user_factory.NewUserViewByUserUpdateEventV1(event)
		return h.userService.Create(ctx, user)
	})
}

func (h *UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *user_events.UserDeleteEventV1) error {
	return h.doSession(ctx, event, func(ctx context.Context) error {
		return h.userService.DeleteById(ctx, event.TenantId, event.EventId)
	})
}

func (h *UserQueryHandler) doSession(ctx context.Context, event ddd.Event, fun func(ctx context.Context) error) error {
	return applog.DoEventLog(ctx, h, event, applog.RunFuncName(1), func() error {
		return ddd_repository.StartSession(ctx, db.NewSession(), func(ctx context.Context) error {
			return fun(ctx)
		})
	})
}

func (h *UserQueryHandler) doNonSession(ctx context.Context, event ddd.Event, fun func(ctx context.Context) error) error {
	return applog.DoEventLog(ctx, h, event, applog.RunFuncName(1), func() error {
		return fun(ctx)
	})
}

func (h *UserQueryHandler) GetClassName() string {
	return "UserQueryHandler"
}
