package adto

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
)

type UserCreateAppCommand struct {
	command.UserCreateCommand
}
type UserDeleteAppCommand struct {
	command.UserDeleteCommand
}
type UserUpdateAppCommand struct {
	command.UserUpdateCommand
}
