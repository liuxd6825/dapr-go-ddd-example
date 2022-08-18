package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// eventFactory
// @Description: 存货档案事件工厂
//
type eventFactory struct {
}

var Event = eventFactory{}

//
// NewInventoryCreateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewInventoryCreateEvent(ctx context.Context, cmd *command.InventoryCreateCommand, metadata *map[string]string) (*event.InventoryCreateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewInventoryCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewInventoryCreateEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewInventoryUpdateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewInventoryUpdateEvent(ctx context.Context, cmd *command.InventoryUpdateCommand, metadata *map[string]string) (*event.InventoryUpdateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewInventoryUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewInventoryUpdateEvent(cmd.CommandId)
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
