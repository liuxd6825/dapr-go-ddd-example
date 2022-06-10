package dto

import (
	base "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/userinterface/rest/dto"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

//
// UserFindByIdRequest
// @Description:  请求内容
//
type UserFindByIdRequest struct {
	TenantId string
	Id       string
}

//
// UserFindByIdResponse
// @Description:  请求内容
//
type UserFindByIdResponse struct {
	UserDto
}

//
// UserFindAllRequest
// @Description:
//
type UserFindAllRequest struct {
	TenantId string
}

type UserFindAllResponse []*UserFindAllResponseItem

type UserFindAllResponseItem struct {
	UserDto
}

//
// UserFindPagingRequest
// @Description:
//
type UserFindPagingRequest struct {
	base.FindPagingQuery
}

//
// UserFindPagingResponse
// @Description:
//
type UserFindPagingResponse struct {
	Data       *[]*UserFindPagingResponseItem `json:"data"`
	TotalRows  int64                          `json:"totalRows"`
	TotalPages int64                          `json:"totalPages"`
	PageNum    int64                          `json:"pageNum"`
	PageSize   int64                          `json:"pageSize"`
	Filter     string                         `json:"filter"`
	Sort       string                         `json:"sort"`
}

type UserFindPagingResponseItem struct {
	UpdateTime *types.JSONTime `json:"updateTime"`
	UserDto
}

//
// UserDto
// @Description:   请求业务数据
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
