package appcmd

import (
	domain "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/command"
)

//
// InventoryCreateAppCmd
// @Description: 应用服务层命令, 创建存货档案
//
type InventoryCreateAppCmd struct {
	domain.InventoryCreateCommand
}

//
// InventoryUpdateAppCmd
// @Description: 应用服务层命令, 更新存货档案
//
type InventoryUpdateAppCmd struct {
	domain.InventoryUpdateCommand
}
