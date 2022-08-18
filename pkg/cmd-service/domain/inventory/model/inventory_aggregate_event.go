package model

import (
	"context"

	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
)

//
// OnInventoryCreateEvent
// @Description: InventoryCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *InventoryAggregate) OnInventoryCreateEvent(ctx context.Context, e *event.InventoryCreateEvent) error {
	return utils.Mapper(e.Data, a)
}

//
// OnInventoryUpdateEvent
// @Description: InventoryUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *InventoryAggregate) OnInventoryUpdateEvent(ctx context.Context, e *event.InventoryUpdateEvent) error {
	return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
}
