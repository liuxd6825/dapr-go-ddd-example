package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/domain/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserQueryDomainService interface {
	Create(ctx context.Context, user *projection.UserView) error
	Update(ctx context.Context, user *projection.UserView) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindById(ctx context.Context, tenantId, userId string) (*projection.UserView, bool, error)
	FindAll(ctx context.Context, tenantId string) (*[]*projection.UserView, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.UserView], bool, error)
}

func NewUserQueryService() UserQueryDomainService {
	return queryservice_impl.NewUserQueryService()
}
