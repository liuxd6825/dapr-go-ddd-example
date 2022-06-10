package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_domain_service"
)

type UserDomainService struct {
	BaseDomainService
}

type UserValidate func(*model.UserAggregate) error

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
func (s *UserDomainService) UserCreate(ctx context.Context, cmd *command.UserCreateCommand, opts ...*ddd_domain_service.DoOptions) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// UserUpdate
// @Description:
// @receiver s
// @param ctx
// @param cmd
// @return *model.UserAggregate
// @return error
//
func (s *UserDomainService) UserUpdate(ctx context.Context, cmd *command.UserUpdateCommand, opts ...*ddd_domain_service.DoOptions) (*model.UserAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// UserDelete
// @Description:
// @receiver s
// @param ctx
// @param cmd
// @return *model.UserAggregate
// @return error
//
func (s *UserDomainService) UserDelete(ctx context.Context, cmd *command.UserDeleteCommand, opts ...*ddd_domain_service.DoOptions) (*model.UserAggregate, error) {
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
func (s *UserDomainService) doCommand(ctx context.Context, cmd ddd.Command, validateFunc func() error, opts ...*ddd_domain_service.DoOptions) (*model.UserAggregate, error) {
	options := ddd_domain_service.NewDoOptionsMerges(opts...)

	// 进行业务检查
	if validateFunc != nil {
		if err := validateFunc(); err != nil {
			return nil, err
		}
	} else if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// 如果只是业务检查，则不执行领域命令，
	validOnly := options.GetIsValidOnly()
	if (validOnly == nil && cmd.GetIsValidOnly()) || (validOnly != nil && *validOnly == true) {
		return nil, nil
	}

	// 执行领域命令
	var err error
	agg := s.NewAggregate()
	if _, ok := cmd.(*command.UserCreateCommand); ok {
		err = ddd.CreateAggregate(ctx, agg, cmd)
	} else {
		err = ddd.CommandAggregate(ctx, agg, cmd)
	}

	// 如果领域命令执行时出错，则返回错误
	if err != nil {
		return nil, err
	}

	return agg, nil
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
func (s *UserDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
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
func (s *UserDomainService) NewAggregate() *model.UserAggregate {
	return model.NewUserAggregate()
}
