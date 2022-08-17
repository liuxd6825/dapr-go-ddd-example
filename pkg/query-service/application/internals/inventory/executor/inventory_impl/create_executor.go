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
// inventoryDeleteAllExecutor
// @Description: 新建
//
type inventoryCreateExecutor struct {
	inventoryService service.InventoryQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryCreateExecutor) Execute(ctx context.Context, view *view.InventoryView) error {
	if err := e.Validate(view); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.inventoryService.Create(ctx, view)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryCreateExecutor) Validate(view *view.InventoryView) error {
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
func newInventoryCreateExecutor() *inventoryCreateExecutor {
	return &inventoryCreateExecutor{
		inventoryService: service_impl.GetInventoryQueryDomainService(),
	}
}
