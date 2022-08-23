package handler

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/logs"
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
		{PubsubName: "pubsub", Topic: event.SaleItemDeleteEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_item_delete_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.SaleItemCreateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_item_create_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.SaleItemUpdateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_item_update_event/ver:v1.0"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewSaleItemQueryHandler())
}

func NewSaleItemQueryHandler() ddd.QueryEventHandler {
	return &SaleItemQueryHandler{
		service: service.GetSaleItemQueryAppService(),
	}
}

//
// OnSaleItemDeleteEventV1s0
// @Description: SaleItemDeleteEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleItemDeleteEvent 领域事件
// @return error 错误
//
func (h *SaleItemQueryHandler) OnSaleItemDeleteEventV1s0(ctx context.Context, event *event.SaleItemDeleteEvent) error {
	logs.Debugln(event)
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleItemView.NewBySaleItemDeleteEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.DeleteMany(ctx, event.GetTenantId(), v)
	})
}

//
// OnSaleItemCreateEventV1s0
// @Description: SaleItemCreateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleItemCreateEvent 领域事件
// @return error 错误
//
func (h *SaleItemQueryHandler) OnSaleItemCreateEventV1s0(ctx context.Context, event *event.SaleItemCreateEvent) error {
	logs.DebugEvent(event, "OnSaleItemCreateEventV1s0")

	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleItemView.NewBySaleItemCreateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.CreateMany(ctx, v)
	})
}

//
// OnSaleItemUpdateEventV1s0
// @Description: SaleItemUpdateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleItemUpdateEvent 领域事件
// @return error 错误
//
func (h *SaleItemQueryHandler) OnSaleItemUpdateEventV1s0(ctx context.Context, event *event.SaleItemUpdateEvent) error {

	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleItemView.NewBySaleItemUpdateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.UpdateMany(ctx, v)
	})
}

func (h *SaleItemQueryHandler) GetStructName() string {
	return "dapr-ddd-demo.sale_bill.SaleItemQueryHandler"
}
