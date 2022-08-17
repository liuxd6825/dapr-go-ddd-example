package sale_item_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/sale_bill/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// saleItemDeleteAllExecutor
// @Description: 按Id删除
//
type saleItemDeleteByIdExecutor struct {
	saleItemService service.SaleItemQueryDomainService
}

type saleItemDeleteByIdExecutorValidate struct {
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
func (e *saleItemDeleteByIdExecutor) Execute(ctx context.Context, tenantId string, id string) error {
	data := &saleItemDeleteByIdExecutorValidate{TenantId: tenantId, Id: id}
	if err := e.Validate(data); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.saleItemService.DeleteById(ctx, tenantId, id)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemDeleteByIdExecutor) Validate(v *saleItemDeleteByIdExecutorValidate) error {
	if v == nil {
		return errors.New("Validate() error: view is nil")
	}
	if len(v.TenantId) == 0 {
		return errors.New("Validate() error: tenantId is nil")
	}
	if len(v.Id) == 0 {
		return errors.New("Validate() error: id is nil")
	}
	return nil
}

//
// newSaleItemDeleteExecutor
// @Description: 新建命令执行器
// @return service.SaleItemDeleteExecutor
//
func newSaleItemDeleteByIdExecutor() *saleItemDeleteByIdExecutor {
	return &saleItemDeleteByIdExecutor{
		saleItemService: service_impl.GetSaleItemQueryDomainService(),
	}
}
