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
// inventoryUpdateExecutor
// @Description: 更新多个
//
type inventoryUpdateExecutor struct {
	inventoryService service.InventoryQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryUpdateExecutor) Execute(ctx context.Context, view *view.InventoryView) error {
	if err := e.Validate(view); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.inventoryService.Update(ctx, view)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryUpdateExecutor) Validate(view *view.InventoryView) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newInventoryCreateExecutor
// @Description: 新建命令执行器
// @return service.InventoryDeleteExecutor
//
func newInventoryUpdateExecutor() *inventoryUpdateExecutor {
	return &inventoryUpdateExecutor{
		inventoryService: service_impl.GetInventoryQueryDomainService(),
	}
}
