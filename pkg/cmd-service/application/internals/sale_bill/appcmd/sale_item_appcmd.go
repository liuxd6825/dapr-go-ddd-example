package appcmd

import (
	domain "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
)

//
// SaleItemCreateAppCmd
// @Description: 应用服务层命令, 创建扫描文件
//
type SaleItemCreateAppCmd struct {
	domain.SaleItemCreateCommand
}

//
// SaleItemUpdateAppCmd
// @Description: 应用服务层命令, 更新扫描文件
//
type SaleItemUpdateAppCmd struct {
	domain.SaleItemUpdateCommand
}

//
// SaleItemDeleteAppCmd
// @Description: 应用服务层命令, 删除扫描单
//
type SaleItemDeleteAppCmd struct {
	domain.SaleItemDeleteCommand
}
