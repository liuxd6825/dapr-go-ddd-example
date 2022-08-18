package dto

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/dto"
)

// 按ID查询

//
// SaleItemFindByIdResponse
// @Description:  按ID查询响应体
//
type SaleItemFindByIdResponse struct {
	SaleItemDto
}

func NewSaleItemFindByIdResponse() *SaleItemFindByIdResponse {
	return &SaleItemFindByIdResponse{}
}

// 按SaleItemId查询

//
// SaleItemFindBySaleBillIdRequest
// @Description: 销售明细项  查询SaleItemId请求体
//
type SaleItemFindBySaleBillIdRequest struct {
	TenantId   string `json:"tenantId" validate:"required"`
	SaleBillId string `json:"saleBillId" validate:"required"`
}

func NewSaleItemFindBySaleBillIdRequest() *SaleItemFindBySaleBillIdRequest {
	return &SaleItemFindBySaleBillIdRequest{}
}

//
// SaleItemFindBySaleBillIdResponse
// @Description: 销售明细项  查询SaleItemId响应体
//
type SaleItemFindBySaleBillIdResponse []*SaleItemFindBySaleBillIdResponseItem

func NewSaleItemFindBySaleBillIdResponse() *SaleItemFindBySaleBillIdResponse {
	return &SaleItemFindBySaleBillIdResponse{}
}

//
// SaleItemFindBySaleBillIdResponseItem
// @Description: 销售明细项  请求SaleItemId响应体
//
type SaleItemFindBySaleBillIdResponseItem struct {
	SaleItemDto
}

func NewSaleItemFindBySaleBillIdResponseItem() *SaleItemFindBySaleBillIdResponseItem {
	return &SaleItemFindBySaleBillIdResponseItem{}
}

// 分页查询

//
// SaleItemFindPagingResponse
// @Description: SaleItem 分页请求数据
//
type SaleItemFindPagingResponse struct {
	base.FindPagingResponse[*SaleItemFindPagingResponseItem]
}

func NewSaleItemFindPagingResponse() *SaleItemFindPagingResponse {
	resp := &SaleItemFindPagingResponse{}
	resp.Init()
	return resp
}

//
// SaleItemFindPagingResponseItem
// @Description: 销售明细项 请求业务数据
//
type SaleItemFindPagingResponseItem struct {
	SaleItemDto
}

func NewSaleItemFindPagingResponseItem() *SaleItemFindPagingResponseItem {
	return &SaleItemFindPagingResponseItem{}
}

// 查询所有

//
// SaleItemFindAllResponse
// @Description: 销售明细项 查询所有响应体
//
type SaleItemFindAllResponse []*SaleItemFindAllResponseItem

func NewSaleItemFindAllResponse() *SaleItemFindAllResponse {
	return &SaleItemFindAllResponse{}
}

//
// SaleItemFindAllResponseItem
// @Description: 销售明细项 请求业务数据
//
type SaleItemFindAllResponseItem struct {
	SaleItemDto
}

func NewSaleItemFindAllResponseItem() *SaleItemFindAllResponseItem {
	return &SaleItemFindAllResponseItem{}
}

//
// SaleItemDto
// @Description: 销售明细项 请求或响应业务数据
//
type SaleItemDto struct {
	base.BaseDto
	InventoryId   string  `json:"inventoryId,omitempty" validate:"required"` // 存货Id
	InventoryName string  `json:"inventoryName,omitempty" validate:"-"`      // 存货名称
	Money         float64 `json:"money,omitempty" validate:"-"`              // 文件大小
	Quantity      int64   `json:"quantity,omitempty" validate:"-"`           // 数量
	SaleBillId    string  `json:"saleBillId,omitempty" validate:"gt=0"`      //
}

func NewSaleItemDto() *SaleItemDto {
	return &SaleItemDto{}
}
