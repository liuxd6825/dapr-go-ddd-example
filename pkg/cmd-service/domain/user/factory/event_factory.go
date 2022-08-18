package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// eventFactory
// @Description: 用户事件工厂
//
type eventFactory struct {
}

var Event = eventFactory{}

//
// NewUserCreateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewUserCreateEvent(ctx context.Context, cmd *command.UserCreateCommand, metadata *map[string]string) (*event.UserCreateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewUserCreateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewUserCreateEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewUserDeleteEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewUserDeleteEvent(ctx context.Context, cmd *command.UserDeleteCommand, metadata *map[string]string) (*event.UserDeleteEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewUserDeleteEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewUserDeleteEvent(cmd.CommandId)
	v.Data = cmd.Data
	return v, nil
}

//
// NewUserUpdateEvent
// @Description: 创建领域事件
//
func (e eventFactory) NewUserUpdateEvent(ctx context.Context, cmd *command.UserUpdateCommand, metadata *map[string]string) (*event.UserUpdateEvent, error) {
	err := e.checkNewEventParas("eventFactory.NewUserUpdateEvent", ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	v := event.NewUserUpdateEvent(cmd.CommandId)
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
