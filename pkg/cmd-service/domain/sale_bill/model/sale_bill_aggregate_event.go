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
func (a *SaleBillAggregate) OnSaleBillConfirmEventV1s0(ctx context.Context, e *event.SaleBillConfirmEvent) error {
	panic("SaleBillAggregate.OnSaleBillConfirmEvent to=SaleBill error")
	return nil
}

//
// OnSaleBillCreateEventV1s0
// @Description: SaleBillCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillCreateEventV1s0(ctx context.Context, e *event.SaleBillCreateEvent) error {
	return utils.Mapper(e.Data, a)
}

//
// OnSaleBillDeleteEventV1s0
// @Description: SaleBillDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillDeleteEventV1s0(ctx context.Context, e *event.SaleBillDeleteEvent) error {
	a.IsDeleted = true
	return nil
}

//
// OnSaleBillUpdateEventV1s0
// @Description: SaleBillUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillUpdateEventV1s0(ctx context.Context, e *event.SaleBillUpdateEvent) error {
	return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
}

//
// OnSaleItemCreateEventV1s0
// @Description: SaleItemCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemCreateEventV1s0(ctx context.Context, e *event.SaleItemCreateEvent) error {
	for _, item := range e.Data.Items {
		a.TotalMoney = a.TotalMoney + item.Money
		if _, err := a.SaleItems.AddMapper(ctx, item.Id, item); err != nil {
			return err
		}
	}
	return nil
}

//
// OnSaleItemDeleteEventV1s0
// @Description: SaleItemDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemDeleteEventV1s0(ctx context.Context, e *event.SaleItemDeleteEvent) error {
	for _, item := range e.Data.Items {
		saleItem, ok := a.SaleItems.Items.Get(item.GetId())
		if ok {
			a.TotalMoney = a.TotalMoney - saleItem.Money
			if err := a.SaleItems.DeleteById(ctx, item.Id); err != nil {
				return err
			}
		}
	}
	return nil
}

//
// OnSaleItemUpdateEventV1s0
// @Description: SaleItemUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemUpdateEventV1s0(ctx context.Context, e *event.SaleItemUpdateEvent) error {
	for _, item := range e.Data.Items {
		saleItem, ok := a.SaleItems.Items.Get(item.GetId())
		if ok {
			a.TotalMoney = a.TotalMoney - saleItem.Money + item.Money
			if _, err := a.SaleItems.AddMapper(ctx, item.Id, item); err != nil {
				return err
			}
		}
	}
	return nil
}
