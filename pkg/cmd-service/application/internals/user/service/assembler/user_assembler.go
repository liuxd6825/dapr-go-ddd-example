package assembler

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/service/adto"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
)

func AssUserCreateCommand(ctx context.Context, appCmd *adto.UserCreateAppCommand) (*command.UserCreateCommand, error) {
	appcmd := *appCmd
	cmd := appcmd.UserCreateCommand
	return &cmd, nil
}

func AssUserUpdateCommand(ctx context.Context, appCmd *adto.UserUpdateAppCommand) (*command.UserUpdateCommand, error) {
	appcmd := *appCmd
	cmd := appcmd.UserUpdateCommand
	return &cmd, nil
}

func AssUserDeleteCommand(ctx context.Context, appCmd *adto.UserDeleteAppCommand) (*command.UserDeleteCommand, error) {
	appcmd := *appCmd
	cmd := appcmd.UserDeleteCommand
	return &cmd, nil
}
