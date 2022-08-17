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
// @Description: 按SaleBillId删除
//
type saleItemDeleteBySaleBillIdExecutor struct {
	saleItemService service.SaleItemQueryDomainService
}

type saleItemDeleteBySaleBillIdExecutorValidate struct {
	TenantId   string
	SaleBillId string
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *saleItemDeleteBySaleBillIdExecutor) Execute(ctx context.Context, tenantId string, saleBillId string) error {
	data := &saleItemDeleteBySaleBillIdExecutorValidate{TenantId: tenantId, SaleBillId: saleBillId}
	if err := e.Validate(data); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.saleItemService.DeleteBySaleBillId(ctx, data.TenantId, data.SaleBillId)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *saleItemDeleteBySaleBillIdExecutor) Validate(v *saleItemDeleteBySaleBillIdExecutorValidate) error {
	if v == nil {
		return errors.New("Validate(v) error: v is nil")
	}
	if len(v.TenantId) == 0 {
		return errors.New("Validate(v) error: v.TenantId is nil")
	}
	if len(v.SaleBillId) == 0 {
		return errors.New("Validate(v) error: v.SaleBillId is nil")
	}
	return nil
}

//
// newSaleItemDeleteBySaleBillIdExecutor
// @Description: 新建命令执行器
// @return service.SaleItemDeleteBySaleBillIdExecutor
//
func newSaleItemDeleteBySaleBillIdExecutor() *saleItemDeleteBySaleBillIdExecutor {
	return &saleItemDeleteBySaleBillIdExecutor{
		saleItemService: service_impl.GetSaleItemQueryDomainService(),
	}
}
