package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/model"
	domain_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/service"
)

type FindAggregateByIdExecutor interface {
	Execute(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error)
}

type findAggregateByIdExecutor struct {
	domainService *domain_service.UserCommandDomainService
}

//
// Execute
// @Description:
// @receiver s
// @param ctx 上下文
// @param tenantId 租户Id
// @param id 聚合根Id
// @return error
//
func (s *findAggregateByIdExecutor) Execute(ctx context.Context, tenantId string, id string) (*model.UserAggregate, bool, error) {
	return s.domainService.GetAggregateById(ctx, tenantId, id)
}

//
// newFindAggregateByIdExecutor
// @Description: 新建命令执行器
// @return service.FindAggregateByIdExecutor
//
func newFindAggregateByIdExecutor() *findAggregateByIdExecutor {
	return &findAggregateByIdExecutor{
		domainService: domain_service.GetUserCommandDomainService(),
	}
}
