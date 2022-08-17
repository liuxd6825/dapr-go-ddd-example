package appcmd

import (
	domain "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
)

//
// SaleBillConfirmAppCmd
// @Description: 应用服务层命令, 下单确认命令
//
type SaleBillConfirmAppCmd struct {
	domain.SaleBillConfirmCommand
}

//
// SaleBillDeleteAppCmd
// @Description: 应用服务层命令, 删除销售订单
//
type SaleBillDeleteAppCmd struct {
	domain.SaleBillDeleteCommand
}

//
// SaleBillCreateAppCmd
// @Description: 应用服务层命令, 创建销售订单
//
type SaleBillCreateAppCmd struct {
	domain.SaleBillCreateCommand
}

//
// SaleBillUpdateAppCmd
// @Description: 应用服务层命令, 更新销售订单
//
type SaleBillUpdateAppCmd struct {
	domain.SaleBillUpdateCommand
}
