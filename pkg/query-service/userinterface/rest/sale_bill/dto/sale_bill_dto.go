package dto

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/userinterface/rest/dto"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
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

// 按多个ID查询

//
// SaleBillFindByIdsResponse
// @Description: 销售订单  查询所有响应体
//
type SaleBillFindByIdsResponse []*SaleBillFindByIdsResponse

func NewSaleBillFindByIdsResponse() *SaleBillFindByIdsResponse {
	return &SaleBillFindByIdsResponse{}
}

//
// SaleBillFindByIdsResponseItem
// @Description: 销售订单  请求业务数据
//
type SaleBillFindByIdsResponseItem struct {
	SaleBillDto
}

func NewSaleBillFindByIdsResponseItem() *SaleBillFindByIdsResponseItem {
	return &SaleBillFindByIdsResponseItem{}
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
	SaleItems []*SaleItemDto       `json:"saleItems,omitempty" validate:"-""`    //
	SaleMoney float64              `json:"saleMoney,omitempty" validate:"-"`     // 销售金额
	SaleTime  *types.JSONTime      `json:"saleTime,omitempty" validate:"-"`      // 文件大小
	Statue    field.SaleBillStatue `json:"statue,omitempty" validate:"-"`        // 单据状态
	UserId    string               `json:"userId,omitempty" validate:"required"` // 用户Id
	UserName  string               `json:"userName,omitempty" validate:"-"`      // 用户名称
}

func NewSaleBillDto() *SaleBillDto {
	return &SaleBillDto{}
}
