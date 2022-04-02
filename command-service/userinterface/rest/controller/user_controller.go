package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/application/appservice"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-sdk/rest"
)

type UserController struct {
	userAppService *appservice.UserAppService
}

func NewUserController() *UserController {
	return &UserController{
		userAppService: appservice.NewUserAppService(),
	}
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/aggregate/{id}", "GetAggregateById")
	b.Handle("POST", "/tenants/{tenantId}/users", "CreateUser")
	b.Handle("PATCH", "/tenants/{tenantId}/users", "UpdateUser")
}

func (m *UserController) CreateUser(ctx iris.Context) {
	rest.Result(ctx, func(ctx iris.Context) (interface{}, error) {
		cmd := &user_commands.UserCreateCommand{}
		if err := ctx.ReadBody(cmd); err != nil {
			return nil, err
		}
		if err := m.userAppService.CreateUser(NewContext(ctx), cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	})
}

func (m *UserController) UpdateUser(ctx iris.Context) {
	rest.Result(ctx, func(ctx iris.Context) (interface{}, error) {
		cmd := &user_commands.UserUpdateCommand{}
		if err := ctx.ReadBody(cmd); err != nil {
			return nil, err
		}
		if err := m.userAppService.UpdateUser(NewContext(ctx), cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	})
}

func (m *UserController) GetAggregateById(ctx iris.Context, tenantId string, id string) {
	rest.ResultOneData(ctx, func(ctx iris.Context) (interface{}, bool, error) {
		return m.userAppService.GetAggregateById(NewContext(ctx), tenantId, id)
	})
}
