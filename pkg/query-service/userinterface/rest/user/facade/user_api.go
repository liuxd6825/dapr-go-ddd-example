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

type UserQueryApi struct {
	base.BaseApi
	queryService *service.UserQueryAppService
}

var userAssembler = assembler.User

func NewUserApi() *UserQueryApi {
	return &UserQueryApi{
		queryService: service.GetUserQueryAppService(),
	}
}

func (m *UserQueryApi) BeforeActivation(b mvc.BeforeActivation) {
	restapp.Handle(b, "GET", "/tenants/{tenantId}/users/{id}", "FindById")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/users:all", "FindAll")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/users", "FindPaging")
}

// FindById godoc
// @Summary      按ID获取用户
// @Description  get string by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path       int           true  "Tenant ID"
// @Param        id         path       int           true  "User ID"
// @Success      200        {object}   dto.UserFindByIdResponse
// @Failure      404        {object}   string        "按ID找到数据"
// @Failure      500        {object}   string        "应用错误"
// @Router       /tenants/{tenantId}/users/{id} [get]
func (m *UserQueryApi) FindById(ctx iris.Context) {
	_, _, _ = restapp.DoQueryOne(ctx, func(c context.Context) (interface{}, bool, error) {
		req, err := userAssembler.AssFindByIdRequest(ctx)
		if err != nil {
			return nil, false, err
		}
		v, b, e := m.queryService.FindById(c, req.TenantId, req.Id)
		return userAssembler.AssFindByIdResponse(ctx, v, b, e)
	})
}

// FindAll godoc
// @Summary      获取所有用户
// @Description  get string by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId  path      int     true    "Tenant ID"
// @Success      200       {object}  dto.UserFindAllResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/users:all [get]
func (m *UserQueryApi) FindAll(ctx iris.Context) {
	_, _, _ = restapp.DoQuery(ctx, func(c context.Context) (interface{}, bool, error) {
		req, err := userAssembler.AssFindAllRequest(ctx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := m.queryService.FindAll(c, req.TenantId)
		return userAssembler.AssFindAllResponse(ctx, fpr, b, e)
	})
}

// FindPaging godoc
// @Summary      分页查询
// @Description  分页查询
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        int         true    "Tenant ID"
// @Success      200        {object}    dto.UserFindPagingResponse
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users [get]
func (m *UserQueryApi) FindPaging(ctx iris.Context) {
	_, _, _ = restapp.DoQuery(ctx, func(c context.Context) (interface{}, bool, error) {
		req, err := userAssembler.AssFindPagingRequest(ctx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := m.queryService.FindPaging(c, req)
		return userAssembler.AssFindPagingResponse(ctx, fpr, b, e)
	})
}
