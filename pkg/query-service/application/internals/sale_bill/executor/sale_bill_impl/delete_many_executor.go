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
// SaleBillDeleteManyExecutor
// @Description: 删除多个
//
type saleBillDeleteManyExecutor struct {
	saleItemDomainService service.SaleItemQueryDomainService
	saleBillService       service.SaleBillQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleBillDeleteManyExecutor) Execute(ctx context.Context, tenantId string, vList []*view.SaleBillView) error {
	if err := e.Validate(vList); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		for _, item := range vList {
			if err := e.saleItemDomainService.DeleteBySaleBillId(ctx, tenantId, item.Id); err != nil {
				return err
			}
		}
		return e.saleBillService.DeleteMany(ctx, tenantId, vList)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleBillDeleteManyExecutor) Validate(view []*view.SaleBillView) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newSaleBillDeleteManyExecutor
// @Description: 新建命令执行器
// @return service.SaleBillDeleteExecutor
//
func newSaleBillDeleteManyExecutor() *saleBillDeleteManyExecutor {
	return &saleBillDeleteManyExecutor{
		saleBillService: service_impl.GetSaleBillQueryDomainService(),
	}
}
