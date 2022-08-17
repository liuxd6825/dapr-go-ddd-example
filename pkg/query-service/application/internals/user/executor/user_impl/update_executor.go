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
// userUpdateExecutor
// @Description: 更新多个
//
type userUpdateExecutor struct {
	userService service.UserQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *userUpdateExecutor) Execute(ctx context.Context, view *view.UserView) error {
	if err := e.Validate(view); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.userService.Update(ctx, view)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userUpdateExecutor) Validate(view *view.UserView) error {
	if view == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newUserCreateExecutor
// @Description: 新建命令执行器
// @return service.UserDeleteExecutor
//
func newUserUpdateExecutor() *userUpdateExecutor {
	return &userUpdateExecutor{
		userService: service_impl.GetUserQueryDomainService(),
	}
}
