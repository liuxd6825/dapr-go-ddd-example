package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/inventory/assembler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var InventoryAssembler = assembler.Inventory

type InventoryQueryApi struct {
	queryService *service.InventoryQueryAppService
}

func NewInventoryQueryApi() *InventoryQueryApi {
	return &InventoryQueryApi{
		queryService: service.GetInventoryQueryAppService(),
	}
}

func (a *InventoryQueryApi) BeforeActivation(b mvc.BeforeActivation) {
	restapp.Handle(b, "GET", "/tenants/{tenantId}/inventories/{id}", "FindById")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/inventories:all", "FindAll")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/inventories:ids", "FindByIds")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/inventories", "FindPaging")
}

// FindById godoc
// @Summary      按ID查询
// @Description  get string by ID
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "Inventory ID"
// @Success      200        {object}   dto.InventoryFindByIdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/inventories/{id} [get]
func (a *InventoryQueryApi) FindById(ictx iris.Context) {
	_, _, _ = restapp.DoQueryOne(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := InventoryAssembler.AssFindByIdRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		v, b, e := a.queryService.FindById(ctx, req.TenantId, req.Id)
		return InventoryAssembler.AssFindByIdResponse(ictx, v, b, e)
	})
}

// FindByIds godoc
// @Summary      按多个ID获取<no value>
// @Description  get string by ID
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId  path      string     true    "Tenant ID"
// @Param        id        path      string     true    "Inventory ID"
// @Success      200       {object}  dto.InventoryFindByIdsResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/inventories:ids [get]
func (a *InventoryQueryApi) FindByIds(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := InventoryAssembler.AssFindByIdsRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindByIds(ctx, req.TenantId, req.Ids)
		return InventoryAssembler.AssFindByIdsResponse(ictx, fpr, b, e)
	})
}

// FindAll godoc
// @Summary      获取所有<no value>
// @Description  get string by ID
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId  path      int     true    "Tenant ID"
// @Success      200       {object}  dto.InventoryFindAllResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/inventories:all [get]
func (a *InventoryQueryApi) FindAll(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := InventoryAssembler.AssFindAllRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindAll(ctx, req.TenantId)
		return InventoryAssembler.AssFindAllResponse(ictx, fpr, b, e)
	})
}

// FindPaging godoc
// @Summary      分页查询<no value>
// @Description  分页查询<no value>
// @Tags         inventories
// @Accept       json
// @Produce      json
// @Param        tenantId   path        int         true    "Tenant ID"
// @Success      200        {object}    dto.InventoryFindPagingResponse
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/inventories [get]
func (a *InventoryQueryApi) FindPaging(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := InventoryAssembler.AssFindPagingRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindPaging(ctx, req)
		return InventoryAssembler.AssFindPagingResponse(ictx, fpr, b, e)
	})
}
