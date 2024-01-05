package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/sale_bill/assembler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var saleBillAssembler = assembler.SaleBillAssembler{}

type SaleBillCommandApi struct {
	service *service.SaleBillCommandAppService
}

func NewSaleBillCommandApi() *SaleBillCommandApi {
	return &SaleBillCommandApi{
		service: service.NewSaleBillCommandAppService(),
	}
}

func (c *SaleBillCommandApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/sale-bills/aggregate/{id}", "FindAggregateById")
	b.Handle("PATCH", "/tenants/{tenantId}/sale-bills:sale-bill-confirm", "SaleBillConfirm")
	b.Handle("POST", "/tenants/{tenantId}/sale-bills", "SaleBillCreate")
	b.Handle("POST", "/tenants/{tenantId}/sale-bills:get", "SaleBillCreateAndGet")
	b.Handle("DELETE", "/tenants/{tenantId}/sale-bills", "SaleBillDelete")
	b.Handle("PATCH", "/tenants/{tenantId}/sale-bills", "SaleBillUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/sale-bills:get", "SaleBillUpdateAndGet")
}

// FindAggregateById godoc
// @Summary      按聚合根ID查找聚合对象
// @Description  按聚合根ID查找聚合对象
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills/aggregate/{id} [get]
func (c *SaleBillCommandApi) FindAggregateById(ictx iris.Context, tenantId string, id string) {
	_, _, _ = restapp.DoQueryOne(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		return c.service.FindAggregateById(ctx, tenantId, id)
	})
}

// SaleBillConfirm godoc
// @Summary      下单确认命令
// @Description  下单确认命令
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills [PATCH]
func (c *SaleBillCommandApi) SaleBillConfirm(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleBillConfirmAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.SaleBillConfirm(ctx, appCmd)
	})
}

// SaleBillConfirmAndGet godoc
// @Summary      下单确认命令
// @Description  下单确认命令
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills:get [PATCH]
func (c *SaleBillCommandApi) SaleBillConfirmAndGet(ictx iris.Context, tenantId string) {
	_ = restapp.Do(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleBillConfirmAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, tenantId, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.SaleBillConfirm(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
		})
		return err
	})
}

// SaleBillCreate godoc
// @Summary      创建销售订单
// @Description  创建销售订单
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills [POST]
func (c *SaleBillCommandApi) SaleBillCreate(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleBillCreateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.SaleBillCreate(ctx, appCmd)
	})
}

// SaleBillCreateAndGet godoc
// @Summary      创建销售订单
// @Description  创建销售订单
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills:get [POST]
func (c *SaleBillCommandApi) SaleBillCreateAndGet(ictx iris.Context, tenantId string) {
	_ = restapp.Do(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleBillCreateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, tenantId, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.SaleBillCreate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
		})
		return err
	})
}

// SaleBillDelete godoc
// @Summary      删除销售订单
// @Description  删除销售订单
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills [DELETE]
func (c *SaleBillCommandApi) SaleBillDelete(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleBillDeleteAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.SaleBillDelete(ctx, appCmd)
	})
}

// SaleBillUpdate godoc
// @Summary      更新销售订单
// @Description  更新销售订单
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills [PATCH]
func (c *SaleBillCommandApi) SaleBillUpdate(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleBillUpdateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.SaleBillUpdate(ctx, appCmd)
	})
}

// SaleBillUpdateAndGet godoc
// @Summary      更新销售订单
// @Description  更新销售订单
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills:get [PATCH]
func (c *SaleBillCommandApi) SaleBillUpdateAndGet(ictx iris.Context, tenantId string) {
	_ = restapp.Do(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := saleBillAssembler.AssSaleBillUpdateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, tenantId, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.SaleBillUpdate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
		})
		return err
	})
}
