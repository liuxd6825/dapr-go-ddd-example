package user_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/session"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/user/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// userDeleteAllExecutor
// @Description: 删除所有
//
type userDeleteAllExecutor struct {
	userService service.UserQueryDomainService
}

type userDeleteAllExecutorValidate struct {
	TenantId string
	Id       string
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *userDeleteAllExecutor) Execute(ctx context.Context, tenantId string) error {
	data := &userDeleteAllExecutorValidate{TenantId: tenantId}
	if err := e.Validate(data); err != nil {
		return err
	}

	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.userService.DeleteAll(ctx, tenantId)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userDeleteAllExecutor) Validate(v *userDeleteAllExecutorValidate) error {
	if v == nil {
		return errors.New("Validate() error: view is nil")
	}
	return nil
}

//
// newUserDeleteExecutor
// @Description: 新建命令执行器
// @return service.UserDeleteExecutor
//
func newUserDeleteAllExecutor() *userDeleteAllExecutor {
	return &userDeleteAllExecutor{
		userService: service_impl.GetUserQueryDomainService(),
	}
}
