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
// saleItemFindByIdExecutor
// @Description: 按Id查询
//
type saleItemFindByIdExecutor struct {
	domainService service.SaleItemQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemFindByIdExecutor) Execute(ctx context.Context, aq *appquery.SaleItemFindByIdAppQuery) (*view.SaleItemView, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindById(ctx, query.NewSaleItemFindByIdQuery(aq.TenantId, aq.Id))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemFindByIdExecutor) Validate(aq *appquery.SaleItemFindByIdAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// NewSaleItemCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleItemDeleteExecutor
//
func newSaleItemFindByIdExecutor() *saleItemFindByIdExecutor {
	return &saleItemFindByIdExecutor{
		domainService: service_impl.GetSaleItemQueryDomainService(),
	}
}
