package dto

import (
	base "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/userinterface/rest/dto"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/types"
)

//
// UserFindByIdRequest
// @Description:  请求内容
//
type UserFindByIdRequest struct {
	base.FindByIdRequest
}

//
// UserFindByIdResponse
// @Description:  请求内容
//
type UserFindByIdResponse struct {
	base.FindByIdResponse
	UserViewDto
}

//
// UserFindAllRequest
// @Description:
//
type UserFindAllRequest struct {
	base.FindAllRequest
}

//
// UserFindPagingRequest
// @Description:
//
type UserFindPagingRequest struct {
	base.FindPagingRequest[*UserViewDto]
}

//
// UserFindPagingResponse
// @Description:
//
type UserFindPagingResponse struct {
	base.FindPagingResponse
}

type UserViewList *[]*UserViewDto

//
// UserViewDto
// @Description:   请求业务数据
//
type UserViewDto struct {
	Id         string          `json:"id"`
	TenantId   string          `json:"tenantId"`
	UserCode   string          `json:"userCode"`
	UserName   string          `json:"userName"`
	Email      string          `json:"email"`
	Telephone  string          `json:"telephone"`
	Address    string          `json:"address"`
	CreateTime *types.DateTime `json:"createTime"`
}
