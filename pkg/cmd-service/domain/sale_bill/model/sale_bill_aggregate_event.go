package model

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
)

//
// OnSaleBillConfirmEvent
// @Description: SaleBillConfirmEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillConfirmEvent(ctx context.Context, e *event.SaleBillConfirmEvent) error {
	panic("SaleBillAggregate.OnSaleBillConfirmEvent to=SaleBill error")
	return nil
}

//
// OnSaleBillCreateEvent
// @Description: SaleBillCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillCreateEvent(ctx context.Context, e *event.SaleBillCreateEvent) error {
	return utils.Mapper(e.Data, a)
}

//
// OnSaleBillDeleteEvent
// @Description: SaleBillDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillDeleteEvent(ctx context.Context, e *event.SaleBillDeleteEvent) error {
	a.IsDeleted = true
	return nil
}

//
// OnSaleBillUpdateEvent
// @Description: SaleBillUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillUpdateEvent(ctx context.Context, e *event.SaleBillUpdateEvent) error {
	return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
}

//
// OnSaleItemCreateEvent
// @Description: SaleItemCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemCreateEvent(ctx context.Context, e *event.SaleItemCreateEvent) error {
	for _, item := range e.Data.Items {
		if _, err := a.SaleItems.AddMapper(ctx, item.Id, item); err != nil {
			return err
		}
	}
	return nil
}

//
// OnSaleItemDeleteEvent
// @Description: SaleItemDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemDeleteEvent(ctx context.Context, e *event.SaleItemDeleteEvent) error {
	for _, item := range e.Data.Items {
		if err := a.SaleItems.DeleteById(ctx, item.Id); err != nil {
			return err
		}
	}
	return nil
}

//
// OnSaleItemUpdateEvent
// @Description: SaleItemUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemUpdateEvent(ctx context.Context, e *event.SaleItemUpdateEvent) error {
	for _, item := range e.Data.Items {
		if _, err := a.SaleItems.AddMapper(ctx, item.Id, item); err != nil {
			return err
		}
	}
	return nil
}
