package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/inventory/assembler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var inventoryAssembler = assembler.InventoryAssembler{}

type InventoryCommandApi struct {
	service *service.InventoryCommandAppService
}

func NewInventoryCommandApi() *InventoryCommandApi {
	return &InventoryCommandApi{
		service: service.NewInventoryCommandAppService(),
	}
}

func (c *InventoryCommandApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/inventories/aggregate/{id}", "FindAggregateById")
	b.Handle("POST", "/tenants/{tenantId}/inventories", "InventoryCreate")
	b.Handle("POST", "/tenants/{tenantId}/inventories:get", "InventoryCreateAndGet")
	b.Handle("PATCH", "/tenants/{tenantId}/inventories", "InventoryUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/inventories:get", "InventoryUpdateAndGet")
}

//
// FindAggregateById godoc
// @Summary      按聚合根ID查找聚合对象
// @Description  按聚合根ID查找聚合对象
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/inventories/aggregate/{id} [get]
//
func (c *InventoryCommandApi) FindAggregateById(ictx iris.Context, tenantId string, id string) {
	_, _, _ = restapp.DoQueryOne(ictx, func(ctx context.Context) (interface{}, bool, error) {
		return c.service.FindAggregateById(ctx, tenantId, id)
	})
}

//
// InventoryCreate godoc
// @Summary      创建存货档案
// @Description  创建存货档案
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/inventories [POST]
//
func (c *InventoryCommandApi) InventoryCreate(ictx iris.Context) {
	_ = restapp.DoCmd(ictx, func(ctx context.Context) error {
		appCmd, err := inventoryAssembler.AssInventoryCreateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.InventoryCreate(ctx, appCmd)
	})
}

//
// InventoryCreateAndGet godoc
// @Summary      创建存货档案
// @Description  创建存货档案
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/inventories:get [POST]
//
func (c *InventoryCommandApi) InventoryCreateAndGet(ictx iris.Context) {
	_ = restapp.Do(ictx, func() error {
		appCmd, err := inventoryAssembler.AssInventoryCreateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.InventoryCreate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
		})
		return err
	})
}

//
// InventoryUpdate godoc
// @Summary      更新存货档案
// @Description  更新存货档案
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/inventories [PATCH]
//
func (c *InventoryCommandApi) InventoryUpdate(ictx iris.Context) {
	_ = restapp.DoCmd(ictx, func(ctx context.Context) error {
		appCmd, err := inventoryAssembler.AssInventoryUpdateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.InventoryUpdate(ctx, appCmd)
	})
}

//
// InventoryUpdateAndGet godoc
// @Summary      更新存货档案
// @Description  更新存货档案
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/inventories:get [PATCH]
//
func (c *InventoryCommandApi) InventoryUpdateAndGet(ictx iris.Context) {
	_ = restapp.Do(ictx, func() error {
		appCmd, err := inventoryAssembler.AssInventoryUpdateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.InventoryUpdate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
		})
		return err
	})
}
