package dto

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
)

type UserCreateCommandDto struct {
	command.UserCreateCommand
}
type UserDeleteCommandDto struct {
	command.UserDeleteCommand
}
type UserUpdateCommandDto struct {
	command.UserUpdateCommand
}
