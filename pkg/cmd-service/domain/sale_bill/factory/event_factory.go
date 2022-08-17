package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// NewSaleBillConfirmEvent
// @Description: 创建领域事件
//
func NewSaleBillConfirmEvent(ctx context.Context, cmd *command.SaleBillConfirmCommand, metadata *map[string]string) (*event.SaleBillConfirmEvent, error) {
	err := checkNewEventParas("NewSaleBillConfirmEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewSaleBillConfirmEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewSaleBillCreateEvent
// @Description: 创建领域事件
//
func NewSaleBillCreateEvent(ctx context.Context, cmd *command.SaleBillCreateCommand, metadata *map[string]string) (*event.SaleBillCreateEvent, error) {
	err := checkNewEventParas("NewSaleBillCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewSaleBillCreateEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewSaleBillDeleteEvent
// @Description: 创建领域事件
//
func NewSaleBillDeleteEvent(ctx context.Context, cmd *command.SaleBillDeleteCommand, metadata *map[string]string) (*event.SaleBillDeleteEvent, error) {
	err := checkNewEventParas("NewSaleBillDeleteEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewSaleBillDeleteEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewSaleBillUpdateEvent
// @Description: 创建领域事件
//
func NewSaleBillUpdateEvent(ctx context.Context, cmd *command.SaleBillUpdateCommand, metadata *map[string]string) (*event.SaleBillUpdateEvent, error) {
	err := checkNewEventParas("NewSaleBillUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewSaleBillUpdateEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewSaleItemCreateEvent
// @Description: 创建领域事件
//
func NewSaleItemCreateEvent(ctx context.Context, cmd *command.SaleItemCreateCommand, metadata *map[string]string) (*event.SaleItemCreateEvent, error) {
	err := checkNewEventParas("NewSaleItemCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewSaleItemCreateEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewSaleItemDeleteEvent
// @Description: 创建领域事件
//
func NewSaleItemDeleteEvent(ctx context.Context, cmd *command.SaleItemDeleteCommand, metadata *map[string]string) (*event.SaleItemDeleteEvent, error) {
	err := checkNewEventParas("NewSaleItemDeleteEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewSaleItemDeleteEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewSaleItemUpdateEvent
// @Description: 创建领域事件
//
func NewSaleItemUpdateEvent(ctx context.Context, cmd *command.SaleItemUpdateCommand, metadata *map[string]string) (*event.SaleItemUpdateEvent, error) {
	err := checkNewEventParas("NewSaleItemUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewSaleItemUpdateEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// checkNewEventParas
// @Description: 检查参数是否正确
//
func checkNewEventParas(funcName string, ctx context.Context, cmd interface{}, metadata *map[string]string) error {
	if ctx == nil {
		return errors.ErrorOf("%s(ctx, cmd, metadata) error: ctx is nil", funcName)
	}
	if cmd == nil {
		return errors.ErrorOf("%s(ctx, cmd, metadata) error: cmd is nil", funcName)
	}
	if metadata == nil {
		return errors.ErrorOf("%s(ctx, cmd, metadata) error: metadata is nil", funcName)
	}
	return nil
}
