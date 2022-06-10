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
	if err := utils.AssemblerRequestBody(ctx, &request, &cmd); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func (a *UserAssembler) AssUserUpdateAppCommand(ctx iris.Context) (*adto.UserUpdateAppCommand, error) {
	var request dto.UserUpdateRequest
	var cmd adto.UserUpdateAppCommand
	if err := utils.AssemblerRequestBody(ctx, &request, &cmd); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func (a *UserAssembler) AssUserDeleteAppCommand(ctx iris.Context) (*adto.UserDeleteAppCommand, error) {
	var request dto.UserDeleteByIdRequest
	var cmd adto.UserDeleteAppCommand
	if err := utils.AssemblerRequestBody(ctx, &request, &cmd); err != nil {
		return nil, err
	}
	cmd.Id = request.Data.Id
	cmd.CommandId = request.CommandId
	cmd.IsValidOnly = request.IsValidOnly
	cmd.TenantId = request.Data.TenantId
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	return &cmd, nil
}
