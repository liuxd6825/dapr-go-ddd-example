package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// SaleItemDeleteCommandExecutor
// @Description: 删除销售明细项 命令执行器接口
//
type SaleItemDeleteCommandExecutor interface {
	Execute(context.Context, *appcmd.SaleItemDeleteAppCmd) error
}

//
// saleItemDeleteCommandCommandExecutor
// @Description: 删除销售明细项 命令执行器实现类
//
type saleItemDeleteCommandExecutor struct {
	domainService *domain_service.SaleBillCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemDeleteCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.SaleItemDeleteAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssSaleItemDeleteCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.SaleItemDelete(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemDeleteCommandExecutor) Validate(appCmd *appcmd.SaleItemDeleteAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewSaleItemDeleteCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleItemDeleteCommandExecutor
//
func newSaleItemDeleteCommandExecutor() *saleItemDeleteCommandExecutor {
	return &saleItemDeleteCommandExecutor{
		domainService: domain_service.GetSaleBillCommandDomainService(),
	}
}
