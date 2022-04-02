package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/common/rest"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/application/appservice"
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
	b.Handle("GET", "/tenants/{tenantId}/users", "GetList")
}

func (m *UserController) GetById(ctx iris.Context, tenantId, id string) {
	rest.ResultOneData(ctx, func(ctx iris.Context) (interface{}, bool, error) {
		return m.appQueryService.FindById(ctx, tenantId, id)
	})
}

func (m UserController) GetList(ctx iris.Context) {

}
