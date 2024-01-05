package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/service"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type SaleItemCommandApi struct {
	service *service.SaleBillCommandAppService
}

func NewSaleItemCommandApi() *SaleItemCommandApi {
	return &SaleItemCommandApi{
		service: service.NewSaleBillCommandAppService(),
	}
}

// BeforeActivation
// @Description: 注册http
// @receiver c
// @param ctx
func (c *SaleItemCommandApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("DELETE", "/tenants/{tenantId}/sale-bills/sale-items", "SaleItemDelete")
	b.Handle("POST", "/tenants/{tenantId}/sale-bills/sale-items", "SaleItemCreate")
	b.Handle("POST", "/tenants/{tenantId}/sale-bills/sale-items:get", "SaleItemCreateAndGet")
	b.Handle("PATCH", "/tenants/{tenantId}/sale-bills/sale-items", "SaleItemUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/sale-bills/sale-items:get", "SaleItemUpdateAndGet")
}

// SaleItemDelete
// @Description: 删除销售明细项
// @receiver c
// @param ctx
func (c *SaleItemCommandApi) SaleItemDelete(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		cmd, err := saleBillAssembler.AssSaleItemDeleteAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.SaleItemDelete(ctx, cmd)
	})
}

// SaleItemCreate
// @Description: 添加明细
// @receiver c
// @param ctx
func (c *SaleItemCommandApi) SaleItemCreate(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		cmd, err := saleBillAssembler.AssSaleItemCreateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.SaleItemCreate(ctx, cmd)
	})
}

// SaleItemCreateAndGet
// @Description: 添加明细
// @receiver c
// @param ctx
func (c *SaleItemCommandApi) SaleItemCreateAndGet(ictx iris.Context, tenantId string) {
	_ = restapp.Do(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleItemCreateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, tenantId, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.SaleItemCreate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryByIds(ctx, appCmd.GetTenantId(), appCmd.Data.GetIds())
		})

		return err
	})
}

// SaleItemUpdate
// @Description: 更新明细
// @receiver c
// @param ctx
func (c *SaleItemCommandApi) SaleItemUpdate(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		cmd, err := saleBillAssembler.AssSaleItemUpdateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.SaleItemUpdate(ctx, cmd)
	})
}

// SaleItemUpdateAndGet
// @Description: 更新明细
// @receiver c
// @param ctx
func (c *SaleItemCommandApi) SaleItemUpdateAndGet(ictx iris.Context, tenantId string) {
	_ = restapp.Do(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleItemUpdateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, tenantId, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.SaleItemUpdate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryByIds(ctx, appCmd.GetTenantId(), appCmd.Data.GetIds())
		})

		return err
	})
}
