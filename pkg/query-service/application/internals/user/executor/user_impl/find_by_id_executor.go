package user_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/user/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// userFindByIdExecutor
// @Description: 按Id查询
//
type userFindByIdExecutor struct {
	domainService service.UserQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *userFindByIdExecutor) Execute(ctx context.Context, aq *appquery.UserFindByIdAppQuery) (*view.UserView, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}

	return e.domainService.FindById(ctx, query.NewUserFindByIdQuery(aq.TenantId, aq.Id))
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userFindByIdExecutor) Validate(aq *appquery.UserFindByIdAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// NewUserCreateCommandExecutor
// @Description: 新建命令执行器
// @return service.UserDeleteExecutor
//
func newUserFindByIdExecutor() *userFindByIdExecutor {
	return &userFindByIdExecutor{
		domainService: service_impl.GetUserQueryDomainService(),
	}
}
