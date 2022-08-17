package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// NewUserCreateEvent
// @Description: 创建领域事件
//
func NewUserCreateEvent(ctx context.Context, cmd *command.UserCreateCommand, metadata *map[string]string) (*event.UserCreateEvent, error) {
	err := checkNewEventParas("NewUserCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewUserCreateEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewUserDeleteEvent
// @Description: 创建领域事件
//
func NewUserDeleteEvent(ctx context.Context, cmd *command.UserDeleteCommand, metadata *map[string]string) (*event.UserDeleteEvent, error) {
	err := checkNewEventParas("NewUserDeleteEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewUserDeleteEvent(cmd.CommandId)
	e.Data = cmd.Data
	return e, nil
}

//
// NewUserUpdateEvent
// @Description: 创建领域事件
//
func NewUserUpdateEvent(ctx context.Context, cmd *command.UserUpdateCommand, metadata *map[string]string) (*event.UserUpdateEvent, error) {
	err := checkNewEventParas("NewUserUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	e := event.NewUserUpdateEvent(cmd.CommandId)
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
