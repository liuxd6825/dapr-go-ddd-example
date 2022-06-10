package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/service/adto"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/userinterface/rest/user/dto"
)

type UserAssembler struct {
}

var User = &UserAssembler{}

func (a *UserAssembler) AssUserCreateAppCommand(ctx iris.Context) (*adto.UserCreateAppCommand, error) {
	var request dto.UserCreateRequest
	var cmd adto.UserCreateAppCommand
	if err := utils.AssemblerRequestBodyMapper(ctx, &request, &cmd, /*func() error {
			cmd.CommandId = request.CommandId
			cmd.IsValidOnly = request.IsValidOnly
			cmd.Data.Id = request.Data.Id
			cmd.Data.UserCode = request.Data.UserCode
			cmd.Data.Email = request.Data.Email
			cmd.Data.TenantId = request.Data.TenantId
			cmd.Data.Address = request.Data.Address
			cmd.Data.UserName = request.Data.UserName
			return nil
		}*/nil, func() error {
			return cmd.Validate()
		}); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func (a *UserAssembler) AssUserUpdateAppCommand(ctx iris.Context) (*adto.UserUpdateAppCommand, error) {
	var request dto.UserUpdateRequest
	var cmd adto.UserUpdateAppCommand
	if err := utils.AssemblerRequestBodyMapper(ctx, &request, &cmd, func() error {
		cmd.CommandId = request.CommandId
		cmd.IsValidOnly = request.IsValidOnly
		cmd.UpdateMask = request.UpdateMask
		cmd.Data.Id = request.Data.Id
		cmd.Data.UserCode = request.Data.UserCode
		cmd.Data.Email = request.Data.Email
		cmd.Data.TenantId = request.Data.TenantId
		cmd.Data.Address = request.Data.Address
		cmd.Data.UserName = request.Data.UserName
		return nil
	}, func() error {
		return cmd.Validate()
	}); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func (a *UserAssembler) AssUserDeleteAppCommand(ctx iris.Context) (*adto.UserDeleteAppCommand, error) {
	var request dto.UserDeleteByIdRequest
	var cmd adto.UserDeleteAppCommand
	if err := utils.AssemblerRequestBodyMapper(ctx, &request, &cmd, func() error {
		cmd.CommandId = request.CommandId
		cmd.IsValidOnly = request.IsValidOnly
		cmd.TenantId = request.Data.TenantId
		cmd.Id = request.Data.Id
		return nil
	}, func() error {
		return cmd.Validate()
	}); err != nil {
		return nil, err
	}
	return &cmd, nil
}
