package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/domain/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type AddressQueryDomainService interface {
	FindById(ctx context.Context, tenantId, id string) (*projection.AddressView, bool, error)
	Create(ctx context.Context, user *projection.AddressView) error
	Update(ctx context.Context, user *projection.AddressView) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindPagingData(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.AddressView], bool, error)
}

func NewAddressQueryService() AddressQueryDomainService {
	return queryservice_impl.NewAddressQueryService()
}
