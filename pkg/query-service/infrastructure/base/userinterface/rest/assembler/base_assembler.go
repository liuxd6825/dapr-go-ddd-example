package utils

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/userinterface/rest/dto"
)

type BaseAssembler struct {
}

func (a *BaseAssembler) GetTenantId(ctx iris.Context) string {
	return ctx.Params().GetStringDefault("tenantId", "")
}

func (a *BaseAssembler) GetFindByIdQuery(ctx iris.Context) *dto.FindByIdQuery {
	tenantId := a.GetTenantId(ctx)
	id := ctx.URLParamDefault("id", "")
	return &dto.FindByIdQuery{
		TenantId: tenantId,
		Id:       id,
	}
}

func (a *BaseAssembler) GetFindAllQuery(ctx iris.Context) *dto.FindAllQuery {
	tenantId := a.GetTenantId(ctx)
	return &dto.FindAllQuery{
		TenantId: tenantId,
	}
}

func (a *BaseAssembler) GetFindPagingQuery(ctx iris.Context) *dto.FindPagingQuery {
	tenantId := a.GetTenantId(ctx)
	pageNum := ctx.URLParamInt64Default("page-num", 0)
	pageSize := ctx.URLParamInt64Default("page-size", 20)
	filter := ctx.URLParamDefault("filter", "")
	sort := ctx.URLParamDefault("sort", "")
	fields := ctx.URLParamDefault("fields", "")
	return &dto.FindPagingQuery{
		TenantId: tenantId,
		PageNum:  pageNum,
		PageSize: pageSize,
		Filter:   filter,
		Sort:     sort,
		Fields:   fields,
	}
}
