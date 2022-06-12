package assembler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/dto"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
)

//
// AssUserCreateCommand
// @Description:
// @param ctx
// @param appCmd
// @return *command.UserCreateCommand
// @return error
//
func AssUserCreateCommand(ctx context.Context, appCmd *dto.UserCreateCommandDto) (*command.UserCreateCommand, error) {
	appcmd := *appCmd
	cmd := appcmd.UserCreateCommand
	return &cmd, nil
}

func AssUserUpdateCommand(ctx context.Context, appCmd *dto.UserUpdateCommandDto) (*command.UserUpdateCommand, error) {
	appcmd := *appCmd
	cmd := appcmd.UserUpdateCommand
	return &cmd, nil
}

func AssUserDeleteCommand(ctx context.Context, appCmd *dto.UserDeleteCommandDto) (*command.UserDeleteCommand, error) {
	appcmd := *appCmd
	cmd := appcmd.UserDeleteCommand
	return &cmd, nil
}
