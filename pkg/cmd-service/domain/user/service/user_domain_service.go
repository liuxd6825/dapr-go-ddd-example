package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

// UserCommandDomainService
// @Description:  <no value> 命令领域服务
type UserCommandDomainService struct {
	ddd_service.CommandHelper[*model.UserAggregate]
}

// NewUserCommandDomainService
// @Description: 创建领域服务
// @return *UserCommandDomainService
func NewUserCommandDomainService() *UserCommandDomainService {
	return singleutils.Create[*UserCommandDomainService]("cmd-service.domain.user.UserCommandDomainService", func() *UserCommandDomainService {
		return &UserCommandDomainService{}
	})
}

// UserCreate
// @Description: 创建用户
// @receiver s
// @param ctx 上下文
// @param cmd 创建用户
// @return *model.UserCommandDomainService
// @return error
func (s *UserCommandDomainService) UserCreate(ctx context.Context, cmd *command.UserCreateCommand, opts ...ddd.DoCommandOption) (*model.UserAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// UserDelete
// @Description: 删除用户
// @receiver s
// @param ctx 上下文
// @param cmd 删除用户
// @return *model.UserCommandDomainService
// @return error
func (s *UserCommandDomainService) UserDelete(ctx context.Context, cmd *command.UserDeleteCommand, opts ...ddd.DoCommandOption) (*model.UserAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// UserUpdate
// @Description: 更新用户
// @receiver s
// @param ctx 上下文
// @param cmd 更新用户
// @return *model.UserCommandDomainService
// @return error
func (s *UserCommandDomainService) UserUpdate(ctx context.Context, cmd *command.UserUpdateCommand, opts ...ddd.DoCommandOption) (*model.UserAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}
