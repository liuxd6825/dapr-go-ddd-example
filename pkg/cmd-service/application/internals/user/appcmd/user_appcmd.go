package appcmd

import (
	domain "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
)

//
// UserUpdateAppCmd
// @Description: 应用服务层命令, 更新用户
//
type UserUpdateAppCmd struct {
	domain.UserUpdateCommand
}

//
// UserDeleteAppCmd
// @Description: 应用服务层命令, 删除用户
//
type UserDeleteAppCmd struct {
	domain.UserDeleteCommand
}

//
// UserCreateAppCmd
// @Description: 应用服务层命令, 创建用户
//
type UserCreateAppCmd struct {
	domain.UserCreateCommand
}