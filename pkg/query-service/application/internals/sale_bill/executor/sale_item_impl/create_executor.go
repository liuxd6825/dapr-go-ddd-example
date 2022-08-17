package sale_item_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/sale_bill/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// saleItemDeleteAllExecutor
// @Description: 新建
//
type saleItemCreateExecutor struct {
	saleItemService service.SaleItemQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemCreateExecutor) Execute(ctx context.Context, view *view.SaleItemView) error {
	if err := e.Validate(view); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.saleItemService.Create(ctx, view)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemCreateExecutor) Validate(view *view.SaleItemView) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newSaleItemCreateExecutor
// @Description: 新建命令执行器
// @return service.SaleItemDeleteExecutor
//
func newSaleItemCreateExecutor() *saleItemCreateExecutor {
	return &saleItemCreateExecutor{
		saleItemService: service_impl.GetSaleItemQueryDomainService(),
	}
}
