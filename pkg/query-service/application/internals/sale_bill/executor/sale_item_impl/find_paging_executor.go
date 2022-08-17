package sale_item_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/sale_bill/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// saleItemCreateCommandCommandExecutor
// @Description: 分页查询
//
type saleItemFindPagingExecutor struct {
	domainService service.SaleItemQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemFindPagingExecutor) Execute(ctx context.Context, aq *appquery.SaleItemFindPagingAppQuery) (*appquery.SaleItemFindPagingResult, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}
	qry := query.NewSaleItemFindPagingQuery(aq.TenantId, aq.Fields, aq.Filter, aq.Sort, aq.PageNum, aq.PageSize)
	fpr, ok, err := e.domainService.FindPaging(ctx, qry)
	if err != nil {
		return nil, false, err
	}
	res := assembler.AssSaleItemFindPagingResult(fpr)
	return res, ok, nil
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemFindPagingExecutor) Validate(aq *appquery.SaleItemFindPagingAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// newSaleItemSaleItemFindPagingExecutor
// @Description: 新建命令执行器
// @return *saleItemFindPagingExecutor
//
func newSaleItemFindPagingExecutor() *saleItemFindPagingExecutor {
	return &saleItemFindPagingExecutor{
		domainService: service_impl.GetSaleItemQueryDomainService(),
	}
}
