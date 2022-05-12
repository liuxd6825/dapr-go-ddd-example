package queryservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type AddressQueryService interface {
	FindById(ctx context.Context, tenantId, id string) (*projection.AddressView, bool, error)
	Create(ctx context.Context, user *projection.AddressView) error
	Update(ctx context.Context, user *projection.AddressView) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindPagingData(ctx context.Context, query *ddd_repository.PagingQuery) (*ddd_repository.PagingData, bool, error)
}

func NewAddressQueryService() AddressQueryService {
	return queryservice_impl.NewAddressQueryService()
}
