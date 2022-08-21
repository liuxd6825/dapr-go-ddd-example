package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/model"
	base_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/base/domain/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"sync"
)

//
// UserCommandDomainService
// @Description:  <no value> 命令领域服务
//
type UserCommandDomainService struct {
	base_service.BaseCommandDomainService
}

// 单例应用服务
var userCommandDomainService *UserCommandDomainService

// 并发安全
var onceUser sync.Once

//
// GetUserCommandDomainService
// @Description: 获取单例领域服务
// @return service.UserQueryDomainService
//
func GetUserCommandDomainService() *UserCommandDomainService {
	onceUser.Do(func() {
		userCommandDomainService = newUserCommandDomainService()
	})
	return userCommandDomainService
}

//
// NewUserCommandDomainService
// @Description: 创建领域服务
// @return *UserCommandDomainService
//
func newUserCommandDomainService() *UserCommandDomainService {
	return &UserCommandDomainService{}
}

//
// UserCreate
// @Description: 创建用户
// @receiver s
// @param ctx 上下文
// @param cmd 创建用户
// @return *model.UserCommandDomainService
// @return error
//
func (s *UserCommandDomainService) UserCreate(ctx context.Context, cmd *command.UserCreateCommand, opts ...ddd.DoCommandOption) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// UserDelete
// @Description: 删除用户
// @receiver s
// @param ctx 上下文
// @param cmd 删除用户
// @return *model.UserCommandDomainService
// @return error
//
func (s *UserCommandDomainService) UserDelete(ctx context.Context, cmd *command.UserDeleteCommand, opts ...ddd.DoCommandOption) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// UserUpdate
// @Description: 更新用户
// @receiver s
// @param ctx 上下文
// @param cmd 更新用户
// @return *model.UserCommandDomainService
// @return error
//
func (s *UserCommandDomainService) UserUpdate(ctx context.Context, cmd *command.UserUpdateCommand, opts ...ddd.DoCommandOption) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
//  doCommand
//  @Description:
//  @receiver s
//  @param ctx
//  @param cmd
//  @return *model.UserAggregate
//  @return error
//
func (s *UserCommandDomainService) doCommand(ctx context.Context, cmd ddd.Command, validateFunc func() error, opts ...ddd.DoCommandOption) (*model.UserAggregate, error) {
	option := ddd.NewDoCommandOptionMerges(opts...)

	// 进行业务检查
	if validateFunc != nil {
		if err := validateFunc(); err != nil {
			return nil, err
		}
	} else if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// 如果只是业务检查，则不执行领域命令，
	validOnly := option.GetIsValidOnly()
	if (validOnly == nil && cmd.GetIsValidOnly()) || (validOnly != nil && *validOnly == true) {
		return nil, nil
	}

	// 新建聚合根对象
	agg := s.NewAggregate()

	// 如果领域命令执行时出错，则返回错误
	if err := ddd.ApplyCommand(ctx, agg, cmd); err != nil {
		return nil, err
	}

	return agg, nil
}

//
// GetAggregateById
// @Description: 获取聚合对象
// @receiver s
// @param ctx 上下文
// @param tenantId 租户id
// @param id 主键id
// @return *user_model.UserCommandDomainService  聚合对象
// @return bool 是否找到聚合根对象
// @return error 错误对象
//
func (s *UserCommandDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	agg := s.NewAggregate()
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, agg)
	return agg, ok, err
}

//
// NewAggregate
// @Description: 新建聚合对象
// @receiver s
// @return *user_model.UserCommandDomainService 聚合对象
//
func (s *UserCommandDomainService) NewAggregate() *model.UserAggregate {
	return model.NewUserAggregate()
}
