package sale_bill_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/sale_bill/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// saleBillDeleteAllExecutor
// @Description: 删除所有
//
type saleBillDeleteAllExecutor struct {
	saleItemDomainService service.SaleItemQueryDomainService
	saleBillService       service.SaleBillQueryDomainService
}

type saleBillDeleteAllExecutorValidate struct {
	TenantId string
	Id       string
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleBillDeleteAllExecutor) Execute(ctx context.Context, tenantId string) error {
	data := &saleBillDeleteAllExecutorValidate{TenantId: tenantId}
	if err := e.Validate(data); err != nil {
		return err
	}

	return session.StartSession(ctx, func(ctx context.Context) error {
		if err := e.saleItemDomainService.DeleteAll(ctx, tenantId); err != nil {
			return err
		}
		return e.saleBillService.DeleteAll(ctx, tenantId)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleBillDeleteAllExecutor) Validate(v *saleBillDeleteAllExecutorValidate) error {
	if v == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newSaleBillDeleteExecutor
// @Description: 新建命令执行器
// @return service.SaleBillDeleteExecutor
//
func newSaleBillDeleteAllExecutor() *saleBillDeleteAllExecutor {
	return &saleBillDeleteAllExecutor{
		saleBillService: service_impl.GetSaleBillQueryDomainService(),
	}
}
