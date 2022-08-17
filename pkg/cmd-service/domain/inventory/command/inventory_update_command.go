package command

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// InventoryUpdateCommand
// @Description: 更新存货档案
//
type InventoryUpdateCommand struct {
	CommandId   string                      `json:"commandId" validate:"required"` // 命令ID
	IsValidOnly bool                        `json:"isValidOnly" validate:"-"`      // 是否仅验证，不执行
	UpdateMask  []string                    `json:"updateMask" validate:"-"`       // 要更新的字段项，空值：更新所有字段
	Data        field.InventoryUpdateFields `json:"data" validate:"-"`             // 业务数据
}

//
// GetAggregateId
// @Description: 获取聚合根Id
//
func (c *InventoryUpdateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.Id)
}

//
// GetCommandId
// @Description: 获取命令Id
//
func (c *InventoryUpdateCommand) GetCommandId() string {
	return c.CommandId
}

//
// GetTenantId
// @Description: 获取租户Id
//
func (c *InventoryUpdateCommand) GetTenantId() string {
	return c.Data.TenantId
}

//
// GetIsValidOnly
// @Description: 是否只验证不执行。
//
func (c *InventoryUpdateCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

//
// Validate
// @Description: 命令数据验证
//
func (c *InventoryUpdateCommand) Validate() error {
	ve := ddd.ValidateCommand(c, nil)
	/* 添加其他数据检查
	   if c.Data.Id == "" {
	       ve.AppendField("data.id", "不能为空")
	   }
	*/
	return ve.GetError()
}
