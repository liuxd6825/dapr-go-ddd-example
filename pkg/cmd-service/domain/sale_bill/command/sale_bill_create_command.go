package command

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// SaleBillCreateCommand
// @Description: 创建销售订单
//
type SaleBillCreateCommand struct {
	CommandId   string                     `json:"commandId" validate:"required"` // 命令ID
	IsValidOnly bool                       `json:"isValidOnly" validate:"-"`      // 是否仅验证，不执行
	Data        field.SaleBillCreateFields `json:"data" validate:"-"`             // 业务数据
}

//
// GetAggregateId
// @Description: 获取聚合根Id
//
func (c *SaleBillCreateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.Id)
}

//
// GetCommandId
// @Description: 获取命令Id
//
func (c *SaleBillCreateCommand) GetCommandId() string {
	return c.CommandId
}

//
// GetTenantId
// @Description: 获取租户Id
//
func (c *SaleBillCreateCommand) GetTenantId() string {
	return c.Data.TenantId
}

//
// GetIsValidOnly
// @Description: 是否只验证不执行。
//
func (c *SaleBillCreateCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

//
// IsAggregateCreateCommand
// @Description: 标识此命令为是聚合根创建命令，DDD框架层使用。
//
func (c *SaleBillCreateCommand) IsAggregateCreateCommand() {

}

//
// Validate
// @Description: 命令数据验证
//
func (c *SaleBillCreateCommand) Validate() error {
	ve := ddd.ValidateCommand(c, nil)
	/* 添加其他数据检查
	   if c.Data.Id == "" {
	       ve.AppendField("data.id", "不能为空")
	   }
	*/
	return ve.GetError()
}
