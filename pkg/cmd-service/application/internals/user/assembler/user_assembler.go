package assembler

import (
	"context"
	"errors"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
)

// AssUserDeleteCommand
// @Description:  UserDeleteAppCmd 删除用户转换器
// @param ctx
// @param cmdDto
// @return *command.UserDeleteCommand
// @return error
func AssUserDeleteCommand(ctx context.Context, appCmd *appcmd.UserDeleteAppCmd) (*command.UserDeleteCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssUserDeleteCommand() appCmd is nil")
	}
	return appCmd, nil
}

// AssUserCreateCommand
// @Description:  UserCreateAppCmd 创建用户转换器
// @param ctx
// @param cmdDto
// @return *command.UserCreateCommand
// @return error
func AssUserCreateCommand(ctx context.Context, appCmd *appcmd.UserCreateAppCmd) (*command.UserCreateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssUserCreateCommand() appCmd is nil")
	}
	return appCmd, nil
}

// AssUserUpdateCommand
// @Description:  UserUpdateAppCmd 更新用户转换器
// @param ctx
// @param cmdDto
// @return *command.UserUpdateCommand
// @return error
func AssUserUpdateCommand(ctx context.Context, appCmd *appcmd.UserUpdateAppCmd) (*command.UserUpdateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssUserUpdateCommand() appCmd is nil")
	}
	return appCmd, nil
}
