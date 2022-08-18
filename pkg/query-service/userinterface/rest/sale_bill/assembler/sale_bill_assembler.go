package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/sale_bill/dto"
	"github.com/kataras/iris/v12"
)

type SaleBillAssembler struct {
	base.BaseAssembler
}

var SaleBill = &SaleBillAssembler{}

func (a *SaleBillAssembler) AssFindByIdResponse(ictx iris.Context, v *view.SaleBillView, isFound bool, findErr error) (*dto.SaleBillFindByIdResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	res := dto.NewSaleBillFindByIdResponse()
	err := utils.Mapper(v, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *SaleBillAssembler) AssFindPagingRequest(ctx iris.Context) (*appquery.SaleBillFindPagingAppQuery, error) {
	fpr, err := a.BaseAssembler.AssFindPagingRequest(ctx)
	if err != nil {
		return nil, err
	}
	query := &appquery.SaleBillFindPagingAppQuery{}
	query.Filter = fpr.Filter
	query.TenantId = fpr.TenantId
	query.Sort = fpr.Sort
	query.PageSize = fpr.PageSize
	query.PageNum = fpr.PageNum
	query.Fields = fpr.Fields
	return query, nil
}

func (a *SaleBillAssembler) AssFindPagingResponse(ictx iris.Context, fpr *appquery.SaleBillFindPagingResult, isFound bool, findErr error) (*dto.SaleBillFindPagingResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewSaleBillFindPagingResponse()
	err := utils.Mapper(fpr, res)
	if err != nil {
		return nil, false, err
	}
	return res, isFound, nil
}

func (a *SaleBillAssembler) AssFindByIdsResponse(ictx iris.Context, vList []*view.SaleBillView, isFound bool, findErr error) (*dto.SaleBillFindByIdsResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewSaleBillFindByIdsResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *SaleBillAssembler) AssFindAllResponse(ictx iris.Context, vList []*view.SaleBillView, isFound bool, findErr error) (*dto.SaleBillFindAllResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewSaleBillFindAllResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}
