package model

import (
	"context"

	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

// InventoryCreateCommand
// @Description: 执行 InventoryCreateCommand 创建存货档案 命令
// @receiver a
// @param ctx 上下文
// @param cmd InventoryCreateCommand 命令
// @param metadata 元数据
// @return error 错误
func (a *InventoryAggregate) InventoryCreateCommand(ctx context.Context, cmd *command.InventoryCreateCommand, metadata map[string]string) (any, error) {
	e, err := factory.Event.NewInventoryCreateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.CreateEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

// InventoryUpdateCommand
// @Description: 执行 InventoryUpdateCommand 更新存货档案 命令
// @receiver a
// @param ctx 上下文
// @param cmd InventoryUpdateCommand 命令
// @param metadata 元数据
// @return error 错误
func (a *InventoryAggregate) InventoryUpdateCommand(ctx context.Context, cmd *command.InventoryUpdateCommand, metadata map[string]string) (any, error) {
	e, err := factory.Event.NewInventoryUpdateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}
