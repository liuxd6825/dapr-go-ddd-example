package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

// SaleBillConfirmCommandExecutor
// @Description: 下单确认命令 命令执行器接口
type SaleBillConfirmCommandExecutor interface {
	Execute(context.Context, *appcmd.SaleBillConfirmAppCmd) error
}

// saleBillConfirmCommandCommandExecutor
// @Description: 下单确认命令 命令执行器实现类
type saleBillConfirmCommandExecutor struct {
	domainService *domain_service.SaleBillCommandDomainService
}

// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
func (e *saleBillConfirmCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.SaleBillConfirmAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssSaleBillConfirmCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.SaleBillConfirm(ctx, cmd)
	return err
}

// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
func (e *saleBillConfirmCommandExecutor) Validate(appCmd *appcmd.SaleBillConfirmAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

// NewSaleBillConfirmCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleBillConfirmCommandExecutor
func newSaleBillConfirmCommandExecutor() *saleBillConfirmCommandExecutor {
	return &saleBillConfirmCommandExecutor{
		domainService: domain_service.NewSaleBillCommandDomainService(),
	}
}
