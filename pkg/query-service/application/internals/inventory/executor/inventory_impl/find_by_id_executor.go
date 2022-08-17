package inventory_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/inventory/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// inventoryFindByIdExecutor
// @Description: 按Id查询
//
type inventoryFindByIdExecutor struct {
	domainService service.InventoryQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryFindByIdExecutor) Execute(ctx context.Context, aq *appquery.InventoryFindByIdAppQuery) (*view.InventoryView, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindById(ctx, query.NewInventoryFindByIdQuery(aq.TenantId, aq.Id))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryFindByIdExecutor) Validate(aq *appquery.InventoryFindByIdAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// NewInventoryCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.InventoryDeleteExecutor
//
func newInventoryFindByIdExecutor() *inventoryFindByIdExecutor {
	return &inventoryFindByIdExecutor{
		domainService: service_impl.GetInventoryQueryDomainService(),
	}
}
