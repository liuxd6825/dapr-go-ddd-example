package controller

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/application/internals/query_appservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type UserController struct {
	appQueryService *query_appservice.UserAppQueryService
}

func NewUserController() *UserController {
	return &UserController{
		appQueryService: query_appservice.NewUserAppQueryService(),
	}
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/{id}", "GetById")
	b.Handle("GET", "/tenants/{tenantId}/users", "GetPagingData")
}

func (m *UserController) GetById(ctx iris.Context, tenantId, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return m.appQueryService.FindById(ctx, tenantId, id)
	})
}

func (m *UserController) GetPagingData(ctx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ctx, func(ctx context.Context) (interface{}, bool, error) {
		query, err := restapp.NewListQuery(ctx, tenantId)
		if err != nil {
			return nil, false, err
		}
		return m.appQueryService.FindPagingData(ctx, query)
	})
}
