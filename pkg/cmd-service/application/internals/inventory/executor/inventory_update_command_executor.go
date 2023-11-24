package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/assembler"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

// InventoryUpdateCommandExecutor
// @Description: 更新存货档案 命令执行器接口
type InventoryUpdateCommandExecutor interface {
	Execute(context.Context, *appcmd.InventoryUpdateAppCmd) error
}

// inventoryUpdateCommandCommandExecutor
// @Description: 更新存货档案 命令执行器实现类
type inventoryUpdateCommandExecutor struct {
	domainService *domain_service.InventoryCommandDomainService
}

// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
func (e *inventoryUpdateCommandExecutor) Execute(ctx context.Context, appCmd *appcmd.InventoryUpdateAppCmd) error {
	if err := e.Validate(appCmd); err != nil {
		return err
	}

	cmd, err := assembler.AssInventoryUpdateCommand(ctx, appCmd)
	if err != nil {
		return err
	}

	_, err = e.domainService.InventoryUpdate(ctx, cmd)
	return err
}

// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
func (e *inventoryUpdateCommandExecutor) Validate(appCmd *appcmd.InventoryUpdateAppCmd) error {
	if appCmd == nil {
		return errors.New("appCmd is nil")
	}
	return nil
}

// NewInventoryUpdateCommandExecutor
// @Description: 新建命令执行器
// @return service.InventoryUpdateCommandExecutor
func newInventoryUpdateCommandExecutor() *inventoryUpdateCommandExecutor {
	return &inventoryUpdateCommandExecutor{
		domainService: domain_service.NewInventoryCommandDomainService(),
	}
}
