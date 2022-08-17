package assembler

import (
	"errors"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/dto"
	"github.com/kataras/iris/v12"
)

type BaseAssembler struct {
}

func (a *BaseAssembler) AssFindByIdRequest(ctx iris.Context) (*dto.FindByIdRequest, error) {
	tenantId, err := a.GetTenantId(ctx)
	if err != nil {
		return nil, err
	}
	id, err := a.GetId(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.FindByIdRequest{
		TenantId: tenantId,
		Id:       id,
	}, nil
}

func (a *BaseAssembler) AssFindAllRequest(ctx iris.Context) (*dto.FindAllRequest, error) {
	tenantId, err := a.GetTenantId(ctx)
	if err != nil {
		return nil, err
	}
	return &dto.FindAllRequest{
		TenantId: tenantId,
	}, nil
}

func (a *BaseAssembler) AssFindPagingRequest(ctx iris.Context) (*dto.FindPagingRequest, error) {
	tenantId, err := a.GetTenantId(ctx)
	if err != nil {
		return nil, err
	}
	pageNum := ctx.URLParamInt64Default("page-num", 0)
	pageSize := ctx.URLParamInt64Default("page-size", 20)
	filter := ctx.URLParamDefault("filter", "")
	sort := ctx.URLParamDefault("sort", "")
	fields := ctx.URLParamDefault("fields", "")
	return &dto.FindPagingRequest{
		TenantId: tenantId,
		PageNum:  pageNum,
		PageSize: pageSize,
		Filter:   filter,
		Sort:     sort,
		Fields:   fields,
	}, nil
}

func (a *BaseAssembler) GetTenantId(ctx iris.Context) (string, error) {
	return a.GetIdParam(ctx, "tenantId")
}

func (a *BaseAssembler) GetId(ctx iris.Context) (string, error) {
	return a.GetIdParam(ctx, "id")
}

func (a *BaseAssembler) GetIdParam(ctx iris.Context, name string) (string, error) {
	id := ctx.Params().GetStringDefault(name, "")
	if id == "" {
		return "", errors.New(name + " is not empty")
	}
	return id, nil
}
