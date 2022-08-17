package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// UserDeleteCommandExecutor
// @Description: 删除用户 命令执行器接口
//
type UserDeleteCommandExecutor interface {
	Execute(context.Context, *appcmd.UserDeleteAppCmd) error
}

//
// userDeleteCommandCommandExecutor
// @Description: 删除用户 命令执行器实现类
//
type userDeleteCommandExecutor struct {
	domainService *domain_service.UserCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *userDeleteCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.UserDeleteAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssUserDeleteCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.UserDelete(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userDeleteCommandExecutor) Validate(appCmd *appcmd.UserDeleteAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewUserDeleteCommandExecutor
// @Description: 新建命令执行器
// @return service.UserDeleteCommandExecutor
//
func newUserDeleteCommandExecutor() *userDeleteCommandExecutor {
	return &userDeleteCommandExecutor{
		domainService: domain_service.GetUserCommandDomainService(),
	}
}
