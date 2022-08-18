package dto

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
)

// SaleItemDeleteCommand

//
// SaleItemDeleteCommandRequest
// @Description: 删除销售明细项
//
type SaleItemDeleteCommandRequest struct {
	CommandId   string                           `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                             `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	Data        SaleItemDeleteCommandRequestData `json:"data"  validate:"required"`
}

//
// SaleItemDeleteCommandRequestData
// @Description: 删除销售明细项
//
type SaleItemDeleteCommandRequestData struct {
	Items      field.SaleItemDeleteItem `json:"items,omitempty" validate:"-"`             //
	SaleBillId string                   `json:"saleBillId,omitempty" validate:"required"` // 销售单Id
	TenantId   string                   `json:"tenantId,omitempty" validate:"required"`   // 租户ID
}

//
// SaleItemDeleteCommandResponse
// @Description: 删除销售明细项
type SaleItemDeleteCommandResponse struct {
}

// SaleItemCreateCommand

//
// SaleItemCreateCommandRequest
// @Description: 添加明细
//
type SaleItemCreateCommandRequest struct {
	CommandId   string                           `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                             `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	Data        SaleItemCreateCommandRequestData `json:"data"  validate:"required"`
}

//
// SaleItemCreateCommandRequestData
// @Description: 添加明细
//
type SaleItemCreateCommandRequestData struct {
	Items      field.SaleItemCreateItem `json:"items,omitempty" validate:"-"`             //
	SaleBillId string                   `json:"saleBillId,omitempty" validate:"required"` //
	TenantId   string                   `json:"tenantId,omitempty" validate:"required"`   // 租户ID
}

//
// SaleItemCreateCommandResponse
// @Description: 添加明细
type SaleItemCreateCommandResponse struct {
}

// SaleItemUpdateCommand

//
// SaleItemUpdateCommandRequest
// @Description: 更新明细
//
type SaleItemUpdateCommandRequest struct {
	CommandId   string                           `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                             `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	UpdateMask  []string                         `json:"updateMask"  validate:"-"`       // 要更新的字段项，空值：更新所有字段
	Data        SaleItemUpdateCommandRequestData `json:"data"  validate:"required"`
}

//
// SaleItemUpdateCommandRequestData
// @Description: 更新明细
//
type SaleItemUpdateCommandRequestData struct {
	Items      field.SaleItemUpdateItem `json:"items,omitempty" validate:"-"`             //
	SaleBillId string                   `json:"saleBillId,omitempty" validate:"required"` //
	TenantId   string                   `json:"tenantId,omitempty" validate:"required"`   // 租户ID
}

//
// SaleItemUpdateCommandResponse
// @Description: 更新明细
type SaleItemUpdateCommandResponse struct {
}

//
// SaleItemDto
// @Description: 销售明细项  请求或响应业务数据
//
type SaleItemDto struct {
	Id            string  `json:"id,omitempty" validate:"required"`          // 主键
	InventoryId   string  `json:"inventoryId,omitempty" validate:"required"` // 存货Id
	InventoryName string  `json:"inventoryName,omitempty" validate:"-"`      // 存货名称
	Money         float64 `json:"money,omitempty" validate:"-"`              // 文件大小
	Quantity      int64   `json:"quantity,omitempty" validate:"-"`           // 数量
	Remarks       string  `json:"remarks,omitempty" validate:"-"`            // 备注
	SaleBillId    string  `json:"saleBillId,omitempty" validate:"gt=0"`      //
	TenantId      string  `json:"tenantId,omitempty" validate:"required"`    // 租户ID
}
