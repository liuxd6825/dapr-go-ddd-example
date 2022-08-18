package assembler

import (
	"context"
	"errors"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
)

//
// AssSaleBillCreateCommand
// @Description:  SaleBillCreateAppCmd 创建销售订单转换器
// @param ctx
// @param cmdDto
// @return *command.SaleBillCreateCommand
// @return error
//
func AssSaleBillCreateCommand(ctx context.Context, appCmd *appcmd.SaleBillCreateAppCmd) (*command.SaleBillCreateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssSaleBillCreateCommand() appCmd is nil")
	}
	cmd := (*appCmd).SaleBillCreateCommand
	return &cmd, nil
}

//
// AssSaleBillUpdateCommand
// @Description:  SaleBillUpdateAppCmd 更新销售订单转换器
// @param ctx
// @param cmdDto
// @return *command.SaleBillUpdateCommand
// @return error
//
func AssSaleBillUpdateCommand(ctx context.Context, appCmd *appcmd.SaleBillUpdateAppCmd) (*command.SaleBillUpdateCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssSaleBillUpdateCommand() appCmd is nil")
	}
	cmd := (*appCmd).SaleBillUpdateCommand
	return &cmd, nil
}

//
// AssSaleBillConfirmCommand
// @Description:  SaleBillConfirmAppCmd 下单确认命令转换器
// @param ctx
// @param cmdDto
// @return *command.SaleBillConfirmCommand
// @return error
//
func AssSaleBillConfirmCommand(ctx context.Context, appCmd *appcmd.SaleBillConfirmAppCmd) (*command.SaleBillConfirmCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssSaleBillConfirmCommand() appCmd is nil")
	}
	cmd := (*appCmd).SaleBillConfirmCommand
	return &cmd, nil
}

//
// AssSaleBillDeleteCommand
// @Description:  SaleBillDeleteAppCmd 删除销售订单转换器
// @param ctx
// @param cmdDto
// @return *command.SaleBillDeleteCommand
// @return error
//
func AssSaleBillDeleteCommand(ctx context.Context, appCmd *appcmd.SaleBillDeleteAppCmd) (*command.SaleBillDeleteCommand, error) {
	if appCmd == nil {
		return nil, errors.New("AssSaleBillDeleteCommand() appCmd is nil")
	}
	cmd := (*appCmd).SaleBillDeleteCommand
	return &cmd, nil
}
