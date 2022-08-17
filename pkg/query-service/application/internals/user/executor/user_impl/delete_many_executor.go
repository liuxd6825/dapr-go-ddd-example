package user_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/user/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// UserDeleteManyExecutor
// @Description: 删除多个
//
type userDeleteManyExecutor struct {
	userService service.UserQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *userDeleteManyExecutor) Execute(ctx context.Context, tenantId string, vList []*view.UserView) error {
	if err := e.Validate(vList); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.userService.DeleteMany(ctx, tenantId, vList)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userDeleteManyExecutor) Validate(view []*view.UserView) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newUserDeleteManyExecutor
// @Description: 新建命令执行器
// @return service.UserDeleteExecutor
//
func newUserDeleteManyExecutor() *userDeleteManyExecutor {
	return &userDeleteManyExecutor{
		userService: service_impl.GetUserQueryDomainService(),
	}
}
