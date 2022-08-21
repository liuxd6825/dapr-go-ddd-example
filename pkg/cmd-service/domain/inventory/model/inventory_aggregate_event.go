package model

import (
	"context"

	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
)

//
// OnInventoryCreateEventV1s0
// @Description: InventoryCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *InventoryAggregate) OnInventoryCreateEventV1s0(ctx context.Context, e *event.InventoryCreateEvent) error {
	return utils.Mapper(e.Data, a)
}

//
// OnInventoryUpdateEventV1s0
// @Description: InventoryUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *InventoryAggregate) OnInventoryUpdateEventV1s0(ctx context.Context, e *event.InventoryUpdateEvent) error {
	return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
}
