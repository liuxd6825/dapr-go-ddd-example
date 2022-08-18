package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// eventFactory
// @Description: 销售订单事件工厂
//
type eventFactory struct {
}

var Event = eventFactory{}

//
// NewSaleBillConfirmEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewSaleBillConfirmEvent(ctx context.Context, cmd *command.SaleBillConfirmCommand, metadata *map[string]string) (*event.SaleBillConfirmEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewSaleBillConfirmEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewSaleBillConfirmEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewSaleBillCreateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewSaleBillCreateEvent(ctx context.Context, cmd *command.SaleBillCreateCommand, metadata *map[string]string) (*event.SaleBillCreateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewSaleBillCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewSaleBillCreateEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewSaleBillDeleteEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewSaleBillDeleteEvent(ctx context.Context, cmd *command.SaleBillDeleteCommand, metadata *map[string]string) (*event.SaleBillDeleteEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewSaleBillDeleteEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewSaleBillDeleteEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewSaleBillUpdateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewSaleBillUpdateEvent(ctx context.Context, cmd *command.SaleBillUpdateCommand, metadata *map[string]string) (*event.SaleBillUpdateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewSaleBillUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewSaleBillUpdateEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewSaleItemCreateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewSaleItemCreateEvent(ctx context.Context, cmd *command.SaleItemCreateCommand, metadata *map[string]string) (*event.SaleItemCreateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewSaleItemCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewSaleItemCreateEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewSaleItemDeleteEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewSaleItemDeleteEvent(ctx context.Context, cmd *command.SaleItemDeleteCommand, metadata *map[string]string) (*event.SaleItemDeleteEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewSaleItemDeleteEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewSaleItemDeleteEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewSaleItemUpdateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewSaleItemUpdateEvent(ctx context.Context, cmd *command.SaleItemUpdateCommand, metadata *map[string]string) (*event.SaleItemUpdateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewSaleItemUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewSaleItemUpdateEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// checkNewEventParas
// @Description: 检查参数是否正确
//
func (e eventFactory) checkNewEventParas(funcName string, ctx context.Context, cmd interface{}, metadata *map[string]string) error {
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
