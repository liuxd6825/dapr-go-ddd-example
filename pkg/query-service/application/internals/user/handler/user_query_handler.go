package handler

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/factory"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/application/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserQueryHandler struct {
	service *service.UserQueryAppService
	handler.BaseQueryHandler
}

//
// NewUserSubscribe
// @Description: 创建dapr消息订阅器，用于接受领域事件
// @return restapp.RegisterSubscribe  消息注册器
//
func NewUserSubscribe() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event.UserCreateEventType.String(), Route: "/dapr-ddd-demo/domain-event/user/user_create_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.UserDeleteEventType.String(), Route: "/dapr-ddd-demo/domain-event/user/user_delete_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.UserUpdateEventType.String(), Route: "/dapr-ddd-demo/domain-event/user/user_update_event/ver:v1.0"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewUserQueryHandler())
}

//
// NewUserQueryHandler
// @Description: 创建<no value>领域事件处理器
// @return ddd.QueryEventHandler 领域事件处理器
//
func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		service: service.GetUserQueryAppService(),
	}
}

//
// OnUserCreateEvent
// @Description: UserCreateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event UserCreateEvent 领域事件
// @return error 错误
//
func (h *UserQueryHandler) OnUserCreateEvent(ctx context.Context, event *event.UserCreateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.UserView.NewByUserCreateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Create(ctx, v)
	})
}

//
// OnUserDeleteEvent
// @Description: UserDeleteEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event UserDeleteEvent 领域事件
// @return error 错误
//
func (h *UserQueryHandler) OnUserDeleteEvent(ctx context.Context, event *event.UserDeleteEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		return h.service.DeleteById(ctx, event.GetTenantId(), event.Data.Id)
	})
}

//
// OnUserUpdateEvent
// @Description: UserUpdateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event UserUpdateEvent 领域事件
// @return error 错误
//
func (h *UserQueryHandler) OnUserUpdateEvent(ctx context.Context, event *event.UserUpdateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.UserView.NewByUserUpdateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Update(ctx, v)
	})
}

func (h *UserQueryHandler) GetStructName() string {
	return "dapr-ddd-demo.user.UserQueryHandler"
}