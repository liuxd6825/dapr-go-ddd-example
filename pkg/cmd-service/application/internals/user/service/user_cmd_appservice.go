package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/adto"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/assembler"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/model"
	domain_service "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/service"
)

//
// UserCommandAppService
// @Description:
//
type UserCommandAppService struct {
	userDomainService *domain_service.UserDomainService
}

//
// NewCommandUserAppService
// @Description:
// @return *UserCommandAppService
//
func NewCommandUserAppService() *UserCommandAppService {
	return &UserCommandAppService{
		userDomainService: &domain_service.UserDomainService{},
	}
}

//
// CreateUser
// @Description:
// @receiver s
// @param ctx
// @param cmd
// @return error
//
func (s *UserCommandAppService) CreateUser(ctx context.Context, cmdDto *adto.UserCreateCommandDto) error {
	cmd, err := assembler.AssUserCreateCommand(ctx, cmdDto)
	if err != nil {
		return err
	}
	_, err = s.userDomainService.UserCreate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) UpdateUser(ctx context.Context, cmdDto *adto.UserUpdateCommandDto) error {
	cmd, err := assembler.AssUserUpdateCommand(ctx, cmdDto)
	if err != nil {
		return err
	}
	_, err = s.userDomainService.UserUpdate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) DeleteUser(ctx context.Context, cmdDto *adto.UserDeleteCommandDto) error {
	cmd, err := assembler.AssUserDeleteCommand(ctx, cmdDto)
	if err != nil {
		return err
	}
	_, err = s.userDomainService.UserDelete(ctx, cmd)
	return err
}

func (s *UserCommandAppService) CreateAddress(ctx context.Context, cmd *command.AddressCreateCommand) error {
	_, err := s.userDomainService.AddressCreate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) UpdateAddress(ctx context.Context, cmd *command.AddressUpdateCommand) error {
	_, err := s.userDomainService.AddressUpdate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) DeleteAddress(ctx context.Context, cmd *command.AddressDeleteCommand) error {
	_, err := s.userDomainService.AddressDelete(ctx, cmd)
	return err
}

func (s *UserCommandAppService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	return s.userDomainService.GetAggregateById(ctx, tenantId, id)
}
