package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// UserCreateCommandExecutor
// @Description: 创建用户 命令执行器接口
//
type UserCreateCommandExecutor interface {
	Execute(context.Context, *appcmd.UserCreateAppCmd) error
}

//
// userCreateCommandCommandExecutor
// @Description: 创建用户 命令执行器实现类
//
type userCreateCommandExecutor struct {
	domainService *domain_service.UserCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *userCreateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.UserCreateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssUserCreateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.UserCreate(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userCreateCommandExecutor) Validate(appCmd *appcmd.UserCreateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewUserCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.UserCreateCommandExecutor
//
func newUserCreateCommandExecutor() *userCreateCommandExecutor {
	return &userCreateCommandExecutor{
		domainService: domain_service.GetUserCommandDomainService(),
	}
}
