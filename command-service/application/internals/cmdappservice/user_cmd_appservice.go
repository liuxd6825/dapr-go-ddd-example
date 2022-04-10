package cmdappservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/model"
	domain_service "github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/service"
)

type UserCommandAppService struct {
	userDomainService *domain_service.UserDomainService
}

func NewCommandUserAppService() *UserCommandAppService {
	return &UserCommandAppService{
		userDomainService: &domain_service.UserDomainService{},
	}
}

func (s *UserCommandAppService) CreateUser(ctx context.Context, cmd *user_commands.UserCreateCommand) error {
	_, err := s.userDomainService.UserCreate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) UpdateUser(ctx context.Context, cmd *user_commands.UserUpdateCommand) error {
	if err := cmd.Validate(); err != nil {
		return err
	}
	if cmd.GetIsValidOnly() {
		return nil
	}
	_, err := s.userDomainService.UserUpdate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) DeleteUser(ctx context.Context, cmd *user_commands.UserUpdateCommand) error {
	if err := cmd.Validate(); err != nil {
		return err
	}
	if cmd.GetIsValidOnly() {
		return nil
	}
	_, err := s.userDomainService.UserUpdate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) CreateAddress(ctx context.Context, cmd *user_commands.AddressCreateCommand) error {
	_, err := s.userDomainService.AddressCreate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) UpdateAddress(ctx context.Context, cmd *user_commands.AddressUpdateCommand) error {
	_, err := s.userDomainService.AddressUpdate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) DeleteAddress(ctx context.Context, cmd *user_commands.AddressDeleteCommand) error {
	_, err := s.userDomainService.DeleteAddress(ctx, cmd)
	return err
}

func (s *UserCommandAppService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	return s.userDomainService.GetAggregateById(ctx, tenantId, id)
}
