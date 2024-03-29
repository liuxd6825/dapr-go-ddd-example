package assembler

import (
	"context"
	"errors"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
)

//
// AssSaleItemDeleteCommand
// @Description:  SaleItemDeleteAppCmd 删除销售明细项转换器
// @param ctx
// @param cmdDto
// @return *command.SaleItemDeleteCommand
// @return error
//
func AssSaleItemDeleteCommand(ctx context.Context, appCmd *appcmd.SaleItemDeleteAppCmd) (*command.SaleItemDeleteCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssSaleItemDeleteCommand() appCmd is nil")
	}
	cmd := (*appCmd).SaleItemDeleteCommand
	return &cmd, nil
}

//
// AssSaleItemCreateCommand
// @Description:  SaleItemCreateAppCmd 添加明细转换器
// @param ctx
// @param cmdDto
// @return *command.SaleItemCreateCommand
// @return error
//
func AssSaleItemCreateCommand(ctx context.Context, appCmd *appcmd.SaleItemCreateAppCmd) (*command.SaleItemCreateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssSaleItemCreateCommand() appCmd is nil")
	}
	cmd := (*appCmd).SaleItemCreateCommand
	return &cmd, nil
}

//
// AssSaleItemUpdateCommand
// @Description:  SaleItemUpdateAppCmd 更新明细转换器
// @param ctx
// @param cmdDto
// @return *command.SaleItemUpdateCommand
// @return error
//
func AssSaleItemUpdateCommand(ctx context.Context, appCmd *appcmd.SaleItemUpdateAppCmd) (*command.SaleItemUpdateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssSaleItemUpdateCommand() appCmd is nil")
	}
	cmd := (*appCmd).SaleItemUpdateCommand
	return &cmd, nil
}
