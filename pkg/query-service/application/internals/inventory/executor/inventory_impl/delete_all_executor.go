package inventory_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/inventory/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// inventoryDeleteAllExecutor
// @Description: 删除所有
//
type inventoryDeleteAllExecutor struct {
	inventoryService service.InventoryQueryDomainService
}

type inventoryDeleteAllExecutorValidate struct {
	TenantId string
	Id       string
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryDeleteAllExecutor) Execute(ctx context.Context, tenantId string) error {
	data := &inventoryDeleteAllExecutorValidate{TenantId: tenantId}
	if err := e.Validate(data); err != nil {
		return err
	}

	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.inventoryService.DeleteAll(ctx, tenantId)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryDeleteAllExecutor) Validate(v *inventoryDeleteAllExecutorValidate) error {
	if v == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newInventoryDeleteExecutor
// @Description: 新建命令执行器
// @return service.InventoryDeleteExecutor
//
func newInventoryDeleteAllExecutor() *inventoryDeleteAllExecutor {
	return &inventoryDeleteAllExecutor{
		inventoryService: service_impl.GetInventoryQueryDomainService(),
	}
}
