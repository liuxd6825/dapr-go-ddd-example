package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	base "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/user/dto"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type userAssembler struct {
	base.BaseAssembler
}

var User = &userAssembler{}

func (a *userAssembler) AssFindByIdRequest(ctx iris.Context, tenantId, id string) (*dto.UserFindByIdRequest, error) {
	request := dto.UserFindByIdRequest{
		TenantId: tenantId,
		Id:       id,
	}
	return &request, nil
}

func (a *userAssembler) AssFindByIdResponse(ctx iris.Context, v *view.UserView, isFound bool, findErr error) (*dto.UserFindByIdResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	res := &dto.UserFindByIdResponse{}
	err := mapper.Mapper(v, res)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, false, err
	}
	return res, true, nil
}

func (a *userAssembler) AssFindPagingRequest(ctx iris.Context, tenantId string) (ddd_repository.FindPagingQuery, error) {
	var request dto.UserFindPagingRequest
	query := a.GetFindPagingQuery(ctx)
	if err := mapper.Mapper(query, &request); err != nil {
		return nil, err
	}
	ddd_repository.NewFindPagingQuery()
	return &request, nil
}

func (a *userAssembler) AssFindPagingResponse(ctx iris.Context, v *ddd_repository.FindPagingResult[*view.UserView], isFound bool, findErr error) (*dto.UserFindPagingResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}

	var response dto.UserFindPagingResponse
	err := mapper.Mapper(v, &response)

	if err != nil {
		restapp.SetError(ctx, err)
		return nil, false, err
	}

	return &response, isFound, nil
}

func (a *userAssembler) AssFindAllRequest(ctx iris.Context) (*dto.UserFindAllRequest, error) {
	var request dto.UserFindAllRequest
	query := a.GetFindAllQuery(ctx)
	if err := mapper.Mapper(query, &request); err != nil {
		return nil, err
	}
	return &request, nil
}

func (a *userAssembler) AssFindAllResponse(ctx iris.Context, vList *[]*view.UserView, isFound bool, findErr error) (*dto.UserFindAllResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}

	res := &dto.UserFindAllResponse{}
	err := mapper.Mapper(vList, res)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, false, err
	}
	return res, true, nil
}
