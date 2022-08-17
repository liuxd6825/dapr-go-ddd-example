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
// saleItemFindBySaleBillIdExecutor
// @Description: 按SaleBillId查询
//
type saleItemFindBySaleBillIdExecutor struct {
	domainService service.SaleItemQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemFindBySaleBillIdExecutor) Execute(ctx context.Context, aq *appquery.SaleItemFindBySaleBillIdAppQuery) ([]*view.SaleItemView, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindBySaleBillId(ctx, query.NewSaleItemFindBySaleBillIdQuery(aq.TenantId, aq.SaleBillId))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemFindBySaleBillIdExecutor) Validate(aq *appquery.SaleItemFindBySaleBillIdAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	if len(aq.SaleBillId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.SaleBillId is nil")
	}
	return nil
}

//
// newSaleItemFindBySaleBillIdExecutor()
// @Description: 新建命令执行器
// @return *saleItemFindBySaleBillIdExecutor
//
func newSaleItemFindBySaleBillIdExecutor() *saleItemFindBySaleBillIdExecutor {
	return &saleItemFindBySaleBillIdExecutor{
		domainService: service_impl.GetSaleItemQueryDomainService(),
	}
}
