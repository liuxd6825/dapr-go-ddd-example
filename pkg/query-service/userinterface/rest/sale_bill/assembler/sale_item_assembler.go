package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/sale_bill/dto"
	"github.com/kataras/iris/v12"
)

type SaleItemAssembler struct {
	base.BaseAssembler
}

var SaleItem = &SaleItemAssembler{}

func (a *SaleItemAssembler) AssFindByIdResponse(ictx iris.Context, v *view.SaleItemView, isFound bool, findErr error) (*dto.SaleItemFindByIdResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	res := dto.NewSaleItemFindByIdResponse()
	err := utils.Mapper(v, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *SaleItemAssembler) AssFindPagingRequest(ctx iris.Context) (*appquery.SaleItemFindPagingAppQuery, error) {
	fpr, err := a.BaseAssembler.AssFindPagingRequest(ctx)
	if err != nil {
		return nil, err
	}
	query := &appquery.SaleItemFindPagingAppQuery{}
	query.Filter = fpr.Filter
	query.TenantId = fpr.TenantId
	query.Sort = fpr.Sort
	query.PageSize = fpr.PageSize
	query.PageNum = fpr.PageNum
	return query, nil
}

func (a *SaleItemAssembler) AssFindPagingResponse(ictx iris.Context, fpr *appquery.SaleItemFindPagingResult, isFound bool, findErr error) (*dto.SaleItemFindPagingResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	response := dto.NewSaleItemFindPagingResponse()
	err := utils.Mapper(fpr, response)
	if err != nil {
		return nil, false, err
	}
	return response, isFound, nil
}

func (a *SaleItemAssembler) AssFindAllResponse(ictx iris.Context, vList []*view.SaleItemView, isFound bool, findErr error) (*dto.SaleItemFindAllResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewSaleItemFindAllResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *SaleItemAssembler) AssFindBySaleBillIdRequest(ictx iris.Context) (*dto.SaleItemFindBySaleBillIdRequest, error) {
	tenantId, err := a.BaseAssembler.GetTenantId(ictx)
	if err != nil {
		return nil, err
	}

	saleBillId, err := a.BaseAssembler.GetIdParam(ictx, "saleBillId")
	if err != nil {
		return nil, err
	}

	res := dto.NewSaleItemFindBySaleBillIdRequest()
	res.TenantId = tenantId
	res.SaleBillId = saleBillId

	return res, nil
}

func (a *SaleItemAssembler) AssFindBySaleBillIdResponse(ictx iris.Context, vList []*view.SaleItemView, isFound bool, findErr error) (*dto.SaleItemFindBySaleBillIdResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}

	res := dto.NewSaleItemFindBySaleBillIdResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}
