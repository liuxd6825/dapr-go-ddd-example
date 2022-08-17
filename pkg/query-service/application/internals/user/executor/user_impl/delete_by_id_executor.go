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
// @Description: 按Id删除
//
type userDeleteByIdExecutor struct {
	userService service.UserQueryDomainService
}

type userDeleteByIdExecutorValidate struct {
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
func (e *userDeleteByIdExecutor) Execute(ctx context.Context, tenantId string, id string) error {
	data := &userDeleteByIdExecutorValidate{TenantId: tenantId, Id: id}
	if err := e.Validate(data); err != nil {
		return err
	}
	return session.StartSession(ctx, func(ctx context.Context) error {
		return e.userService.DeleteById(ctx, tenantId, id)
	})
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userDeleteByIdExecutor) Validate(v *userDeleteByIdExecutorValidate) error {
	if v == nil {
		return errors.New("Validate() error: view is nil")
	}
	if len(v.TenantId) == 0 {
		return errors.New("Validate() error: tenantId is nil")
	}
	if len(v.Id) == 0 {
		return errors.New("Validate() error: id is nil")
	}
	return nil
}

//
// newUserDeleteExecutor
// @Description: 新建命令执行器
// @return service.UserDeleteExecutor
//
func newUserDeleteByIdExecutor() *userDeleteByIdExecutor {
	return &userDeleteByIdExecutor{
		userService: service_impl.GetUserQueryDomainService(),
	}
}
