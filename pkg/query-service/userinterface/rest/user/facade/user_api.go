package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/application/internals/service"
	base "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/userinterface/rest/facade"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/user/assembler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserApi struct {
	base.BaseApi
	appQueryService *service.UserQueryAppService
}

var userAssembler = assembler.User

func NewUserApi() *UserApi {
	return &UserApi{
		appQueryService: service.NewUserQueryAppService(),
	}
}

func (m *UserApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/{id}", "FindById")
	b.Handle("GET", "/tenants/{tenantId}/users:all", "FindAll")
	b.Handle("GET", "/tenants/{tenantId}/users", "FindPaging")
}

func (m *UserApi) FindById(ctx iris.Context, tenantId, id string) {
	req, err := userAssembler.AssFindByIdRequest(ctx)
	if err != nil {
		return
	}
	_, _, _ = restapp.DoQueryOne(ctx, func(c context.Context) (interface{}, bool, error) {
		v, b, e := m.appQueryService.FindById(c, req.TenantId(), req.Id())
		return userAssembler.AssOneResponse(ctx, v, b, e)
	})
}

func (m *UserApi) FindAll(ctx iris.Context, tenantId string) {
	req, err := userAssembler.AssFindAllRequest(ctx)
	if err != nil {
		return
	}
	_, _, _ = restapp.DoQuery(ctx, func(c context.Context) (interface{}, bool, error) {
		fpr, b, e := m.appQueryService.FindAll(c, req.TenantId())
		return userAssembler.AssListResponse(ctx, fpr, b, e)
	})
}

func (m *UserApi) FindPaging(ctx iris.Context, tenantId string) {
	req, err := userAssembler.AssGetPagingRequest(ctx)
	if err != nil {
		return
	}
	_, _, _ = restapp.DoQuery(ctx, func(c context.Context) (interface{}, bool, error) {
		fpr, b, e := m.appQueryService.FindPaging(c, req)
		return userAssembler.AssFindPagingResponse(ctx, fpr, b, e)
	})
}
