package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/application/internals/cmdappservice"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/application/internals/queryappservice"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserController struct {
	userAppService *cmdappservice.UserCommandAppService
}

func NewUserController() *UserController {
	return &UserController{
		userAppService: cmdappservice.NewCommandUserAppService(),
	}
}

func (m *UserController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/aggregate/{id}", "GetAggregateById")

	b.Handle("POST", "/tenants/{tenantId}/users", "UserCreate")
	b.Handle("POST", "/tenants/{tenantId}/users:get", "UserCreateAndGetUser")
	b.Handle("PATCH", "/tenants/{tenantId}/users", "UserUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/users:get", "UserUpdateAndGetUser")

	b.Handle("POST", "/tenants/{tenantId}/users/addresses", "AddressCreate")
	b.Handle("POST", "/tenants/{tenantId}/users/addresses:get", "AddressCreateAndGet")
	b.Handle("PATCH", "/tenants/{tenantId}/users/addresses", "AddressUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/users/addresses:get", "AddressUpdateAndGet")
	b.Handle("DELETE", "/tenants/{tenantId}/users/addresses", "AddressDelete")
}

func (m *UserController) GetAggregateById(ctx iris.Context, tenantId string, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func() (interface{}, bool, error) {
		return m.userAppService.GetAggregateById(NewContext(ctx), tenantId, id)
	})
}

func (m *UserController) UserCreate(ctx iris.Context) {
	cmd := &user_commands.UserCreateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func() error {
		return m.userAppService.CreateUser(NewContext(ctx), cmd)
	})
}

func (m *UserController) UserCreateAndGetUser(ctx iris.Context) {
	cmd := &user_commands.UserCreateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, subAppId, cmd, func() error {
		return m.userAppService.CreateUser(NewContext(ctx), cmd)
	}, func() (interface{}, bool, error) {
		return queryappservice.GetUserByUserId(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

func (m *UserController) UserUpdate(ctx iris.Context) {
	cmd := &user_commands.UserUpdateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func() error {
		return m.userAppService.UpdateUser(NewContext(ctx), cmd)
	})
}

func (m *UserController) UserUpdateAndGetUser(ctx iris.Context) {
	cmd := &user_commands.UserUpdateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, subAppId, cmd, func() error {
		return m.userAppService.UpdateUser(NewContext(ctx), cmd)
	}, func() (interface{}, bool, error) {
		return queryappservice.GetUserByUserId(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

func (m *UserController) AddressCreate(ctx iris.Context) {
	cmd := &user_commands.AddressCreateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func() error {
		return m.userAppService.CreateAddress(NewContext(ctx), cmd)
	})
}

func (m *UserController) AddressCreateAndGet(ctx iris.Context) {
	cmd := &user_commands.AddressCreateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, subAppId, cmd, func() error {
		return m.userAppService.CreateAddress(NewContext(ctx), cmd)
	}, func() (interface{}, bool, error) {
		return queryappservice.GetUserByUserId(ctx, cmd.GetTenantId(), cmd.Data.UserId)
	})
}

func (m *UserController) AddressUpdate(ctx iris.Context) {
	cmd := &user_commands.AddressUpdateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func() error {
		return m.userAppService.UpdateAddress(NewContext(ctx), cmd)
	})
}

func (m *UserController) AddressUpdateAndGet(ctx iris.Context) {
	cmd := &user_commands.AddressUpdateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, subAppId, cmd, func() error {
		return m.userAppService.UpdateAddress(NewContext(ctx), cmd)
	}, func() (interface{}, bool, error) {
		return queryappservice.GetUserByUserId(ctx, cmd.GetTenantId(), cmd.Data.UserId)
	})
}

func (m *UserController) AddressDelete(ctx iris.Context) {
	cmd := &user_commands.AddressDeleteCommand{}
	_ = restapp.DoCmd(ctx, cmd, func() error {
		return m.userAppService.DeleteAddress(NewContext(ctx), cmd)
	})
}
