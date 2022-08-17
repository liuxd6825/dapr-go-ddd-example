package handler

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/factory"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/application/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type InventoryQueryHandler struct {
	service *service.InventoryQueryAppService
	handler.BaseQueryHandler
}

//
// NewInventorySubscribe
// @Description: 创建dapr消息订阅器，用于接受领域事件
// @return restapp.RegisterSubscribe  消息注册器
//
func NewInventorySubscribe() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event.InventoryCreateEventType.String(), Route: "/dapr-ddd-demo/domain-event/inventory/inventory_create_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.InventoryUpdateEventType.String(), Route: "/dapr-ddd-demo/domain-event/inventory/inventory_update_event/ver:v1.0"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewInventoryQueryHandler())
}

//
// NewInventoryQueryHandler
// @Description: 创建<no value>领域事件处理器
// @return ddd.QueryEventHandler 领域事件处理器
//
func NewInventoryQueryHandler() ddd.QueryEventHandler {
	return &InventoryQueryHandler{
		service: service.GetInventoryQueryAppService(),
	}
}

//
// OnInventoryCreateEvent
// @Description: InventoryCreateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event InventoryCreateEvent 领域事件
// @return error 错误
//
func (h *InventoryQueryHandler) OnInventoryCreateEvent(ctx context.Context, event *event.InventoryCreateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.InventoryView.NewByInventoryCreateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Create(ctx, v)
	})
}

//
// OnInventoryUpdateEvent
// @Description: InventoryUpdateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event InventoryUpdateEvent 领域事件
// @return error 错误
//
func (h *InventoryQueryHandler) OnInventoryUpdateEvent(ctx context.Context, event *event.InventoryUpdateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.InventoryView.NewByInventoryUpdateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Update(ctx, v)
	})
}

func (h *InventoryQueryHandler) GetStructName() string {
	return "dapr-ddd-demo.inventory.InventoryQueryHandler"
}
