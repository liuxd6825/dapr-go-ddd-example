package dto

import (
	base "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/userinterface/rest/dto"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

//
// UserFindByIdResponse
// @Description: 按Id查询的响应体
//
type UserFindByIdResponse struct {
	UserDto
}

//
// UserFindAllResponse
// @Description: 查询所有的响应体
//

type UserFindAllResponse []*UserFindAllResponseItem

type UserFindAllResponseItem struct {
	UserDto
}

//
// UserFindPagingResponse
// @Description: 分页查询的响应体
//
type UserFindPagingResponse struct {
	base.FindPagingResponse[*UserFindPagingResponseItem]
}

type UserFindPagingResponseItem struct {
	UserDto
}

//
// UserDto
// @Description:   响应业务数据
//
type UserDto struct {
	Id         string          `json:"id"`
	TenantId   string          `json:"tenantId"`
	UserCode   string          `json:"userCode"`
	UserName   string          `json:"userName"`
	Email      string          `json:"email"`
	Telephone  string          `json:"telephone"`
	Address    string          `json:"address"`
	CreateTime *types.JSONTime `json:"createTime"`
}
