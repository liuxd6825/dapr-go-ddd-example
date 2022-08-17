package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// SaleBillUpdateCommandExecutor
// @Description: 更新销售订单 命令执行器接口
//
type SaleBillUpdateCommandExecutor interface {
	Execute(context.Context, *appcmd.SaleBillUpdateAppCmd) error
}

//
// saleBillUpdateCommandCommandExecutor
// @Description: 更新销售订单 命令执行器实现类
//
type saleBillUpdateCommandExecutor struct {
	domainService *domain_service.SaleBillCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleBillUpdateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.SaleBillUpdateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssSaleBillUpdateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.SaleBillUpdate(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleBillUpdateCommandExecutor) Validate(appCmd *appcmd.SaleBillUpdateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewSaleBillUpdateCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleBillUpdateCommandExecutor
//
func newSaleBillUpdateCommandExecutor() *saleBillUpdateCommandExecutor {
	return &saleBillUpdateCommandExecutor{
		domainService: domain_service.GetSaleBillCommandDomainService(),
	}
}
