package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/sale_bill/assembler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var SaleBillAssembler = assembler.SaleBill

type SaleBillQueryApi struct {
	queryService *service.SaleBillQueryAppService
}

func NewSaleBillQueryApi() *SaleBillQueryApi {
	return &SaleBillQueryApi{
		queryService: service.GetSaleBillQueryAppService(),
	}
}

func (a *SaleBillQueryApi) BeforeActivation(b mvc.BeforeActivation) {
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-bills/{id}", "FindById")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-bills:all", "FindAll")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-bills:ids", "FindByIds")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/sale-bills", "FindPaging")
}

// FindById godoc
// @Summary      按ID查询
// @Description  get string by ID
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "SaleBill ID"
// @Success      200        {object}   dto.SaleBillFindByIdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/sale-bills/{id} [get]
func (a *SaleBillQueryApi) FindById(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQueryOne(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleBillAssembler.AssFindByIdRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		v, b, e := a.queryService.FindById(ctx, req.TenantId, req.Id)
		return SaleBillAssembler.AssFindByIdResponse(ictx, v, b, e)
	})
}

// FindByIds godoc
// @Summary      按多个ID获取<no value>
// @Description  get string by ID
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId  path      string     true    "Tenant ID"
// @Param        id        path      string     true    "SaleBill ID"
// @Success      200       {object}  dto.SaleBillFindByIdsResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/sale-bills:ids [get]
func (a *SaleBillQueryApi) FindByIds(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleBillAssembler.AssFindByIdsRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindByIds(ctx, req.TenantId, req.Ids)
		return SaleBillAssembler.AssFindByIdsResponse(ictx, fpr, b, e)
	})
}

// FindAll godoc
// @Summary      获取所有<no value>
// @Description  get string by ID
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId  path      int     true    "Tenant ID"
// @Success      200       {object}  dto.SaleBillFindAllResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/sale-bills:all [get]
func (a *SaleBillQueryApi) FindAll(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleBillAssembler.AssFindAllRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindAll(ctx, req.TenantId)
		return SaleBillAssembler.AssFindAllResponse(ictx, fpr, b, e)
	})
}

// FindPaging godoc
// @Summary      分页查询<no value>
// @Description  分页查询<no value>
// @Tags         sale-bills
// @Accept       json
// @Produce      json
// @Param        tenantId   path        int         true    "Tenant ID"
// @Success      200        {object}    dto.SaleBillFindPagingResponse
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/sale-bills [get]
func (a *SaleBillQueryApi) FindPaging(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		req, err := SaleBillAssembler.AssFindPagingRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindPaging(ctx, req)
		return SaleBillAssembler.AssFindPagingResponse(ictx, fpr, b, e)
	})
}
