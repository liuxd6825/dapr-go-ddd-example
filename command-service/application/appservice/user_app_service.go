package appservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/model"
	domain_service "github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/service"
)

type UserAppService struct {
	userDomainService *domain_service.UserDomainService
}

func NewUserAppService() *UserAppService {
	return &UserAppService{
		userDomainService: &domain_service.UserDomainService{},
	}
}

func (s *UserAppService) CreateUser(ctx context.Context, cmd *user_commands.UserCreateCommand) error {
	a, err := s.userDomainService.CreateUser(ctx, cmd)
	print(a)
	return err
}

func (s *UserAppService) UpdateUser(ctx context.Context, cmd *user_commands.UserUpdateCommand) error {
	if err := cmd.Validate(); err != nil {
		return err
	}
	if cmd.GetIsValidOnly() {
		return nil
	}
	_, err := s.userDomainService.UpdateUser(ctx, cmd)
	return err
}

func (s *UserAppService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	return s.userDomainService.GetAggregateById(ctx, tenantId, id)
}

func (s *UserAppService) Get(ctx context.Context, cmd *user_commands.UserCreateCommand) error {
	a, err := s.userDomainService.CreateUser(ctx, cmd)
	print(a)
	return err
}
