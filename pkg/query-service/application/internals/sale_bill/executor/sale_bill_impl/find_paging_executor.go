package sale_bill_impl

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
// saleBillCreateCommandCommandExecutor
// @Description: 分页查询
//
type saleBillFindPagingExecutor struct {
	domainService service.SaleBillQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleBillFindPagingExecutor) Execute(ctx context.Context, aq *appquery.SaleBillFindPagingAppQuery) (*appquery.SaleBillFindPagingResult, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}
	qry := query.NewSaleBillFindPagingQuery(aq.TenantId, aq.Fields, aq.Filter, aq.Sort, aq.PageNum, aq.PageSize)
	fpr, ok, err := e.domainService.FindPaging(ctx, qry)
	if err != nil {
		return nil, false, err
	}
	res := assembler.SaleBill.AssFindPagingResult(fpr)
	return res, ok, nil
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleBillFindPagingExecutor) Validate(aq *appquery.SaleBillFindPagingAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// newSaleBillSaleBillFindPagingExecutor
// @Description: 新建命令执行器
// @return *saleBillFindPagingExecutor
//
func newSaleBillFindPagingExecutor() *saleBillFindPagingExecutor {
	return &saleBillFindPagingExecutor{
		domainService: service_impl.GetSaleBillQueryDomainService(),
	}
}
