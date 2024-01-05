package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/user/assembler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var userAssembler = assembler.UserAssembler{}

type UserCommandApi struct {
	service *service.UserCommandAppService
}

func NewUserCommandApi() *UserCommandApi {
	return &UserCommandApi{
		service: service.NewUserCommandAppService(),
	}
}

func (c *UserCommandApi) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/tenants/{tenantId}/users/aggregate/{id}", "FindAggregateById")
	b.Handle("POST", "/tenants/{tenantId}/users", "UserCreate")
	b.Handle("POST", "/tenants/{tenantId}/users:get", "UserCreateAndGet")
	b.Handle("DELETE", "/tenants/{tenantId}/users", "UserDelete")
	b.Handle("PATCH", "/tenants/{tenantId}/users", "UserUpdate")
	b.Handle("PATCH", "/tenants/{tenantId}/users:get", "UserUpdateAndGet")
}

// FindAggregateById godoc
// @Summary      按聚合根ID查找聚合对象
// @Description  按聚合根ID查找聚合对象
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users/aggregate/{id} [get]
func (c *UserCommandApi) FindAggregateById(ictx iris.Context, tenantId string, id string) {
	_, _, _ = restapp.DoQueryOne(ictx, tenantId, func(ctx context.Context) (interface{}, bool, error) {
		return c.service.FindAggregateById(ctx, tenantId, id)
	})
}

// UserCreate godoc
// @Summary      创建用户
// @Description  创建用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users [POST]
func (c *UserCommandApi) UserCreate(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := userAssembler.AssUserCreateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.UserCreate(ctx, appCmd)
	})
}

// UserCreateAndGet godoc
// @Summary      创建用户
// @Description  创建用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users:get [POST]
func (c *UserCommandApi) UserCreateAndGet(ictx iris.Context, tenantId string) {
	_ = restapp.Do(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := userAssembler.AssUserCreateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, tenantId, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.UserCreate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
		})
		return err
	})
}

// UserDelete godoc
// @Summary      删除用户
// @Description  删除用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users [DELETE]
func (c *UserCommandApi) UserDelete(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := userAssembler.AssUserDeleteAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.UserDelete(ctx, appCmd)
	})
}

// UserUpdate godoc
// @Summary      更新用户
// @Description  更新用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users [PATCH]
func (c *UserCommandApi) UserUpdate(ictx iris.Context, tenantId string) {
	_ = restapp.DoCmd(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := userAssembler.AssUserUpdateAppCmd(ictx)
		if err != nil {
			return err
		}
		return c.service.UserUpdate(ctx, appCmd)
	})
}

// UserUpdateAndGet godoc
// @Summary      更新用户
// @Description  更新用户
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        string         true    "Tenant ID"
// @Param        id         path        string         true    "Aggregate ID"
// @Success      200        {object}    any
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users:get [PATCH]
func (c *UserCommandApi) UserUpdateAndGet(ictx iris.Context, tenantId string) {
	_ = restapp.Do(ictx, tenantId, func(ctx context.Context) error {
		appCmd, err := userAssembler.AssUserUpdateAppCmd(ictx)
		if err != nil {
			return err
		}

		_, _, err = restapp.DoCmdAndQueryOne(ictx, tenantId, c.service.QueryAppId, appCmd, func(ctx context.Context) error {
			return c.service.UserUpdate(ctx, appCmd)
		}, func(ctx context.Context) (interface{}, bool, error) {
			return c.service.QueryById(ctx, appCmd.GetTenantId(), appCmd.Data.Id)
		})
		return err
	})
}
