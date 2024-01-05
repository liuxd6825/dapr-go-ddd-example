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

type SaleBillQueryHandler struct {
	service *service.SaleBillQueryAppService
	handler.BaseQueryHandler
}

// NewSaleBillSubscribe
// @Description: 创建dapr消息订阅器，用于接受领域事件
// @return restapp.RegisterSubscribe  消息注册器
func NewSaleBillSubscribe() restapp.RegisterSubscribe {
	subscribes := []*ddd.Subscribe{
		{PubsubName: "pubsub", Topic: event.SaleBillConfirmEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_confirm_event"},
		{PubsubName: "pubsub", Topic: event.SaleBillCreateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_create_event"},
		{PubsubName: "pubsub", Topic: event.SaleBillDeleteEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_delete_event"},
		{PubsubName: "pubsub", Topic: event.SaleBillUpdateEventType.String(), Route: "/dapr-ddd-demo/domain-event/sale_bill/sale_bill_update_event"},
	}
	return restapp.NewRegisterSubscribe(subscribes, NewSaleBillQueryHandler())
}

// NewSaleBillQueryHandler
// @Description: 创建<no value>领域事件处理器
// @return ddd.QueryEventHandler 领域事件处理器
func NewSaleBillQueryHandler() ddd.QueryEventHandler {
	return &SaleBillQueryHandler{
		service: service.GetSaleBillQueryAppService(),
	}
}

// OnSaleBillConfirmEventV1s0
// @Description: SaleBillConfirmEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillConfirmEvent 领域事件
// @return error 错误
func (h *SaleBillQueryHandler) OnSaleBillConfirmEventV1s0(ctx context.Context, event *event.SaleBillConfirmEvent) error {
	logs.DebugEvent(event, "OnOnSaleBillConfirmEventV1s0")
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleBillView.NewBySaleBillConfirmEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Update(ctx, v)
	})
}

// OnSaleBillCreateEventV1s0
// @Description: SaleBillCreateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillCreateEvent 领域事件
// @return error 错误
func (h *SaleBillQueryHandler) OnSaleBillCreateEventV1s0(ctx context.Context, event *event.SaleBillCreateEvent) error {
	logs.DebugEvent(event, "OnOnSaleBillCreateEventV1s0")
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		v, err := factory.SaleBillView.NewBySaleBillCreateEvent(ctx, event)
		if err != nil {
			return err
		}
		return h.service.Create(ctx, v)
	})
}

// OnSaleBillDeleteEventV1s0
// @Description: SaleBillDeleteEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillDeleteEvent 领域事件
// @return error 错误
func (h *SaleBillQueryHandler) OnSaleBillDeleteEventV1s0(ctx context.Context, event *event.SaleBillDeleteEvent) error {
	logs.DebugEvent(event, "OnOnSaleBillDeleteEventV1s0")
	return h.DoSession(ctx, h, event, func(ctx context.Context) error {
		return h.service.DeleteById(ctx, event.GetTenantId(), event.Data.Id)
	})
}

// OnSaleBillUpdateEventV1s0
// @Description: SaleBillUpdateEvent事件处理器
// @receiver h
// @param ctx 上下文
// @param event SaleBillUpdateEvent 领域事件
// @return error 错误
func (h *SaleBillQueryHandler) OnSaleBillUpdateEventV1s0(ctx context.Context, event *event.SaleBillUpdateEvent) error {
	logs.DebugEvent(event, "OnOnSaleBillUpdateEventV1s0")
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
