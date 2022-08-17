package dto

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/dto"
	"time"
)

// 按ID查询

//
// SaleBillFindByIdResponse
// @Description:  按ID查询响应体
//
type SaleBillFindByIdResponse struct {
	SaleBillDto
}

func NewSaleBillFindByIdResponse() *SaleBillFindByIdResponse {
	return &SaleBillFindByIdResponse{}
}

// 分页查询

//
// SaleBillFindPagingRequest
// @Description: 分页请求数据
//
type SaleBillFindPagingRequest struct {
	base.FindPagingRequest
}

func NewSaleBillFindPagingRequest() *SaleBillFindPagingRequest {
	return &SaleBillFindPagingRequest{}
}

//
// SaleBillFindPagingResponse
// @Description: SaleBill 分页响应数据
//
type SaleBillFindPagingResponse struct {
	base.FindPagingResponse[*SaleBillFindPagingResponseItem]
}

func NewSaleBillFindPagingResponse() *SaleBillFindPagingResponse {
	resp := &SaleBillFindPagingResponse{}
	resp.Init()
	return resp
}

//
// SaleBillFindPagingResponseItem
// @Description: 销售订单  请求业务数据
//
type SaleBillFindPagingResponseItem struct {
	SaleBillDto
}

func NewSaleBillFindPagingResponseItem() *SaleBillFindPagingResponseItem {
	return &SaleBillFindPagingResponseItem{}
}

// 查询所有

//
// SaleBillFindAllResponse
// @Description: 销售订单  查询所有响应体
//
type SaleBillFindAllResponse []*SaleBillFindAllResponseItem

func NewSaleBillFindAllResponse() *SaleBillFindAllResponse {
	return &SaleBillFindAllResponse{}
}

//
// SaleBillFindAllResponseItem
// @Description: 销售订单  请求业务数据
//
type SaleBillFindAllResponseItem struct {
	SaleBillDto
}

func NewSaleBillFindAllResponseItem() *SaleBillFindAllResponseItem {
	return &SaleBillFindAllResponseItem{}
}

//
// SaleBillDto
// @Description: 销售订单 请求或响应业务数据
//
type SaleBillDto struct {
	base.BaseDto
	SaleItems []*SaleItemDto `json:"saleItems,omitempty" validate:"-""`    //
	SaleMoney float64        `json:"saleMoney,omitempty" validate:"-"`     // 销售金额
	SaleTime  time.Time      `json:"saleTime,omitempty" validate:"-"`      // 文件大小
	UserId    string         `json:"userId,omitempty" validate:"required"` // 用户Id
	UserName  string         `json:"userName,omitempty" validate:"-"`      // 用户名称
}

func NewSaleBillDto() *SaleBillDto {
	return &SaleBillDto{}
}
