package controller

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/cmdappservice"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/queryappservice"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// UserController
// @Description:
//
type UserController struct {
	userAppService *cmdappservice.UserCommandAppService
	queryAppId     string
}

func NewUserController() *UserController {
	return &UserController{
		userAppService: cmdappservice.NewCommandUserAppService(),
		queryAppId:     queryappservice.GetUserQueryAppService().AppId(),
	}
}

func (c *UserController) BeforeActivation(b mvc.BeforeActivation) {
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

func (c *UserController) GetAggregateById(ctx iris.Context, tenantId string, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return c.userAppService.GetAggregateById(ctx, tenantId, id)
	})
}

func (c *UserController) UserCreate(ctx iris.Context) {
	cmd := &user_commands.UserCreateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.userAppService.CreateUser(ctx, cmd)
	})
}

func (c *UserController) UserCreateAndGetUser(ctx iris.Context) {
	cmd := &user_commands.UserCreateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.CreateUser(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getUserById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

func (c *UserController) UserUpdate(ctx iris.Context) {
	cmd := &user_commands.UserUpdateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.userAppService.UpdateUser(ctx, cmd)
	})
}

func (c *UserController) UserUpdateAndGetUser(ctx iris.Context) {
	cmd := &user_commands.UserUpdateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.UpdateUser(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getUserById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

func (c *UserController) AddressCreate(ctx iris.Context) {
	cmd := &user_commands.AddressCreateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.userAppService.CreateAddress(ctx, cmd)
	})
}

func (c *UserController) AddressCreateAndGet(ctx iris.Context) {
	cmd := &user_commands.AddressCreateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.CreateAddress(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getUserById(ctx, cmd.GetTenantId(), cmd.Data.UserId)
	})
}

func (c *UserController) AddressUpdate(ctx iris.Context) {
	cmd := &user_commands.AddressUpdateCommand{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.userAppService.UpdateAddress(ctx, cmd)
	})
}

func (c *UserController) AddressUpdateAndGet(ctx iris.Context) {
	cmd := &user_commands.AddressUpdateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.UpdateAddress(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getUserById(ctx, cmd.GetTenantId(), cmd.Data.UserId)
	})
}

func (c *UserController) AddressDelete(ctx iris.Context) {
	cmd := &user_commands.AddressDeleteCommand{}
	_ = restapp.DoCmd(ctx, cmd, func(ctx context.Context) error {
		return c.userAppService.DeleteAddress(ctx, cmd)
	})
}

func (c *UserController) getUserById(ctx context.Context, tenantId, userId string) (data interface{}, isFound bool, err error) {
	return queryappservice.GetUserQueryAppService().GetById(ctx, tenantId, userId)
}
