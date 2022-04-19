package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/application/appservice"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserController struct {
	appQueryService *appservice.UserAppQueryService
}

func NewUserController() *UserController {
	return &UserController{
		appQueryService: appservice.NewUserAppQueryService(),
	}
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/{id}", "GetById")
	b.Handle("GET", "/tenants/{tenantId}/users", "GetPagingData")
}

func (m *UserController) GetById(ctx iris.Context, tenantId, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func() (interface{}, bool, error) {
		return m.appQueryService.FindById(ctx, tenantId, id)
	})
}

func (m *UserController) GetPagingData(ctx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ctx, func() (interface{}, bool, error) {
		query, err := restapp.NewListQuery(ctx, tenantId)
		if err != nil {
			return nil, false, err
		}
		return m.appQueryService.GetPagingData(ctx, query)
	})
}
