package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/user/dto"
	"github.com/kataras/iris/v12"
)

type UserAssembler struct {
	base.BaseAssembler
}

var User = &UserAssembler{}

func (a *UserAssembler) AssFindByIdResponse(ictx iris.Context, v *view.UserView, isFound bool, findErr error) (*dto.UserFindByIdResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	res := dto.NewUserFindByIdResponse()
	err := utils.Mapper(v, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *UserAssembler) AssFindPagingRequest(ctx iris.Context) (*appquery.UserFindPagingAppQuery, error) {
	fpr, err := a.BaseAssembler.AssFindPagingRequest(ctx)
	if err != nil {
		return nil, err
	}
	query := &appquery.UserFindPagingAppQuery{}
	query.Filter = fpr.Filter
	query.TenantId = fpr.TenantId
	query.Sort = fpr.Sort
	query.PageSize = fpr.PageSize
	query.PageNum = fpr.PageNum
	query.Fields = fpr.Fields
	return query, nil
}

func (a *UserAssembler) AssFindPagingResponse(ictx iris.Context, fpr *appquery.UserFindPagingResult, isFound bool, findErr error) (*dto.UserFindPagingResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewUserFindPagingResponse()
	err := utils.Mapper(fpr, res)
	if err != nil {
		return nil, false, err
	}
	return res, isFound, nil
}

func (a *UserAssembler) AssFindByIdsResponse(ictx iris.Context, vList []*view.UserView, isFound bool, findErr error) (*dto.UserFindByIdsResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewUserFindByIdsResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}

func (a *UserAssembler) AssFindAllResponse(ictx iris.Context, vList []*view.UserView, isFound bool, findErr error) (*dto.UserFindAllResponse, bool, error) {
	if findErr != nil {
		return nil, isFound, findErr
	}
	res := dto.NewUserFindAllResponse()
	err := utils.Mapper(vList, res)
	if err != nil {
		return nil, false, err
	}
	return res, true, nil
}
