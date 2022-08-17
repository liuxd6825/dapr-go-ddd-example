package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// SaleItemUpdateCommandExecutor
// @Description: 更新扫描文件 命令执行器接口
//
type SaleItemUpdateCommandExecutor interface {
	Execute(context.Context, *appcmd.SaleItemUpdateAppCmd) error
}

//
// saleItemUpdateCommandCommandExecutor
// @Description: 更新扫描文件 命令执行器实现类
//
type saleItemUpdateCommandExecutor struct {
	domainService *domain_service.SaleBillCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemUpdateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.SaleItemUpdateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssSaleItemUpdateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.SaleItemUpdate(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemUpdateCommandExecutor) Validate(appCmd *appcmd.SaleItemUpdateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewSaleItemUpdateCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleItemUpdateCommandExecutor
//
func newSaleItemUpdateCommandExecutor() *saleItemUpdateCommandExecutor {
	return &saleItemUpdateCommandExecutor{
		domainService: domain_service.GetSaleBillCommandDomainService(),
	}
}
