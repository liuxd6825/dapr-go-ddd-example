package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/sale_bill/assembler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var SaleItemAssembler = assembler.SaleItem

type SaleItemQueryApi struct {
	queryService *service.SaleItemQueryAppService
}

func NewSaleItemQueryApi() *SaleItemQueryApi {
	return &SaleItemQueryApi{
		queryService: service.GetSaleItemQueryAppService(),
	}
}

func (a *SaleItemQueryApi) BeforeActivation(b mvc.BeforeActivation) {
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-items/{id}", "FindById")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-bills/{saleBillId}/sale-items", "FindBySaleBillId")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-items:sale-bill-id/sale-items", "FindBySaleBillId")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-items:all", "FindAll")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-items", "FindPaging")
}

// FindById godoc
// @Summary      按ID获取用户
// @Description  get string by ID
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "User ID"
// @Success      200        {object}   dto.SaleItemFindByIdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/sale-bills/{id} [get]
func (a *SaleItemQueryApi) FindById(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQueryOne(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleItemAssembler.AssFindByIdRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		v, b, e := a.queryService.FindById(ctx, req.TenantId, req.Id)
		return SaleItemAssembler.AssFindByIdResponse(ictx, v, b, e)
	})
}

// FindBySaleBillId godoc
// @Summary      按SaleBillId获取
// @Description  get string by ID
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "User ID"
// @Success      200        {object}   dto.SaleBillFindBySaleBillIdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/sale-items:sale-bill-id/saleBillId [get]
func (a *SaleItemQueryApi) FindBySaleBillId(ictx iris.Context, tenantId, saleBillId string) {
	_, _, _ = restapp.DoQuery(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleItemAssembler.AssFindBySaleBillIdRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		v, b, e := a.queryService.FindBySaleBillId(ctx, req.TenantId, req.SaleBillId)
		return SaleItemAssembler.AssFindBySaleBillIdResponse(ictx, v, b, e)
	})
}

// FindAll godoc
// @Summary      获取所有
// @Description  get string by ID
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId  path      int     true    "Tenant ID"
// @Success      200       {object}  dto.UserFindAllResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/sale-bills:all [get]
func (a *SaleItemQueryApi) FindAll(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleItemAssembler.AssFindAllRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindAll(ctx, req.TenantId)
		return SaleItemAssembler.AssFindAllResponse(ictx, fpr, b, e)
	})
}

// FindPaging godoc
// @Summary      分页查询
// @Description  分页查询
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        int         true    "Tenant ID"
// @Success      200        {object}    dto.SaleItemFindPagingResponse
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills [get]
func (a *SaleItemQueryApi) FindPaging(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleItemAssembler.AssFindPagingRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindPaging(ctx, req)
		return SaleItemAssembler.AssFindPagingResponse(ictx, fpr, b, e)
	})
}
