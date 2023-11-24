package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

// UserUpdateCommandExecutor
// @Description: 更新用户 命令执行器接口
type UserUpdateCommandExecutor interface {
	Execute(context.Context, *appcmd.UserUpdateAppCmd) error
}

// userUpdateCommandCommandExecutor
// @Description: 更新用户 命令执行器实现类
type userUpdateCommandExecutor struct {
	domainService *domain_service.UserCommandDomainService
}

// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
func (e *userUpdateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.UserUpdateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssUserUpdateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.UserUpdate(ctx, cmd)
	return err
}

// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
func (e *userUpdateCommandExecutor) Validate(appCmd *appcmd.UserUpdateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

// NewUserUpdateCommandExecutor
// @Description: 新建命令执行器
// @return service.UserUpdateCommandExecutor
func newUserUpdateCommandExecutor() *userUpdateCommandExecutor {
	return &userUpdateCommandExecutor{
		domainService: domain_service.NewUserCommandDomainService(),
	}
}
