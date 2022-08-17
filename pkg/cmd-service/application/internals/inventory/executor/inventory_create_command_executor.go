package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// InventoryCreateCommandExecutor
// @Description: 创建存货档案 命令执行器接口
//
type InventoryCreateCommandExecutor interface {
	Execute(context.Context, *appcmd.InventoryCreateAppCmd) error
}

//
// inventoryCreateCommandCommandExecutor
// @Description: 创建存货档案 命令执行器实现类
//
type inventoryCreateCommandExecutor struct {
	domainService *domain_service.InventoryCommandDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryCreateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.InventoryCreateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssInventoryCreateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.InventoryCreate(ctx, cmd)
	return err
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryCreateCommandExecutor) Validate(appCmd *appcmd.InventoryCreateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

//
// NewInventoryCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.InventoryCreateCommandExecutor
//
func newInventoryCreateCommandExecutor() *inventoryCreateCommandExecutor {
	return &inventoryCreateCommandExecutor{
		domainService: domain_service.GetInventoryCommandDomainService(),
	}
}
