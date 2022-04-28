package repository

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type AddressViewRepository interface {
	CreateById(ctx context.Context, addr *projection.AddressView) (*projection.AddressView, error)
	UpdateById(ctx context.Context, addr *projection.AddressView) (*projection.AddressView, error)
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindById(ctx context.Context, tenantId string, id string) (*projection.AddressView, bool, error)
	FindPaging(ctx context.Context, query *ddd_repository.PagingQuery) (res *ddd_repository.PagingData, isFound bool, err error)
}
