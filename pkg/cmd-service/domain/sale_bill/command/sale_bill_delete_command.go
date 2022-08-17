package command

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// SaleBillDeleteCommand
// @Description: 删除销售订单
//
type SaleBillDeleteCommand struct {
	CommandId   string                     `json:"commandId" validate:"required"` // 命令ID
	IsValidOnly bool                       `json:"isValidOnly" validate:"-"`      // 是否仅验证，不执行
	Data        field.SaleBillDeleteFields `json:"data" validate:"-"`             // 业务数据
}

//
// GetAggregateId
// @Description: 获取聚合根Id
//
func (c *SaleBillDeleteCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.Id)
}

//
// GetCommandId
// @Description: 获取命令Id
//
func (c *SaleBillDeleteCommand) GetCommandId() string {
	return c.CommandId
}

//
// GetTenantId
// @Description: 获取租户Id
//
func (c *SaleBillDeleteCommand) GetTenantId() string {
	return c.Data.TenantId
}

//
// GetIsValidOnly
// @Description: 是否只验证不执行。
//
func (c *SaleBillDeleteCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

//
// Validate
// @Description: 命令数据验证
//
func (c *SaleBillDeleteCommand) Validate() error {
	ve := ddd.ValidateCommand(c, nil)
	/* 添加其他数据检查
	   if c.Data.Id == "" {
	       ve.AppendField("data.id", "不能为空")
	   }
	*/
	return ve.GetError()
}
