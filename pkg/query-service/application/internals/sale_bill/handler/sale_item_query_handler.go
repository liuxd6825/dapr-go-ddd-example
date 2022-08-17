package handler

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/factory"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/application/handler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type SaleItemQueryHandler struct {
	service *service.SaleItemQueryAppService
	handler.BaseQueryHandler
}

func NewSaleItemSubscribe() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event.SaleItemCreateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_item_create_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.SaleItemUpdateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_item_update_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.SaleItemDeleteEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_item_delete_event/ver:v1.0"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewSaleItemQueryHandler())
}

func NewSaleItemQueryHandler() ddd.QueryEventHandler {
	return &SaleItemQueryHandler{
		service: service.GetSaleItemQueryAppService(),
	}
}

//
// OnSaleItemCreateEvent
// @Description: SaleItemCreateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleItemCreateEvent 领域事件
// @return error 错误
//
func (h *SaleItemQueryHandler) OnSaleItemCreateEvent(ctx context.Context, event *event.SaleItemCreateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleItemView.NewBySaleItemCreateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Create(ctx, v)
	})
}

//
// OnSaleItemUpdateEvent
// @Description: SaleItemUpdateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleItemUpdateEvent 领域事件
// @return error 错误
//
func (h *SaleItemQueryHandler) OnSaleItemUpdateEvent(ctx context.Context, event *event.SaleItemUpdateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleItemView.NewBySaleItemUpdateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Update(ctx, v)
	})
}

//
// OnSaleItemDeleteEvent
// @Description: SaleItemDeleteEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleItemDeleteEvent 领域事件
// @return error 错误
//
func (h *SaleItemQueryHandler) OnSaleItemDeleteEvent(ctx context.Context, event *event.SaleItemDeleteEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		return h.service.DeleteById(ctx, event.GetTenantId(), event.Data.Id)
	})
}

func (h *SaleItemQueryHandler) GetStructName() string {
	return "dapr-ddd-demo.sale_bill.SaleItemQueryHandler"
}
