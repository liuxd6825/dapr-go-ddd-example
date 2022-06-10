package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	base "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/userinterface/rest/assembler"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/user/dto"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/mapper"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"time"
)

type userAssembler struct {
	base.BaseAssembler
}

var User = &userAssembler{}

func (a *userAssembler) AssFindByIdRequest(ctx iris.Context) (*dto.UserFindByIdRequest, error) {
	res := &dto.UserFindByIdRequest{}
	err := a.SetFindByIdRequest(ctx, res)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
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

func (a *userAssembler) AssFindPagingRequest(ctx iris.Context) (*dto.UserFindPagingRequest, error) {
	res := &dto.UserFindPagingRequest{}
	err := a.SetFindPagingRequest(ctx, res)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
}

func (a *userAssembler) AssFindPagingResponse(ctx iris.Context, v *ddd_repository.FindPagingResult[*view.UserView], isFound bool, findErr error) (*dto.UserFindPagingResponse, bool, error) {
	if findErr != nil || !isFound {
		return nil, isFound, findErr
	}

	var data []*dto.UserFindPagingResponseItem
	for i, item := range *v.Data {
		if i < 5 {
			now := time.Now()
			item.CreateTime = &now
		}
	}
	err := mapper.Mapper(v.Data, &data)

	res := &dto.UserFindPagingResponse{
		TotalPages: v.TotalPages,
		TotalRows:  v.TotalRows,
		Filter:     v.Filter,
		Sort:       v.Sort,
		PageNum:    v.PageNum,
		PageSize:   v.PageSize,
		Data:       &data,
	}
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, false, err
	}

	return res, isFound, nil
}

func (a *userAssembler) AssFindAllRequest(ctx iris.Context) (*dto.UserFindAllRequest, error) {
	res := &dto.UserFindAllRequest{}
	err := a.SetFindAllRequest(ctx, res)
	if err != nil {
		restapp.SetError(ctx, err)
		return nil, err
	}
	return res, nil
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
