package command

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// InventoryCreateCommand
// @Description: 创建存货档案
//
type InventoryCreateCommand struct {
	CommandId   string                      `json:"commandId" validate:"required"` // 命令ID
	IsValidOnly bool                        `json:"isValidOnly" validate:"-"`      // 是否仅验证，不执行
	Data        field.InventoryCreateFields `json:"data" validate:"-"`             // 业务数据
}

//
// GetAggregateId
// @Description: 获取聚合根Id
//
func (c *InventoryCreateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.Id)
}

//
// GetCommandId
// @Description: 获取命令Id
//
func (c *InventoryCreateCommand) GetCommandId() string {
	return c.CommandId
}

//
// GetTenantId
// @Description: 获取租户Id
//
func (c *InventoryCreateCommand) GetTenantId() string {
	return c.Data.TenantId
}

//
// GetIsValidOnly
// @Description: 是否只验证不执行。
//
func (c *InventoryCreateCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

//
// Validate
// @Description: 命令数据验证
//
func (c *InventoryCreateCommand) Validate() error {
	ve := ddd.ValidateCommand(c, nil)
	/* 添加其他数据检查
	   if c.Data.Id == "" {
	       ve.AppendField("data.id", "不能为空")
	   }
	*/
	return ve.GetError()
}
