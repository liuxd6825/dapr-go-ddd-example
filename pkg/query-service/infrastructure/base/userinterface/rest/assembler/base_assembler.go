package utils

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/assert"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type BaseAssembler struct {
}

type FindRequest interface {
	SetTenantId(value string)
	GetTenantId() string
}

type FindByIdRequest interface {
	FindRequest
	SetId(value string)
	GetId() string
}

type FindAllRequest interface {
	FindRequest
}

type FindPagingRequest interface {
	FindRequest

	GetPageNum() int64
	SetPageNum(value int64)

	GetPageSize() int64
	SetPageSize(value int64)

	GetFilter() string
	SetFilter(value string)

	GetSort() string
	SetSort(value string)

	GetFields() string
	SetFields(value string)
}

type FindPagingResponse interface {
	GetPageNum() int64
	SetPageNum(value int64)

	GetPageSize() int64
	SetPageSize(value int64)

	GetFilter() string
	SetFilter(value string)

	GetSort() string
	SetSort(value string)

	GetFields() string
	SetFields(value string)
}

func (b *BaseAssembler) SetFindRequest(ctx iris.Context, r FindRequest) error {
	tenantId := ctx.Params().Get("tenantId")
	r.SetTenantId(tenantId)
	return assert.NotEmpty(tenantId, assert.NewOptions("url \"{tenantId}\" cannot be empty"))
}

func (b *BaseAssembler) SetFindByIdRequest(ctx iris.Context, r FindByIdRequest) error {
	if err := b.SetFindRequest(ctx, r); err != nil {
		return err
	}
	id := ctx.Params().Get("id")
	r.SetId(id)
	return assert.NotEmpty(id, assert.NewOptions("url \"{id}\" cannot be empty"))
}

func (b *BaseAssembler) SetFindAllRequest(ctx iris.Context, r FindAllRequest) error {
	return b.SetFindRequest(ctx, r)
}

func (b *BaseAssembler) SetFindPagingRequest(ctx iris.Context, r FindPagingRequest) error {
	if err := b.SetFindRequest(ctx, r); err != nil {
		return err
	}
	r.SetPageNum(ctx.URLParamInt64Default("pageNum", 0))
	r.SetPageSize(ctx.URLParamInt64Default("pageSize", 20))
	r.SetFilter(ctx.URLParamDefault("filter", ""))
	r.SetSort(ctx.URLParamDefault("sort", ""))
	return nil
}
func (b *BaseAssembler) SetFindPagingResponse(v *ddd_repository.FindPagingResult[*view.UserView], resp FindPagingResponse) {
	resp.SetFields(v.Filter)

}
