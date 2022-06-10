package assembler

import (
	"github.com/kataras/iris/v12"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/adto"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/userinterface/rest/user/dto"
)

type UserAssembler struct {
}

var User = &UserAssembler{}

func (a *UserAssembler) AssUserCreateCommandDto(ctx iris.Context) (*adto.UserCreateCommandDto, error) {
	var request dto.UserCreateRequest
	var cmd adto.UserCreateCommandDto
	if err := utils.AssemblerRequestBody(ctx, &request, &cmd); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func (a *UserAssembler) AssUserUpdateCommandDto(ctx iris.Context) (*adto.UserUpdateCommandDto, error) {
	var request dto.UserUpdateRequest
	var cmd adto.UserUpdateCommandDto
	if err := utils.AssemblerRequestBody(ctx, &request, &cmd); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	return &cmd, nil
}

func (a *UserAssembler) AssUserDeleteCommandDto(ctx iris.Context) (*adto.UserDeleteCommandDto, error) {
	var request dto.UserDeleteByIdRequest
	var cmd adto.UserDeleteCommandDto
	if err := utils.AssemblerRequestBody(ctx, &request, &cmd); err != nil {
		return nil, err
	}
	if err := cmd.Validate(); err != nil {
		return nil, err
	}
	return &cmd, nil
}
