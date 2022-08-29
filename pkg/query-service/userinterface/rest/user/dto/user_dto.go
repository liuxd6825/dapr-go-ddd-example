package dto

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/dto"
)

// 按ID查询

//
// UserFindByIdResponse
// @Description:  按ID查询响应体
//
type UserFindByIdResponse struct {
	UserDto
}

func NewUserFindByIdResponse() *UserFindByIdResponse {
	return &UserFindByIdResponse{}
}

// 按多个ID查询

//
// UserFindByIdsResponse
// @Description: 用户  查询所有响应体
//
type UserFindByIdsResponse []*UserFindByIdsResponse

func NewUserFindByIdsResponse() *UserFindByIdsResponse {
	return &UserFindByIdsResponse{}
}

//
// UserFindByIdsResponseItem
// @Description: 用户  请求业务数据
//
type UserFindByIdsResponseItem struct {
	UserDto
}

func NewUserFindByIdsResponseItem() *UserFindByIdsResponseItem {
	return &UserFindByIdsResponseItem{}
}

// 分页查询

//
// UserFindPagingRequest
// @Description: 分页请求数据
//
type UserFindPagingRequest struct {
	base.FindPagingRequest
}

func NewUserFindPagingRequest() *UserFindPagingRequest {
	return &UserFindPagingRequest{}
}

//
// UserFindPagingResponse
// @Description: User 分页响应数据
//
type UserFindPagingResponse struct {
	base.FindPagingResponse[*UserFindPagingResponseItem]
}

func NewUserFindPagingResponse() *UserFindPagingResponse {
	resp := &UserFindPagingResponse{}
	resp.Init()
	return resp
}

//
// UserFindPagingResponseItem
// @Description: 用户  请求业务数据
//
type UserFindPagingResponseItem struct {
	UserDto
}

func NewUserFindPagingResponseItem() *UserFindPagingResponseItem {
	return &UserFindPagingResponseItem{}
}

// 查询所有

//
// UserFindAllResponse
// @Description: 用户  查询所有响应体
//
type UserFindAllResponse []*UserFindAllResponseItem

func NewUserFindAllResponse() *UserFindAllResponse {
	return &UserFindAllResponse{}
}

//
// UserFindAllResponseItem
// @Description: 用户  请求业务数据
//
type UserFindAllResponseItem struct {
	UserDto
}

func NewUserFindAllResponseItem() *UserFindAllResponseItem {
	return &UserFindAllResponseItem{}
}

//
// UserDto
// @Description: 用户 请求或响应业务数据
//
type UserDto struct {
	base.BaseDto
	Email string `json:"email,omitempty" validate:"-"` // 电子邮箱
	Name  string `json:"name,omitempty" validate:"-"`  // 用户名称
}

func NewUserDto() *UserDto {
	return &UserDto{}
}
