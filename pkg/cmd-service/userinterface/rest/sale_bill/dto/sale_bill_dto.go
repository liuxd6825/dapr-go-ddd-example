package dto

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

// SaleBillCreateCommand

//
// SaleBillCreateCommandRequest
// @Description: 创建销售订单
//
type SaleBillCreateCommandRequest struct {
	CommandId   string                           `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                             `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	Data        SaleBillCreateCommandRequestData `json:"data"  validate:"required"`
}

//
// SaleBillCreateCommandRequestData
// @Description: 创建销售订单
//
type SaleBillCreateCommandRequestData struct {
	Id        string          `json:"id,omitempty" validate:"required"`       // 主键
	Remarks   string          `json:"remarks,omitempty" validate:"-"`         // 备注
	SaleMoney float64         `json:"saleMoney,omitempty" validate:"-"`       // 销售金额
	SaleTime  *types.JSONDate `json:"saleTime,omitempty" validate:"-"`        // 文件大小
	TenantId  string          `json:"tenantId,omitempty" validate:"required"` // 租户ID
	UserId    string          `json:"userId,omitempty" validate:"required"`   // 用户Id
	UserName  string          `json:"userName,omitempty" validate:"-"`        // 用户名称
}

//
// SaleBillCreateCommandResponse
// @Description: 创建销售订单
type SaleBillCreateCommandResponse struct {
}

// SaleBillUpdateCommand

//
// SaleBillUpdateCommandRequest
// @Description: 更新销售订单
//
type SaleBillUpdateCommandRequest struct {
	CommandId   string                           `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                             `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	UpdateMask  []string                         `json:"updateMask"  validate:"-"`       // 要更新的字段项，空值：更新所有字段
	Data        SaleBillUpdateCommandRequestData `json:"data"  validate:"required"`
}

//
// SaleBillUpdateCommandRequestData
// @Description: 更新销售订单
//
type SaleBillUpdateCommandRequestData struct {
	Id        string          `json:"id,omitempty" validate:"required"`       // 主键
	Remarks   string          `json:"remarks,omitempty" validate:"-"`         // 备注
	SaleMoney float64         `json:"saleMoney,omitempty" validate:"-"`       // 销售金额
	SaleTime  *types.JSONDate `json:"saleTime,omitempty" validate:"-"`        // 文件大小
	TenantId  string          `json:"tenantId,omitempty" validate:"required"` // 租户ID
	UserId    string          `json:"userId,omitempty" validate:"required"`   // 用户Id
	UserName  string          `json:"userName,omitempty" validate:"-"`        // 用户名称
}

//
// SaleBillUpdateCommandResponse
// @Description: 更新销售订单
type SaleBillUpdateCommandResponse struct {
}

// SaleBillConfirmCommand

//
// SaleBillConfirmCommandRequest
// @Description: 下单确认命令
//
type SaleBillConfirmCommandRequest struct {
	CommandId   string                            `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                              `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	UpdateMask  []string                          `json:"updateMask"  validate:"-"`       // 要更新的字段项，空值：更新所有字段
	Data        SaleBillConfirmCommandRequestData `json:"data"  validate:"required"`
}

//
// SaleBillConfirmCommandRequestData
// @Description: 下单确认命令
//
type SaleBillConfirmCommandRequestData struct {
	Id       string `json:"id,omitempty" validate:"required"`       // 主键
	Remarks  string `json:"remarks,omitempty" validate:"-"`         // 备注
	TenantId string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}

//
// SaleBillConfirmCommandResponse
// @Description: 下单确认命令
type SaleBillConfirmCommandResponse struct {
}

// SaleBillDeleteCommand

//
// SaleBillDeleteCommandRequest
// @Description: 删除销售订单
//
type SaleBillDeleteCommandRequest struct {
	CommandId   string                           `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                             `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	Data        SaleBillDeleteCommandRequestData `json:"data"  validate:"required"`
}

//
// SaleBillDeleteCommandRequestData
// @Description: 删除销售订单
//
type SaleBillDeleteCommandRequestData struct {
	Id       string `json:"id,omitempty" validate:"required"`       // 主键
	Remarks  string `json:"remarks,omitempty" validate:"-"`         // 备注
	TenantId string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}

//
// SaleBillDeleteCommandResponse
// @Description: 删除销售订单
type SaleBillDeleteCommandResponse struct {
}

//
// SaleBillDto
// @Description: 销售订单  请求或响应业务数据
//
type SaleBillDto struct {
	Id        string               `json:"id,omitempty" validate:"required"`       // 主键
	IsDeleted bool                 `json:"isDeleted,omitempty" validate:"-"`       // 已删除
	Remarks   string               `json:"remarks,omitempty" validate:"-"`         // 备注
	SaleItems []*SaleItemDto       `json:"saleItems,omitempty" validate:"-"`       //
	SaleMoney float64              `json:"saleMoney,omitempty" validate:"-"`       // 销售金额
	SaleTime  *types.JSONTime      `json:"saleTime,omitempty" validate:"-"`        // 文件大小
	Statue    field.SaleBillStatue `json:"statue,omitempty" validate:"-"`          // 单据状态
	TenantId  string               `json:"tenantId,omitempty" validate:"required"` // 租户ID
	UserId    string               `json:"userId,omitempty" validate:"required"`   // 用户Id
	UserName  string               `json:"userName,omitempty" validate:"-"`        // 用户名称
}
