package appcmd

import (
	domain "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
)

//
// SaleItemUpdateAppCmd
// @Description: 应用服务层命令, 更新明细
//
type SaleItemUpdateAppCmd struct {
	domain.SaleItemUpdateCommand
}

//
// SaleItemDeleteAppCmd
// @Description: 应用服务层命令, 删除销售明细项
//
type SaleItemDeleteAppCmd struct {
	domain.SaleItemDeleteCommand
}

//
// SaleItemCreateAppCmd
// @Description: 应用服务层命令, 添加明细
//
type SaleItemCreateAppCmd struct {
	domain.SaleItemCreateCommand
}
