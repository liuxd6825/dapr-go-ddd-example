package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

// SaleBillCreateCommandExecutor
// @Description: 创建销售订单 命令执行器接口
type SaleBillCreateCommandExecutor interface {
	Execute(context.Context, *appcmd.SaleBillCreateAppCmd) error
}

// saleBillCreateCommandCommandExecutor
// @Description: 创建销售订单 命令执行器实现类
type saleBillCreateCommandExecutor struct {
	domainService *domain_service.SaleBillCommandDomainService
}

// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
func (e *saleBillCreateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.SaleBillCreateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssSaleBillCreateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.SaleBillCreate(ctx, cmd)
	return err
}

// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
func (e *saleBillCreateCommandExecutor) Validate(appCmd *appcmd.SaleBillCreateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

// NewSaleBillCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleBillCreateCommandExecutor
func newSaleBillCreateCommandExecutor() *saleBillCreateCommandExecutor {
	return &saleBillCreateCommandExecutor{
		domainService: domain_service.NewSaleBillCommandDomainService(),
	}
}
