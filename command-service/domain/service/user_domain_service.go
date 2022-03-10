package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserDomainService struct {
}

var validate = validator.New()

func (s *UserDomainService) CreateUser(ctx context.Context, cmd *user_commands.UserCreateCommand) (*model.UserAggregate, error) {
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}

	user := model.NewUserAggregate()
	err := ddd.CreateAggregate(ctx, user, cmd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) UpdateUser(ctx context.Context, cmd *user_commands.UserUpdateCommand) (*model.UserAggregate, error) {
	if err := validate.Struct(cmd); err != nil {
		return nil, err
	}
	user := model.NewUserAggregate()
	if err := ddd.CommandAggregate(ctx, user, cmd); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	agg := &model.UserAggregate{}
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, agg)
	return agg, ok, err
}
