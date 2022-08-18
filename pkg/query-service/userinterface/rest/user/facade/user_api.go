package facade

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/user/assembler"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

var UserAssembler = assembler.User

type UserQueryApi struct {
	queryService *service.UserQueryAppService
}

func NewUserQueryApi() *UserQueryApi {
	return &UserQueryApi{
		queryService: service.GetUserQueryAppService(),
	}
}

func (a *UserQueryApi) BeforeActivation(b mvc.BeforeActivation) {
	restapp.Handle(b, "GET", "/tenants/{tenantId}/users/{id}", "FindById")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/users:all", "FindAll")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/users:ids", "FindByIds")
	restapp.Handle(b, "GET", "/tenants/{tenantId}/users", "FindPaging")
}

// FindById godoc
// @Summary      按ID查询
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
func (a *UserQueryApi) FindById(ictx iris.Context) {
	_, _, _ = restapp.DoQueryOne(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := UserAssembler.AssFindByIdRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		v, b, e := a.queryService.FindById(ctx, req.TenantId, req.Id)
		return UserAssembler.AssFindByIdResponse(ictx, v, b, e)
	})
}

// FindByIds godoc
// @Summary      按多个ID获取<no value>
// @Description  get string by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId  path      string     true    "Tenant ID"
// @Param        id        path      string     true    "User ID"
// @Success      200       {object}  dto.UserFindByIdsResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/users:ids [get]
func (a *UserQueryApi) FindByIds(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := UserAssembler.AssFindByIdsRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindByIds(ctx, req.TenantId, req.Ids)
		return UserAssembler.AssFindByIdsResponse(ictx, fpr, b, e)
	})
}

// FindAll godoc
// @Summary      获取所有<no value>
// @Description  get string by ID
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId  path      int     true    "Tenant ID"
// @Success      200       {object}  dto.UserFindAllResponse
// @Failure      500       {object}  string          "应用错误"
// @Router       /tenants/{tenantId}/users:all [get]
func (a *UserQueryApi) FindAll(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := UserAssembler.AssFindAllRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindAll(ctx, req.TenantId)
		return UserAssembler.AssFindAllResponse(ictx, fpr, b, e)
	})
}

// FindPaging godoc
// @Summary      分页查询<no value>
// @Description  分页查询<no value>
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        tenantId   path        int         true    "Tenant ID"
// @Success      200        {object}    dto.UserFindPagingResponse
// @Failure      500        {object}    string      "应用错误"
// @Router       /tenants/{tenantId}/users [get]
func (a *UserQueryApi) FindPaging(ictx iris.Context, tenantId string) {
	_, _, _ = restapp.DoQuery(ictx, func(ctx context.Context) (interface{}, bool, error) {
		req, err := UserAssembler.AssFindPagingRequest(ictx)
		if err != nil {
			return nil, false, err
		}
		fpr, b, e := a.queryService.FindPaging(ctx, req)
		return UserAssembler.AssFindPagingResponse(ictx, fpr, b, e)
	})
}
