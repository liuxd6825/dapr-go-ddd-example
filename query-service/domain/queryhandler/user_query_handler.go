package queryhandler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserQueryHandler struct {
	userService *queryservice.UserQueryService
	addrService *queryservice.AddressQueryService
	restapp.BaseQueryHandler
}

func NewUserSubscribes() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
		{PubsubName: "pubsub", Topic: user_events.UserCreateEventType.String(), Route: "/users/user-create-event"},
		{PubsubName: "pubsub", Topic: user_events.UserUpdateEventType.String(), Route: "/users/user-update-event"},
		{PubsubName: "pubsub", Topic: user_events.UserDeleteEventType.String(), Route: "/users/user-delete-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressCreateEventType.String(), Route: "/users/user-address-create-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressUpdateEventType.String(), Route: "/users/user-address-update-event"},
		{PubsubName: "pubsub", Topic: user_events.AddressDeleteEventType.String(), Route: "/users/user-address-delete-event"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewUserQueryHandler())
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		userService: queryservice.NewUserQueryService(),
		addrService: queryservice.NewAddressQueryService(),
	}
}

func (h *UserQueryHandler) OnUserAddressCreateEventV1s0(ctx context.Context, event *user_events.AddressCreateEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		addr := user_factory.NewAddressViewByUserCreateEventV1(event)
		return h.addrService.Create(ctx, addr)
	})
}

func (h *UserQueryHandler) OnUserAddressUpdateEventV1s0(ctx context.Context, event *user_events.AddressUpdateEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		addr := user_factory.NewAddressViewByUserUpdateEventV1(event)
		return h.addrService.Create(ctx, addr)
	})
}

func (h *UserQueryHandler) OnUserAddressDeleteEventV1s0(ctx context.Context, event *user_events.AddressDeleteEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		return h.addrService.DeleteById(ctx, event.TenantId, event.AddressId)
	})
}

func (h *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		user := user_factory.NewUserViewByUserCreateEventV1(event)
		return h.userService.Create(ctx, user)
	})
}

func (h *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(dbCtx context.Context) error {
		user := user_factory.NewUserViewByUserUpdateEventV1(event)
		return h.userService.Create(dbCtx, user)
	})
}

func (h *UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *user_events.UserDeleteEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		return h.userService.DeleteById(ctx, event.TenantId, event.EventId)
	})
}

func (h *UserQueryHandler) GetStructName() string {
	return "query-service.UserQueryHandler"
}
