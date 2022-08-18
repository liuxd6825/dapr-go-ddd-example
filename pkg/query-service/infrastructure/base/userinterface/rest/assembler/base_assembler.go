package assembler

import (
	"errors"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/dto"
	"github.com/kataras/iris/v12"
)

type BaseAssembler struct {
}

func (a *BaseAssembler) AssFindByIdRequest(ictx iris.Context) (*dto.FindByIdRequest, error) {
	tenantId, err := a.GetTenantId(ictx)
	if err != nil {
		return nil, err
	}
	id, err := a.GetId(ictx)
	if err != nil {
		return nil, err
	}
	return &dto.FindByIdRequest{
		TenantId: tenantId,
		Id:       id,
	}, nil
}

func (a *BaseAssembler) AssFindByIdsRequest(ictx iris.Context) (*dto.FindByIdsRequest, error) {
	tenantId, err := a.GetTenantId(ictx)
	if err != nil {
		return nil, err
	}
	ids, err := a.GetIds(ictx)
	if err != nil {
		return nil, err
	}
	return &dto.FindByIdsRequest{
		TenantId: tenantId,
		Ids:      ids,
	}, nil
}

func (a *BaseAssembler) AssFindAllRequest(ictx iris.Context) (*dto.FindAllRequest, error) {
	tenantId, err := a.GetTenantId(ictx)
	if err != nil {
		return nil, err
	}
	return &dto.FindAllRequest{
		TenantId: tenantId,
	}, nil
}

func (a *BaseAssembler) AssFindPagingRequest(ictx iris.Context) (*dto.FindPagingRequest, error) {
	tenantId, err := a.GetTenantId(ictx)
	if err != nil {
		return nil, err
	}
	pageNum := ictx.URLParamInt64Default("page-num", 0)
	pageSize := ictx.URLParamInt64Default("page-size", 20)
	filter := ictx.URLParamDefault("filter", "")
	sort := ictx.URLParamDefault("sort", "")
	fields := ictx.URLParamDefault("fields", "")
	return &dto.FindPagingRequest{
		TenantId: tenantId,
		PageNum:  pageNum,
		PageSize: pageSize,
		Filter:   filter,
		Sort:     sort,
		Fields:   fields,
	}, nil
}

func (a *BaseAssembler) GetTenantId(ictx iris.Context) (string, error) {
	return a.GetIdParam(ictx, "tenantId")
}

func (a *BaseAssembler) GetId(ictx iris.Context) (string, error) {
	return a.GetIdParam(ictx, "id")
}

func (a *BaseAssembler) GetIds(ictx iris.Context) ([]string, error) {
	ids := ictx.URLParamSlice("id")
	return ids, nil
}

func (a *BaseAssembler) GetIdParam(ictx iris.Context, name string) (string, error) {
	id := ictx.Params().GetStringDefault(name, "")
	if id == "" {
		return "", errors.New(name + " is not empty")
	}
	return id, nil
}
