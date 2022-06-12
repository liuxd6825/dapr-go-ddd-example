package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/assembler"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals/user/dto"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/model"
	domain_service "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/service"
	base "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/infrastructure/base/application/service"
	query_dto "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/user/dto"
)

//
// UserCommandAppService
// @Description:
//
type UserCommandAppService struct {
	base.BaseQueryAppService
	userDomainService *domain_service.UserDomainService
}

//
// NewCommandUserAppService
// @Description:
// @return *UserCommandAppService
//
func NewCommandUserAppService() *UserCommandAppService {
	res := &UserCommandAppService{
		userDomainService: &domain_service.UserDomainService{},
	}
	res.Init("query-service", "users", "v1.0")
	return res
}

//
// CreateUser
// @Description:
// @receiver s
// @param ctx
// @param cmd
// @return error
//
func (s *UserCommandAppService) CreateUser(ctx context.Context, cmdDto *dto.UserCreateCommandDto) error {
	cmd, err := assembler.AssUserCreateCommand(ctx, cmdDto)
	if err != nil {
		return err
	}
	_, err = s.userDomainService.UserCreate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) UpdateUser(ctx context.Context, cmdDto *dto.UserUpdateCommandDto) error {
	cmd, err := assembler.AssUserUpdateCommand(ctx, cmdDto)
	if err != nil {
		return err
	}
	_, err = s.userDomainService.UserUpdate(ctx, cmd)
	return err
}

func (s *UserCommandAppService) DeleteUser(ctx context.Context, cmdDto *dto.UserDeleteCommandDto) error {
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

func (s *UserCommandAppService) QueryById(ctx context.Context, tenantId string, id string) (*query_dto.UserFindByIdResponse, bool, error) {
	var resp query_dto.UserFindByIdResponse
	isFound, err := s.BaseQueryAppService.QueryById(ctx, tenantId, id, resp)
	if err != nil {
		return nil, false, err
	}
	return &resp, isFound, nil
}
