package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/model"
)

func (s *UserDomainService) AddressCreate(ctx context.Context, cmd *command.AddressCreateCommand) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	})
}

func (s *UserDomainService) AddressUpdate(ctx context.Context, cmd *command.AddressUpdateCommand) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	})
}

func (s *UserDomainService) AddressDelete(ctx context.Context, cmd *command.AddressDeleteCommand) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	})
}
