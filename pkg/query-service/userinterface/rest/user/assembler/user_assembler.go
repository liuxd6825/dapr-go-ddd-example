package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/types"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/user/dto"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type userAssembler struct {
}

var User = &userAssembler{}

func (a *userAssembler) AssFindByIdRequest(ctx iris.Context) (*dto.UserFindByIdRequest, error) {
	res := &dto.UserFindByIdRequest{}
	err := res.Init(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
}

func (a *userAssembler) AssGetPagingRequest(ctx iris.Context) (*dto.UserFindPagingRequest, error) {
	res := &dto.UserFindPagingRequest{}
	err := res.Init(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
}

func (a *userAssembler) AssFindPagingResponse(ctx iris.Context, v *ddd_repository.FindPagingResult[*projection.UserView], isFound bool, findErr error) (*dto.UserFindPagingResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}
	dtoList := a.AssViewDtoList(v.Data)
	res := &dto.UserFindPagingResponse{}
	err := res.Init(v, dtoList)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, false, err
	}

	return res, isFound, nil
}

func (a *userAssembler) AssFindAllRequest(ctx iris.Context) (*dto.UserFindAllRequest, error) {
	res := &dto.UserFindAllRequest{}
	err := res.Init(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
}

func (a *userAssembler) AssOneResponse(ctx iris.Context, v *projection.UserView, isFound bool, err error) (*dto.UserViewDto, bool, error) {
	if err != nil || !isFound {
		return nil, isFound, err
	}
	res := a.AssViewDto(v)
	return res, isFound, nil
}

func (a *userAssembler) AssListResponse(ctx iris.Context, v *[]*projection.UserView, isFound bool, findErr error) (*[]*dto.UserViewDto, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}

	res := a.AssViewDtoList(v)
	return res, isFound, nil
}

func (a *userAssembler) AssViewDtoList(vList *[]*projection.UserView) *[]*dto.UserViewDto {
	var dtoList []*dto.UserViewDto
	if vList != nil {
		for _, v := range *vList {
			vDto := a.AssViewDto(v)
			dtoList = append(dtoList, vDto)
		}
	}
	return &dtoList
}

func (a *userAssembler) AssViewDto(v *projection.UserView) *dto.UserViewDto {
	d := dto.UserViewDto{}
	d.Id = v.Id
	d.TenantId = v.TenantId
	d.UserCode = v.UserCode
	d.UserName = v.UserName
	d.Telephone = v.Telephone
	d.UserName = v.UserName
	d.Address = v.Address
	d.Email = v.Email
	d.CreateTime = types.NewDateTime(v.CreateTime)
	return &d
}
