package model

import (
	"context"

	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

// UserCreateCommand
// @Description: 执行 UserCreateCommand 创建用户 命令
// @receiver a
// @param ctx 上下文
// @param cmd UserCreateCommand 命令
// @param metadata 元数据
// @return error 错误
func (a *UserAggregate) UserCreateCommand(ctx context.Context, cmd *command.UserCreateCommand, metadata map[string]string) (any, error) {
	e, err := factory.Event.NewUserCreateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	a.Id = cmd.Data.Id
	a.Name = cmd.Data.Name
	a.Email = cmd.Data.Email
	a.TenantId = cmd.Data.TenantId
	a.IsDeleted = false
	a.Remarks = cmd.Data.Remarks
	return ddd.CreateEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

// UserDeleteCommand
// @Description: 执行 UserDeleteCommand 删除用户 命令
// @receiver a
// @param ctx 上下文
// @param cmd UserDeleteCommand 命令
// @param metadata 元数据
// @return error 错误
func (a *UserAggregate) UserDeleteCommand(ctx context.Context, cmd *command.UserDeleteCommand, metadata map[string]string) (any, error) {
	e, err := factory.Event.NewUserDeleteEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.DeleteEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

// UserUpdateCommand
// @Description: 执行 UserUpdateCommand 更新用户 命令
// @receiver a
// @param ctx 上下文
// @param cmd UserUpdateCommand 命令
// @param metadata 元数据
// @return error 错误
func (a *UserAggregate) UserUpdateCommand(ctx context.Context, cmd *command.UserUpdateCommand, metadata map[string]string) (any, error) {
	e, err := factory.Event.NewUserUpdateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	a.Id = cmd.Data.Id
	a.Name = cmd.Data.Name
	a.Email = cmd.Data.Email
	a.TenantId = cmd.Data.TenantId
	a.IsDeleted = false
	a.Remarks = cmd.Data.Remarks
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}
