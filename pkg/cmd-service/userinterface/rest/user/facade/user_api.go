package facade

import (
	"context"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/service"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/userinterface/rest/user/assembler"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

//
// UserCommandApi
// @Description:
//
type UserCommandApi struct {
	userAppService *service.UserCommandAppService
	queryAppId     string
}

var userAssembler = assembler.User

//
// NewUserController
// @Description:
// @return *UserCommandApi
//
func NewUserController() *UserCommandApi {
	return &UserCommandApi{
		userAppService: service.NewCommandUserAppService(),
		queryAppId:     service.GetUserQueryAppService().AppId(),
	}
}

func (c *UserCommandApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/aggregate/{id}", "FindAggregateById")

	b.Handle("POST", "/tenants/{tenantId}/users", "UserCreate")
	b.Handle("POST", "/tenants/{tenantId}/users:get", "UserCreateAndGet")
	b.Handle("PATCH", "/tenants/{tenantId}/users", "UserUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/users:get", "UserUpdateAndGet")

	b.Handle("POST", "/tenants/{tenantId}/users/addresses", "AddressCreate")
	b.Handle("POST", "/tenants/{tenantId}/users/addresses:get", "AddressCreateAndGet")
	b.Handle("PATCH", "/tenants/{tenantId}/users/addresses", "AddressUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/users/addresses:get", "AddressUpdateAndGet")
	b.Handle("DELETE", "/tenants/{tenantId}/users/addresses", "AddressDelete")
}

//
// FindAggregateById godoc
// @Summary      获取用户聚合根信息
// @Description  获取用户聚合根信息
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path      string  true  "租户ID"
// @Param        id         path      string  true  "聚合根ID"
// @Success      200  {object}  any
// @Failure      500  {object}  string
// @Router       /tenants/{tenantId}/aggregate/{id} [get]
func (c *UserCommandApi) FindAggregateById(ctx iris.Context, tenantId string, id string) {
	_, _, _ = restapp.DoQueryOne(ctx, func(ctx context.Context) (interface{}, bool, error) {
		return c.userAppService.GetAggregateById(ctx, tenantId, id)
	})
}

//
// UserCreate godoc
// @Summary      创建用户数据
// @Description  创建用户数据
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        data body dto.UserCreateRequest true "The input UserCreateRequest struct"
// @Success      200  {object}  any
// @Failure      500  {object}  string
// @Router       /tenants/{tenantId}/users  [post]
func (c *UserCommandApi) UserCreate(ctx iris.Context) {
	cmd, err := userAssembler.AssUserCreateCommandDto(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return
	}
	_ = restapp.DoCmd(ctx, func(ctx context.Context) error {
		return c.userAppService.CreateUser(ctx, cmd)
	})
}

//
// UserCreateAndGet godoc
// @Summary      创建并返回用户数据
// @Description  创建并返回用户数据
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        data body dto.UserCreateRequest true "The input UserUpdateRequest struct"
// @Success      200  {object}  any
// @Failure      500  {object}  string
// @Router       /tenants/{tenantId}/users:get  [post]
func (c *UserCommandApi) UserCreateAndGet(ctx iris.Context) {
	cmd, err := userAssembler.AssUserCreateCommandDto(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return
	}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.CreateUser(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

//
// UserUpdate godoc
// @Summary      更新用户数据
// @Description  更新用户数据
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        data body dto.UserUpdateRequest true "The input UserUpdate struct"
// @Success      200  {object}  any
// @Failure      500  {object}  string
// @Router       /tenants/{tenantId}/users:get  [patch]
func (c *UserCommandApi) UserUpdate(ctx iris.Context) {
	cmd, err := userAssembler.AssUserUpdateCommandDto(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return
	}
	_ = restapp.DoCmd(ctx, func(ctx context.Context) error {
		return c.userAppService.UpdateUser(ctx, cmd)
	})
}

//
// UserUpdateAndGet
// @Summary      更新并返回用户数据
// @Description  更新并返回用户数据
// @Tags         users
// @Param        data body dto.UserUpdateRequest true "The input UserUpdateRequest struct"
// @Accept       json
// @Produce      json
// @Success      200  {object}  any
// @Failure      500  {object}  string
// @Router       /tenants/{tenantId}/users:get  [patch]
func (c *UserCommandApi) UserUpdateAndGet(ctx iris.Context) {
	cmd, err := userAssembler.AssUserUpdateCommandDto(ctx)
	if err != nil {
		restapp.SetError(ctx, err)
		return
	}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.UpdateUser(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getById(ctx, cmd.GetTenantId(), cmd.Data.Id)
	})
}

func (c *UserCommandApi) AddressCreate(ctx iris.Context) {
	cmd := &command.AddressCreateCommand{}
	_ = restapp.DoCmd(ctx, func(ctx context.Context) error {
		return c.userAppService.CreateAddress(ctx, cmd)
	})
}

func (c *UserCommandApi) AddressCreateAndGet(ctx iris.Context) {
	cmd := &command.AddressCreateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.CreateAddress(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getById(ctx, cmd.GetTenantId(), cmd.Data.UserId)
	})
}

func (c *UserCommandApi) AddressUpdate(ctx iris.Context) {
	cmd := &command.AddressUpdateCommand{}
	_ = restapp.DoCmd(ctx, func(ctx context.Context) error {
		return c.userAppService.UpdateAddress(ctx, cmd)
	})
}

func (c *UserCommandApi) AddressUpdateAndGet(ctx iris.Context) {
	cmd := &command.AddressUpdateCommand{}
	_, _, _ = restapp.DoCmdAndQueryOne(ctx, c.queryAppId, cmd, func(ctx context.Context) error {
		return c.userAppService.UpdateAddress(ctx, cmd)
	}, func(ctx context.Context) (interface{}, bool, error) {
		return c.getById(ctx, cmd.GetTenantId(), cmd.Data.UserId)
	})
}

func (c *UserCommandApi) AddressDelete(ctx iris.Context) {
	cmd := &command.AddressDeleteCommand{}
	_ = restapp.DoCmd(ctx, func(ctx context.Context) error {
		return c.userAppService.DeleteAddress(ctx, cmd)
	})
}

func (c *UserCommandApi) getById(ctx context.Context, tenantId, userId string) (data interface{}, isFound bool, err error) {
	return service.GetUserQueryAppService().FindById(ctx, tenantId, userId)
}
