package sale_bill_impl

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
// saleBillCreateCommandCommandExecutor
// @Description: 查找所有
//
type saleBillFindAllExecutor struct {
	domainService service.SaleBillQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleBillFindAllExecutor) Execute(ctx context.Context, aq *appquery.SaleBillFindAllAppQuery) ([]*view.SaleBillView, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindAll(ctx, query.NewSaleBillFindAllQuery(aq.TenantId))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleBillFindAllExecutor) Validate(aq *appquery.SaleBillFindAllAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// NewSaleBillCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.SaleBillDeleteExecutor
//
func newSaleBillFindAllExecutor() *saleBillFindAllExecutor {
	return &saleBillFindAllExecutor{
		domainService: service_impl.GetSaleBillQueryDomainService(),
	}
}
