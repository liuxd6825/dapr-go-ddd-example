package handler

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/logs"
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

// NewUserSubscribe
// @Description: 创建dapr消息订阅器，用于接受领域事件
// @return restapp.RegisterSubscribe  消息注册器
func NewUserSubscribe() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event.UserCreateEventType.String(), Route: "/dapr-ddd-demo/domain-event/user/user_create_event"},
		{PubsubName: "pubsub", Topic: event.UserDeleteEventType.String(), Route: "/dapr-ddd-demo/domain-event/user/user_delete_event"},
		{PubsubName: "pubsub", Topic: event.UserUpdateEventType.String(), Route: "/dapr-ddd-demo/domain-event/user/user_update_event"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewUserQueryHandler())
}

// NewUserQueryHandler
// @Description: 创建<no value>领域事件处理器
// @return ddd.QueryEventHandler 领域事件处理器
func NewUserQueryHandler() ddd.QueryEventHandler {
	return &UserQueryHandler{
		service: service.GetUserQueryAppService(),
	}
}

// OnUserCreateEventV1s0
// @Description: UserCreateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event UserCreateEvent 领域事件
// @return error 错误
func (h *UserQueryHandler) OnUserCreateEventV1s0(ctx context.Context, event *event.UserCreateEvent) error {
	logs.DebugEvent(event, "OnOnUserCreateEventV1s0")
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.UserView.NewByUserCreateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Create(ctx, v)
	})
}

// OnUserDeleteEventV1s0
// @Description: UserDeleteEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event UserDeleteEvent 领域事件
// @return error 错误
func (h *UserQueryHandler) OnUserDeleteEventV1s0(ctx context.Context, event *event.UserDeleteEvent) error {
	logs.DebugEvent(event, "OnOnUserDeleteEventV1s0")
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		return h.service.DeleteById(ctx, event.GetTenantId(), event.Data.Id)
	})
}

// OnUserUpdateEventV1s0
// @Description: UserUpdateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event UserUpdateEvent 领域事件
// @return error 错误
func (h *UserQueryHandler) OnUserUpdateEventV1s0(ctx context.Context, event *event.UserUpdateEvent) error {
	logs.DebugEvent(event, "OnOnUserUpdateEventV1s0")
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
