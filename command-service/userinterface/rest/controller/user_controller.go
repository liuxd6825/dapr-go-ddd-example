package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/application/internals/cmdappservice"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/application/internals/queryappservice"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-sdk/rest"
)

type UserController struct {
	userAppService *cmdappservice.UserCommandAppService
}

func NewUserController() *UserController {
	return &UserController{
		userAppService: cmdappservice.NewUserAppService(),
	}
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/aggregate/{id}", "GetAggregateById")
	b.Handle("POST", "/tenants/{tenantId}/users", "CreateUser")
	b.Handle("POST", "/tenants/{tenantId}/users:get", "CreateUserAndGetUser")
	b.Handle("PATCH", "/tenants/{tenantId}/users", "UpdateUser")
	b.Handle("PATCH", "/tenants/{tenantId}/users:get", "UpdateUserAndGetUser")
}

func (m *UserController) CreateUser(ctx iris.Context) {
	cmd := &user_commands.UserCreateCommand{}
	_, _ = rest.DoCmd(ctx, cmd, func() (interface{}, error) {
		if err := m.userAppService.CreateUser(NewContext(ctx), cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	})
}

func (m *UserController) CreateUserAndGetUser(ctx iris.Context) {
	cmd := &user_commands.UserCreateCommand{}
	_, _, _ = rest.DoCmdAndQueryOne(ctx, subAppId, cmd, func() (interface{}, error) {
		if err := m.userAppService.CreateUser(NewContext(ctx), cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	}, func() (interface{}, bool, error) {
		return queryappservice.GetUserByUserId(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

func (m *UserController) UpdateUser(ctx iris.Context) {
	cmd := &user_commands.UserUpdateCommand{}
	_, _ = rest.DoCmd(ctx, cmd, func() (interface{}, error) {
		if err := m.userAppService.UpdateUser(NewContext(ctx), cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	})
}

func (m *UserController) UpdateUserAndGetUser(ctx iris.Context) {
	cmd := &user_commands.UserUpdateCommand{}
	_, _, _ = rest.DoCmdAndQueryOne(ctx, subAppId, cmd, func() (interface{}, error) {
		if err := m.userAppService.UpdateUser(NewContext(ctx), cmd); err != nil {
			return nil, err
		}
		return cmd, nil
	}, func() (interface{}, bool, error) {
		return queryappservice.GetUserByUserId(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

func (m *UserController) GetAggregateById(ctx iris.Context, tenantId string, id string) {
	_, _, _ = rest.DoQueryOne(ctx, func() (interface{}, bool, error) {
		return m.userAppService.GetAggregateById(NewContext(ctx), tenantId, id)
	})
}
