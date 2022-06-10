package repository

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type AddressViewRepository interface {
	CreateById(ctx context.Context, addr *view.AddressView) (*view.AddressView, error)
	UpdateById(ctx context.Context, addr *view.AddressView) (*view.AddressView, error)
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindById(ctx context.Context, tenantId string, id string) (*view.AddressView, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*view.AddressView]
}
