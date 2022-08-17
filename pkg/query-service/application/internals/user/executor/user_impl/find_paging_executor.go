package user_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/user/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/user/service_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

//
// userCreateCommandCommandExecutor
// @Description: 分页查询
//
type userFindPagingExecutor struct {
	domainService service.UserQueryDomainService
}

//
// Execute
// @Description: 执行命令
// @param ctx 上下文
// @param appCmd 命令
// @return error 错误
//
func (e *userFindPagingExecutor) Execute(ctx context.Context, aq *appquery.UserFindPagingAppQuery) (*appquery.UserFindPagingResult, bool, error) {
	if err := e.Validate(aq); err != nil {
		return nil, false, err
	}
	qry := query.NewUserFindPagingQuery(aq.TenantId, aq.Fields, aq.Filter, aq.Sort, aq.PageNum, aq.PageSize)
	fpr, ok, err := e.domainService.FindPaging(ctx, qry)
	if err != nil {
		return nil, false, err
	}
	res := assembler.AssUserFindPagingResult(fpr)
	return res, ok, nil
}

//
// Validate
// @Description: 命令验证
// @param appCmd 应用层命令
// @return error 错误
//
func (e *userFindPagingExecutor) Validate(aq *appquery.UserFindPagingAppQuery) error {
	if aq == nil {
		return errors.New("Validate(aq) error: aq is nil")
	}
	if len(aq.TenantId) == 0 {
		return errors.ErrorOf("Validate(aq) error: aq.TenantId is nil")
	}
	return nil
}

//
// newUserUserFindPagingExecutor
// @Description: 新建命令执行器
// @return *userFindPagingExecutor
//
func newUserFindPagingExecutor() *userFindPagingExecutor {
	return &userFindPagingExecutor{
		domainService: service_impl.GetUserQueryDomainService(),
	}
}
