package inventory_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/inventory/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// InventoryDeleteManyExecutor
// @Description: 删除多个
//
type inventoryDeleteManyExecutor struct {
	inventoryService service.InventoryQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryDeleteManyExecutor) Execute(ctx context.Context, tenantId string, vList []*view.InventoryView) error {
	if err := e.Validate(vList); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.inventoryService.DeleteMany(ctx, tenantId, vList)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryDeleteManyExecutor) Validate(view []*view.InventoryView) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newInventoryDeleteManyExecutor
// @Description: 新建命令执行器
// @return service.InventoryDeleteExecutor
//
func newInventoryDeleteManyExecutor() *inventoryDeleteManyExecutor {
	return &inventoryDeleteManyExecutor{
		inventoryService: service_impl.GetInventoryQueryDomainService(),
	}
}
