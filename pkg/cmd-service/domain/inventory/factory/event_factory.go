package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// NewInventoryCreateEvent
// @Description: 创建领域事件
//
func NewInventoryCreateEvent(ctx context.Context, cmd *command.InventoryCreateCommand, metadata *map[string]string) (*event.InventoryCreateEvent, error) {
	err := checkNewEventParas("NewInventoryCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewInventoryCreateEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewInventoryUpdateEvent
// @Description: 创建领域事件
//
func NewInventoryUpdateEvent(ctx context.Context, cmd *command.InventoryUpdateCommand, metadata *map[string]string) (*event.InventoryUpdateEvent, error) {
	err := checkNewEventParas("NewInventoryUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewInventoryUpdateEvent(cmd.CommandId)
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
