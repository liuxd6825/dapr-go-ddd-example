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

type SaleBillQueryHandler struct {
	service *service.SaleBillQueryAppService
	handler.BaseQueryHandler
}

//
// NewSaleBillSubscribe
// @Description: 创建dapr消息订阅器，用于接受领域事件
// @return restapp.RegisterSubscribe  消息注册器
//
func NewSaleBillSubscribe() restapp.RegisterSubscribe {
	subscribes := &[]ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event.SaleBillConfirmEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_confirm_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.SaleBillCreateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_create_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.SaleBillDeleteEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_delete_event/ver:v1.0"},
		{PubsubName: "pubsub", Topic: event.SaleBillUpdateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_update_event/ver:v1.0"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewSaleBillQueryHandler())
}

//
// NewSaleBillQueryHandler
// @Description: 创建<no value>领域事件处理器
// @return ddd.QueryEventHandler 领域事件处理器
//
func NewSaleBillQueryHandler() ddd.QueryEventHandler {
	return &SaleBillQueryHandler{
		service: service.GetSaleBillQueryAppService(),
	}
}

//
// OnSaleBillConfirmEvent
// @Description: SaleBillConfirmEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillConfirmEvent 领域事件
// @return error 错误
//
func (h *SaleBillQueryHandler) OnSaleBillConfirmEvent(ctx context.Context, event *event.SaleBillConfirmEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleBillView.NewBySaleBillConfirmEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Update(ctx, v)
	})
}

//
// OnSaleBillCreateEvent
// @Description: SaleBillCreateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillCreateEvent 领域事件
// @return error 错误
//
func (h *SaleBillQueryHandler) OnSaleBillCreateEvent(ctx context.Context, event *event.SaleBillCreateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleBillView.NewBySaleBillCreateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Create(ctx, v)
	})
}

//
// OnSaleBillDeleteEvent
// @Description: SaleBillDeleteEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillDeleteEvent 领域事件
// @return error 错误
//
func (h *SaleBillQueryHandler) OnSaleBillDeleteEvent(ctx context.Context, event *event.SaleBillDeleteEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		return h.service.DeleteById(ctx, event.GetTenantId(), event.Data.Id)
	})
}

//
// OnSaleBillUpdateEvent
// @Description: SaleBillUpdateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillUpdateEvent 领域事件
// @return error 错误
//
func (h *SaleBillQueryHandler) OnSaleBillUpdateEvent(ctx context.Context, event *event.SaleBillUpdateEvent) error {
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleBillView.NewBySaleBillUpdateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Update(ctx, v)
	})
}

func (h *SaleBillQueryHandler) GetStructName() string {
	return "dapr-ddd-demo.sale_bill.SaleBillQueryHandler"
}