package dto

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/dto"
)

// 按ID查询

//
// InventoryFindByIdResponse
// @Description:  按ID查询响应体
//
type InventoryFindByIdResponse struct {
	InventoryDto
}

func NewInventoryFindByIdResponse() *InventoryFindByIdResponse {
	return &InventoryFindByIdResponse{}
}

// 按多个ID查询

//
// InventoryFindByIdsResponse
// @Description: 存货档案  查询所有响应体
//
type InventoryFindByIdsResponse []*InventoryFindByIdsResponse

func NewInventoryFindByIdsResponse() *InventoryFindByIdsResponse {
	return &InventoryFindByIdsResponse{}
}

//
// InventoryFindByIdsResponseItem
// @Description: 存货档案  请求业务数据
//
type InventoryFindByIdsResponseItem struct {
	InventoryDto
}

func NewInventoryFindByIdsResponseItem() *InventoryFindByIdsResponseItem {
	return &InventoryFindByIdsResponseItem{}
}

// 分页查询

//
// InventoryFindPagingRequest
// @Description: 分页请求数据
//
type InventoryFindPagingRequest struct {
	base.FindPagingRequest
}

func NewInventoryFindPagingRequest() *InventoryFindPagingRequest {
	return &InventoryFindPagingRequest{}
}

//
// InventoryFindPagingResponse
// @Description: Inventory 分页响应数据
//
type InventoryFindPagingResponse struct {
	base.FindPagingResponse[*InventoryFindPagingResponseItem]
}

func NewInventoryFindPagingResponse() *InventoryFindPagingResponse {
	resp := &InventoryFindPagingResponse{}
	resp.Init()
	return resp
}

//
// InventoryFindPagingResponseItem
// @Description: 存货档案  请求业务数据
//
type InventoryFindPagingResponseItem struct {
	InventoryDto
}

func NewInventoryFindPagingResponseItem() *InventoryFindPagingResponseItem {
	return &InventoryFindPagingResponseItem{}
}

// 查询所有

//
// InventoryFindAllResponse
// @Description: 存货档案  查询所有响应体
//
type InventoryFindAllResponse []*InventoryFindAllResponseItem

func NewInventoryFindAllResponse() *InventoryFindAllResponse {
	return &InventoryFindAllResponse{}
}

//
// InventoryFindAllResponseItem
// @Description: 存货档案  请求业务数据
//
type InventoryFindAllResponseItem struct {
	InventoryDto
}

func NewInventoryFindAllResponseItem() *InventoryFindAllResponseItem {
	return &InventoryFindAllResponseItem{}
}

//
// InventoryDto
// @Description: 存货档案 请求或响应业务数据
//
type InventoryDto struct {
	base.BaseDto
	Brand    string `json:"brand,omitempty" validate:"-"`    // 品牌
	Keywords string `json:"keywords,omitempty" validate:"-"` // 搜索关键字
	Name     string `json:"name,omitempty" validate:"-"`     // 名称
	Spec     string `json:"spec,omitempty" validate:"-"`     // 规格
}

func NewInventoryDto() *InventoryDto {
	return &InventoryDto{}
}
