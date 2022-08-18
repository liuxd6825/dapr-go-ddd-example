package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/executor"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/executor/user_impl"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	"sync"
)

//
// UserQueryAppService
// @Description: <no value>查询应用服务类
//
type UserQueryAppService struct {
	userCreateExecutor     executor.UserCreateExecutor
	userCreateManyExecutor executor.UserCreateManyExecutor

	userUpdateExecutor     executor.UserUpdateExecutor
	userUpdateManyExecutor executor.UserUpdateManyExecutor

	userDeleteByIdExecutor executor.UserDeleteByIdExecutor
	userDeleteManyExecutor executor.UserDeleteManyExecutor
	userDeleteAllExecutor  executor.UserDeleteAllExecutor

	userFindAllExecutor    executor.UserFindAllExecutor
	userFindByIdExecutor   executor.UserFindByIdExecutor
	userFindByIdsExecutor  executor.UserFindByIdsExecutor
	userFindPagingExecutor executor.UserFindPagingExecutor
}

// 单例应用服务
var userQueryAppService *UserQueryAppService

// 并发安全
var onceUser sync.Once

//
// GetUserQueryAppService
// @Description: 获取单例应用服务
// @return *UserQueryAppService
//
func GetUserQueryAppService() *UserQueryAppService {
	onceUser.Do(func() {
		userQueryAppService = newUserQueryAppService()
	})
	return userQueryAppService
}

//
// NewUserQueryAppService
// @Description: 创建User查询应用服务
// @return *UserQueryAppService
//
func newUserQueryAppService() *UserQueryAppService {
	return &UserQueryAppService{
		userCreateExecutor:     user_impl.GetUserCreateExecutor(),
		userCreateManyExecutor: user_impl.GetUserCreateManyExecutor(),

		userUpdateExecutor:     user_impl.GetUserUpdateExecutor(),
		userUpdateManyExecutor: user_impl.GetUserUpdateManyExecutor(),

		userDeleteByIdExecutor: user_impl.GetUserDeleteByIdExecutor(),
		userDeleteManyExecutor: user_impl.GetUserDeleteManyExecutor(),
		userDeleteAllExecutor:  user_impl.GetUserDeleteAllExecutor(),

		userFindAllExecutor:    user_impl.GetUserFindAllExecutor(),
		userFindByIdExecutor:   user_impl.GetUserFindByIdExecutor(),
		userFindByIdsExecutor:  user_impl.GetUserFindByIdsExecutor(),
		userFindPagingExecutor: user_impl.GetUserFindPagingExecutor(),
	}
}

//
// Create
// @Description: 创建UserView
// @param ctx 上下文
// @param *view.UserView User实体对象
// @return error 错误
//
func (a *UserQueryAppService) Create(ctx context.Context, v *view.UserView) error {
	return a.userCreateExecutor.Execute(ctx, v)
}

//
// CreateMany
// @Description: 创建UserView
// @param ctx
// @return []*view.UserView  User实体对象切片
// @return error 错误
//
func (a *UserQueryAppService) CreateMany(ctx context.Context, vList []*view.UserView) error {
	return a.userCreateManyExecutor.Execute(ctx, vList)
}

//
// Update
// @Description: 按ID更新UserView
// @receiver a
// @param ctx
// @param v  *view.UserView
// @return error 错误
//
func (a *UserQueryAppService) Update(ctx context.Context, v *view.UserView) error {
	return a.userUpdateExecutor.Execute(ctx, v)
}

//
// UpdateMany
// @Description:  创建UserView
// @param ctx
// @return []*view.UserView  User实体对象切片
// @return error 错误
//
func (a *UserQueryAppService) UpdateMany(ctx context.Context, vList []*view.UserView) error {
	return a.userUpdateManyExecutor.Execute(ctx, vList)
}

//
// DeleteById
// @Description: 按ID删除UserView
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return error 错误
//
func (a *UserQueryAppService) DeleteById(ctx context.Context, tenantId, id string) error {
	return a.userDeleteByIdExecutor.Execute(ctx, tenantId, id)
}

//
// DeleteMany
// @Description: 删除多个UserView
// @param ctx
// @param tenantId 租户ID
// @param []*view.UserView  User实体对象切片
// @return error 错误
//
func (a *UserQueryAppService) DeleteMany(ctx context.Context, tenantId string, vList []*view.UserView) error {
	return a.userDeleteManyExecutor.Execute(ctx, tenantId, vList)
}

//
// DeleteAll
// @Description:  删除所有
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return error
//
func (a *UserQueryAppService) DeleteAll(ctx context.Context, tenantId string) error {
	return a.userDeleteAllExecutor.Execute(ctx, tenantId)

}

//
// FindById
// @Description:  按ID查询UserView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.UserView
// @return bool 是否查询到数据
// @return error
//
func (a *UserQueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.UserView, bool, error) {
	qry := assembler.User.AssFindByIdAppQuery(tenantId, id)
	return a.userFindByIdExecutor.Execute(ctx, qry)
}

//
// FindByIds
// @Description:  按多个ID查询UserView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.UserView
// @return bool 是否查询到数据
// @return error
//
func (a *UserQueryAppService) FindByIds(ctx context.Context, tenantId string, ids []string) ([]*view.UserView, bool, error) {
	qry := assembler.User.AssFindByIdsAppQuery(tenantId, ids)
	return a.userFindByIdsExecutor.Execute(ctx, qry)
}

//
// FindAll
// @Description: 查询所有view.UserView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return []*view.UserView
// @return bool 是否查询到数据
// @return error 错误
//
func (a *UserQueryAppService) FindAll(ctx context.Context, tenantId string) ([]*view.UserView, bool, error) {
	qry := assembler.User.AssFindAllAppQuery(tenantId)
	return a.userFindAllExecutor.Execute(ctx, qry)
}

//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param qry 分页查询条件
// @return *appquery.UserFindPagingResult 分页数据
// @return bool 是否查询到数据
// @return error 错误
//
func (a *UserQueryAppService) FindPaging(ctx context.Context, aq *appquery.UserFindPagingAppQuery) (*appquery.UserFindPagingResult, bool, error) {
	return a.userFindPagingExecutor.Execute(ctx, aq)
}
