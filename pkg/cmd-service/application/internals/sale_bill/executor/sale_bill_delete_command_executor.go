package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// SaleBillDeleteCommandExecutor
// @Description: 删除销售订单 命令执行器接口
//
type SaleBillDeleteCommandExecutor interface {
	Execute(context.Context, *appcmd.SaleBillDeleteAppCmd) error
}

//
// saleBillDeleteCommandCommandExecutor
// @Description: 删除销售订单 命令执行器实现类
//
type saleBillDeleteCommandExecutor struct {
	domainService *domain_service.SaleBillCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleBillDeleteCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.SaleBillDeleteAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssSaleBillDeleteCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.SaleBillDelete(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleBillDeleteCommandExecutor) Validate(appCmd *appcmd.SaleBillDeleteAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewSaleBillDeleteCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleBillDeleteCommandExecutor
//
func newSaleBillDeleteCommandExecutor() *saleBillDeleteCommandExecutor {
	return &saleBillDeleteCommandExecutor{
		domainService: domain_service.GetSaleBillCommandDomainService(),
	}
}
