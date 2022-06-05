package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserQueryAppService struct {
	service service.UserQueryDomainService
}

//
// NewUserQueryAppService
// @Description: 创建应用服务
// @return *UserQueryAppService
//
func NewUserQueryAppService() *UserQueryAppService {
	return &UserQueryAppService{
		service: service.NewUserQueryService(),
	}
}

func (a *UserQueryAppService) FindById(ctx context.Context, tenantId string, userId string) (*projection.UserView, bool, error) {
	return a.service.FindById(ctx, tenantId, userId)
}

func (a *UserQueryAppService) FindAll(ctx context.Context, tenantId string) (*[]*projection.UserView, bool, error) {
	return a.service.FindAll(ctx, tenantId)
}

func (a *UserQueryAppService) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.UserView], bool, error) {
	return a.service.FindPaging(ctx, query)
}
