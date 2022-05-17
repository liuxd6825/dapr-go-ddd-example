package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/model/user_model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserDomainService struct {
	BaseDomainService
}

//
// NewUserDomainService
// @Description: 创建领域服务
// @return *UserDomainService
//
func NewUserDomainService() *UserDomainService {
	return &UserDomainService{}
}

//
// UserCreate
// @Description:
// @receiver s
// @param ctx
// @param cmd
// @return *model.UserAggregate
// @return error
//
func (s *UserDomainService) UserCreate(ctx context.Context, cmd *user_commands.UserCreateCommand) (*user_model.UserAggregate, error) {
	if err := s.ValidateCommand(cmd); err != nil {
		return nil, err
	}

	user := s.NewAggregate()
	err := ddd.CreateAggregate(ctx, user, cmd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) UserUpdate(ctx context.Context, cmd *user_commands.UserUpdateCommand) (*user_model.UserAggregate, error) {
	if err := s.ValidateCommand(cmd); err != nil {
		return nil, err
	}
	user := s.NewAggregate()
	if err := ddd.CommandAggregate(ctx, user, cmd); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) AddressCreate(ctx context.Context, cmd *user_commands.AddressCreateCommand) (*user_model.UserAggregate, error) {
	if err := s.ValidateCommand(cmd); err != nil {
		return nil, err
	}

	user := s.NewAggregate()
	err := ddd.CommandAggregate(ctx, user, cmd)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) AddressUpdate(ctx context.Context, cmd *user_commands.AddressUpdateCommand) (*user_model.UserAggregate, error) {
	if err := s.ValidateCommand(cmd); err != nil {
		return nil, err
	}
	user := s.NewAggregate()
	if err := ddd.CommandAggregate(ctx, user, cmd); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserDomainService) DeleteAddress(ctx context.Context, cmd *user_commands.AddressDeleteCommand) (*user_model.UserAggregate, error) {
	if err := s.ValidateCommand(cmd); err != nil {
		return nil, err
	}
	user := s.NewAggregate()
	if err := ddd.CommandAggregate(ctx, user, cmd); err != nil {
		return nil, err
	}
	return user, nil
}

//
// GetAggregateById
// @Description: 获取聚合根对象
// @receiver s
// @param ctx 上下文
// @param tenantId 租户id
// @param id 主键id
// @return *model.UserAggregate  聚合根对象
// @return bool 是否找到聚合根对象
// @return error 错误对象
//
func (s *UserDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*user_model.UserAggregate, bool, error) {
	user := s.NewAggregate()
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, user)
	return user, ok, err
}

//
// NewAggregate
// @Description: 新建聚合根对象
// @receiver s
// @return *model.UserAggregate
//
func (s *UserDomainService) NewAggregate() *user_model.UserAggregate {
	return user_model.NewUserAggregate()
}
