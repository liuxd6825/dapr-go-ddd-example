package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/inventory/dto"
	"github.com/kataras/iris/v12"
)

type InventoryAssembler struct {
	base.BaseAssembler
}

var Inventory = &InventoryAssembler{}

func (a *InventoryAssembler) AssFindByIdResponse(ictx iris.Context, v *view.InventoryView, isFound bool, findErr error) (*dto.InventoryFindByIdResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	res := dto.NewInventoryFindByIdResponse()
	err := utils.Mapper(v, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *InventoryAssembler) AssFindPagingRequest(ctx iris.Context) (*appquery.InventoryFindPagingAppQuery, error) {
	fpr, err := a.BaseAssembler.AssFindPagingRequest(ctx)
	if err != nil {
		return nil, err
	}
	query := &appquery.InventoryFindPagingAppQuery{}
	query.Filter = fpr.Filter
	query.TenantId = fpr.TenantId
	query.Sort = fpr.Sort
	query.PageSize = fpr.PageSize
	query.PageNum = fpr.PageNum
	query.Fields = fpr.Fields
	return query, nil
}

func (a *InventoryAssembler) AssFindPagingResponse(ictx iris.Context, fpr *appquery.InventoryFindPagingResult, isFound bool, findErr error) (*dto.InventoryFindPagingResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewInventoryFindPagingResponse()
	err := utils.Mapper(fpr, res)
	if err != nil {
		return nil, false, err
	}
	return res, isFound, nil
}

func (a *InventoryAssembler) AssFindByIdsResponse(ictx iris.Context, vList []*view.InventoryView, isFound bool, findErr error) (*dto.InventoryFindByIdsResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewInventoryFindByIdsResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *InventoryAssembler) AssFindAllResponse(ictx iris.Context, vList []*view.InventoryView, isFound bool, findErr error) (*dto.InventoryFindAllResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewInventoryFindAllResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}
