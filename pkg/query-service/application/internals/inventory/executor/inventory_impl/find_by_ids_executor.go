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
// inventoryCreateCommandCommandExecutor
// @Description: 按Id查询多个
//
type inventoryFindByIdsExecutor struct {
	domainService service.InventoryQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryFindByIdsExecutor) Execute(ctx context.Context, aq *appquery.InventoryFindByIdsAppQuery) ([]*view.InventoryView, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindByIds(ctx, query.NewInventoryFindByIdsQuery(aq.TenantId, aq.Ids))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryFindByIdsExecutor) Validate(aq *appquery.InventoryFindByIdsAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	if len(aq.Ids) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.Ids is nil")
	}
	return nil
}

//
// NewInventoryCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.InventoryDeleteExecutor
//
func newInventoryFindByIdsExecutor() *inventoryFindByIdsExecutor {
	return &inventoryFindByIdsExecutor{
		domainService: service_impl.GetInventoryQueryDomainService(),
	}
}
