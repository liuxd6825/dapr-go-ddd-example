package model

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// SaleBillConfirmCommand
// @Description: 执行 SaleBillConfirmCommand 下单确认命令 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillConfirmCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillConfirmCommand(ctx context.Context, cmd *command.SaleBillConfirmCommand, metadata *map[string]string) (any, error) {
	e, err := factory.Event.NewSaleBillConfirmEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleBillCreateCommand
// @Description: 执行 SaleBillCreateCommand 创建销售订单 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillCreateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillCreateCommand(ctx context.Context, cmd *command.SaleBillCreateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.Event.NewSaleBillCreateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.CreateEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleBillDeleteCommand
// @Description: 执行 SaleBillDeleteCommand 删除销售订单 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillDeleteCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillDeleteCommand(ctx context.Context, cmd *command.SaleBillDeleteCommand, metadata *map[string]string) (any, error) {
	e, err := factory.Event.NewSaleBillDeleteEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.DeleteEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleBillUpdateCommand
// @Description: 执行 SaleBillUpdateCommand 更新销售订单 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillUpdateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillUpdateCommand(ctx context.Context, cmd *command.SaleBillUpdateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.Event.NewSaleBillUpdateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleItemCreateCommand
// @Description: 执行 SaleItemCreateCommand 添加明细 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleItemCreateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleItemCreateCommand(ctx context.Context, cmd *command.SaleItemCreateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.Event.NewSaleItemCreateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleItemDeleteCommand
// @Description: 执行 SaleItemDeleteCommand 删除销售明细项 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleItemDeleteCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleItemDeleteCommand(ctx context.Context, cmd *command.SaleItemDeleteCommand, metadata *map[string]string) (any, error) {
	e, err := factory.Event.NewSaleItemDeleteEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleItemUpdateCommand
// @Description: 执行 SaleItemUpdateCommand 更新明细 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleItemUpdateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleItemUpdateCommand(ctx context.Context, cmd *command.SaleItemUpdateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.Event.NewSaleItemUpdateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}
