package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/executor"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/model"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/base/application/service"
	query_dto "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/user/dto"
)

type UserCommandAppService struct {
	base.BaseQueryAppService
	userCreateCommandExecutor executor.UserCreateCommandExecutor
	userDeleteCommandExecutor executor.UserDeleteCommandExecutor
	userUpdateCommandExecutor executor.UserUpdateCommandExecutor
	findAggregateByIdExecutor executor.FindAggregateByIdExecutor
}

// NewUserCommandAppService
// @Description:  <no value>
// @return *UserCommandAppService
func NewUserCommandAppService() *UserCommandAppService {
	res := &UserCommandAppService{
		userCreateCommandExecutor: executor.GetUserCreateCommandExecutor(),
		userDeleteCommandExecutor: executor.GetUserDeleteCommandExecutor(),
		userUpdateCommandExecutor: executor.GetUserUpdateCommandExecutor(),
		findAggregateByIdExecutor: executor.GetFindAggregateByIdExecutor(),
	}
	res.Init("example-query-service", "users", "v1.0")
	return res
}

// UserCreate
// @Description: 创建用户
// @receiver s
// @param ctx 上下文
// @param cmd 创建用户命令DTO对象
// @return error
func (s *UserCommandAppService) UserCreate(ctx context.Context, appCmd *appcmd.UserCreateAppCmd) error {
	return s.userCreateCommandExecutor.Execute(ctx, appCmd)
}

// UserDelete
// @Description: 删除用户
// @receiver s
// @param ctx 上下文
// @param cmd 删除用户命令DTO对象
// @return error
func (s *UserCommandAppService) UserDelete(ctx context.Context, appCmd *appcmd.UserDeleteAppCmd) error {
	return s.userDeleteCommandExecutor.Execute(ctx, appCmd)
}

// UserUpdate
// @Description: 更新用户
// @receiver s
// @param ctx 上下文
// @param cmd 更新用户命令DTO对象
// @return error
func (s *UserCommandAppService) UserUpdate(ctx context.Context, appCmd *appcmd.UserUpdateAppCmd) error {
	return s.userUpdateCommandExecutor.Execute(ctx, appCmd)
}

// FindAggregateById
// @Description:
// @receiver s
// @param ctx 上下文
// @param tenantId 租户Id
// @param id 聚合根Id
// @return error
func (s *UserCommandAppService) FindAggregateById(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	return s.findAggregateByIdExecutor.Execute(ctx, tenantId, id)
}

// QueryById
// @Description: 按id获取<no value>投影类
// @receiver s queryAppService
// @param ctx 上下文
// @param tenantId  租户id
// @param id <no value> Id
// @return data <no value> 信息
// @return isFound 是否找到
// @return err 错误信息
func (s *UserCommandAppService) QueryById(ctx context.Context, tenantId string, id string) (*query_dto.UserFindByIdResponse, bool, error) {
	var resp query_dto.UserFindByIdResponse
	isFound, err := s.BaseQueryAppService.QueryById(ctx, tenantId, id, &resp)
	if err != nil {
		return nil, false, err
	}
	return &resp, isFound, nil
}

// QueryByIds
// @Description: 按ids获取<no value>投影类
// @receiver s queryAppService
// @param ctx 上下文
// @param tenantId  租户id
// @param ids 多个<no value>Id
// @return data <no value> 信息
// @return isFound 是否找到
// @return err 错误信息
func (s *UserCommandAppService) QueryByIds(ctx context.Context, tenantId string, ids []string) (*query_dto.UserFindByIdsResponse, bool, error) {
	var resp query_dto.UserFindByIdsResponse
	isFound, err := s.BaseQueryAppService.QueryByIds(ctx, tenantId, ids, &resp)
	if err != nil {
		return nil, false, err
	}
	return &resp, isFound, nil
}
