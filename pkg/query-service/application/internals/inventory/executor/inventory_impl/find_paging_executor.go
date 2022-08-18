package inventory_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/inventory/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// inventoryCreateCommandCommandExecutor
// @Description: 分页查询
//
type inventoryFindPagingExecutor struct {
	domainService service.InventoryQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *inventoryFindPagingExecutor) Execute(ctx context.Context, aq *appquery.InventoryFindPagingAppQuery) (*appquery.InventoryFindPagingResult, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}
	qry := query.NewInventoryFindPagingQuery(aq.TenantId, aq.Fields, aq.Filter, aq.Sort, aq.PageNum, aq.PageSize)
	fpr, ok, err := e.domainService.FindPaging(ctx, qry)
	if err != nil {
		return nil, false, err
	}
	res := assembler.Inventory.AssFindPagingResult(fpr)
	return res, ok, nil
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *inventoryFindPagingExecutor) Validate(aq *appquery.InventoryFindPagingAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// newInventoryInventoryFindPagingExecutor
// @Description: 新建命令执行器
// @return *inventoryFindPagingExecutor
//
func newInventoryFindPagingExecutor() *inventoryFindPagingExecutor {
	return &inventoryFindPagingExecutor{
		domainService: service_impl.GetInventoryQueryDomainService(),
	}
}
