package sale_bill_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/sale_bill/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// saleBillDeleteAllExecutor
// @Description: 新建
//
type saleBillCreateExecutor struct {
	saleBillService service.SaleBillQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleBillCreateExecutor) Execute(ctx context.Context, view *view.SaleBillView) error {
	if err := e.Validate(view); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.saleBillService.Create(ctx, view)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleBillCreateExecutor) Validate(view *view.SaleBillView) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newSaleBillCreateExecutor
// @Description: 新建命令执行器
// @return service.SaleBillDeleteExecutor
//
func newSaleBillCreateExecutor() *saleBillCreateExecutor {
	return &saleBillCreateExecutor{
		saleBillService: service_impl.GetSaleBillQueryDomainService(),
	}
}
