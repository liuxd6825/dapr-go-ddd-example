package handler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/application/internals/service"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserQueryHandler struct {
	userService *service.UserQueryAppService
	restapp.BaseQueryHandler
}

func NewUserSubscribes() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event.UserCreateEventType.String(), Route: "/event/command-service/users/user-create-event"},
		{PubsubName: "pubsub", Topic: event.UserUpdateEventType.String(), Route: "/event/command-service/users/user-update-event"},
		{PubsubName: "pubsub", Topic: event.UserDeleteEventType.String(), Route: "/event/command-service/users/user-delete-event"},
		{PubsubName: "pubsub", Topic: event.AddressCreateEventType.String(), Route: "/event/command-service/users/user-address-create-event"},
		{PubsubName: "pubsub", Topic: event.AddressUpdateEventType.String(), Route: "/event/command-service/users/user-address-update-event"},
		{PubsubName: "pubsub", Topic: event.AddressDeleteEventType.String(), Route: "/event/command-service/users/user-address-delete-event"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewUserQueryHandler())
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		userService: service.GetUserQueryAppService(),
	}
}

//
// OnUserCreateEventV1s0
// @Description: 事件处理器
// @receiver h
// @param ctx 上下文
// @param event 领域事件
// @return error 错误
//
func (h *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *event.UserCreateEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		user := user_factory.NewUserViewByUserCreateEventV1(event)
		return h.userService.Create(ctx, user)
	})
}

func (h *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *event.UserUpdateEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(dbCtx context.Context) error {
		user := user_factory.NewUserViewByUserUpdateEventV1(event)
		return h.userService.Update(dbCtx, user)
	})
}

func (h *UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *event.UserDeleteEventV1) error {
	return h.DoSession(ctx, h.GetStructName, event, func(ctx context.Context) error {
		return h.userService.DeleteById(ctx, event.TenantId, event.EventId)
	})
}

/*
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
*/
func (h *UserQueryHandler) GetStructName() string {
	return "query-service.UserQueryHandler"
}
