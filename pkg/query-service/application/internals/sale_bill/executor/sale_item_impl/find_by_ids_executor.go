package sale_item_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/sale_bill/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// saleItemCreateCommandCommandExecutor
// @Description: 按Id查询多个
//
type saleItemFindByIdsExecutor struct {
	domainService service.SaleItemQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemFindByIdsExecutor) Execute(ctx context.Context, aq *appquery.SaleItemFindByIdsAppQuery) ([]*view.SaleItemView, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindByIds(ctx, query.NewSaleItemFindByIdsQuery(aq.TenantId, aq.Ids))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemFindByIdsExecutor) Validate(aq *appquery.SaleItemFindByIdsAppQuery) error {
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
// NewSaleItemCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleItemDeleteExecutor
//
func newSaleItemFindByIdsExecutor() *saleItemFindByIdsExecutor {
	return &saleItemFindByIdsExecutor{
		domainService: service_impl.GetSaleItemQueryDomainService(),
	}
}
