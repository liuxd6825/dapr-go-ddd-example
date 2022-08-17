package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// SaleItemCreateCommandExecutor
// @Description: 创建扫描文件 命令执行器接口
//
type SaleItemCreateCommandExecutor interface {
	Execute(context.Context, *appcmd.SaleItemCreateAppCmd) error
}

//
// saleItemCreateCommandCommandExecutor
// @Description: 创建扫描文件 命令执行器实现类
//
type saleItemCreateCommandExecutor struct {
	domainService *domain_service.SaleBillCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemCreateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.SaleItemCreateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssSaleItemCreateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.SaleItemCreate(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemCreateCommandExecutor) Validate(appCmd *appcmd.SaleItemCreateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewSaleItemCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleItemCreateCommandExecutor
//
func newSaleItemCreateCommandExecutor() *saleItemCreateCommandExecutor {
	return &saleItemCreateCommandExecutor{
		domainService: domain_service.GetSaleBillCommandDomainService(),
	}
}
