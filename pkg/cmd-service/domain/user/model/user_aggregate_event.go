package model

import (
	"context"

	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
)

//
// OnUserCreateEvent
// @Description: UserCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *UserAggregate) OnUserCreateEvent(ctx context.Context, e *event.UserCreateEvent) error {
	return utils.Mapper(e.Data, a)
}

//
// OnUserDeleteEvent
// @Description: UserDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *UserAggregate) OnUserDeleteEvent(ctx context.Context, e *event.UserDeleteEvent) error {
	a.IsDeleted = true
	return nil
}

//
// OnUserUpdateEvent
// @Description: UserUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *UserAggregate) OnUserUpdateEvent(ctx context.Context, e *event.UserUpdateEvent) error {
	return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
}
