package assembler

import (
	"context"
	"errors"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/command"
)

//
// AssInventoryCreateCommand
// @Description:  InventoryCreateAppCmd 创建存货档案转换器
// @param ctx
// @param cmdDto
// @return *command.InventoryCreateCommand
// @return error
//
func AssInventoryCreateCommand(ctx context.Context, appCmd *appcmd.InventoryCreateAppCmd) (*command.InventoryCreateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssInventoryCreateCommand() appCmd is nil")
	}
	cmd := (*appCmd).InventoryCreateCommand
	return &cmd, nil
}

//
// AssInventoryUpdateCommand
// @Description:  InventoryUpdateAppCmd 更新存货档案转换器
// @param ctx
// @param cmdDto
// @return *command.InventoryUpdateCommand
// @return error
//
func AssInventoryUpdateCommand(ctx context.Context, appCmd *appcmd.InventoryUpdateAppCmd) (*command.InventoryUpdateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssInventoryUpdateCommand() appCmd is nil")
	}
	cmd := (*appCmd).InventoryUpdateCommand
	return &cmd, nil
}
