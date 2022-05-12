package queryservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserQueryService interface {
	FindById(ctx context.Context, tenantId, userId string) (*projection.UserView, bool, error)
	Create(ctx context.Context, user *projection.UserView) error
	Update(ctx context.Context, user *projection.UserView) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindPagingData(ctx context.Context, query *ddd_repository.PagingQuery) (*ddd_repository.PagingData, bool, error)
}

func NewUserQueryService() UserQueryService {
	return queryservice_impl.NewUserQueryService()
}
