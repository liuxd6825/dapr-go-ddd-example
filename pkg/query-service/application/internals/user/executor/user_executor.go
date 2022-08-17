package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
)

//
// UserCreateExecutor
// @Description: 新建
//
type UserCreateExecutor interface {
	Execute(context.Context, *view.UserView) error
}

//
// UserCreateManyExecutor
// @Description: 新建多个
//
type UserCreateManyExecutor interface {
	Execute(context.Context, []*view.UserView) error
}

//
// UserDeleteByIdExecutor
// @Description: 按Id删除
//
type UserDeleteByIdExecutor interface {
	Execute(ctx context.Context, tenantId string, id string) error
}

//
// UserDeleteManyExecutor
// @Description: 删除多个
//
type UserDeleteManyExecutor interface {
	Execute(context.Context, string, []*view.UserView) error
}

//
// UserDeleteAllExecutor
// @Description: 删除所有
//
type UserDeleteAllExecutor interface {
	Execute(ctx context.Context, tenantId string) error
}

//
// UserUpdateExecutor
// @Description: 更新
//
type UserUpdateExecutor interface {
	Execute(context.Context, *view.UserView) error
}

//
// UserUpdateManyExecutor
// @Description: 更新多个
//
type UserUpdateManyExecutor interface {
	Execute(context.Context, []*view.UserView) error
}

//
// UserFindAllExecutor
// @Description: 查询所有
//
type UserFindAllExecutor interface {
	Execute(context.Context, *appquery.UserFindAllAppQuery) ([]*view.UserView, bool, error)
}

//
// UserFindByIdExecutor
// @Description: 按Id查询
//
type UserFindByIdExecutor interface {
	Execute(context.Context, *appquery.UserFindByIdAppQuery) (*view.UserView, bool, error)
}

//
// UserFindByIdsExecutor
// @Description: 按Id列表查询多个
//
type UserFindByIdsExecutor interface {
	Execute(context.Context, *appquery.UserFindByIdsAppQuery) ([]*view.UserView, bool, error)
}

//
// UserFindPagingExecutor
// @Description: 分页查询
//
type UserFindPagingExecutor interface {
	Execute(context.Context, *appquery.UserFindPagingAppQuery) (*appquery.UserFindPagingResult, bool, error)
}
