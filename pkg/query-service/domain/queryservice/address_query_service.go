package queryservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type AddressQueryService struct {
	repos repository.AddressViewRepository
}

func NewAddressQueryService() *AddressQueryService {
	return &AddressQueryService{
		repos: mongodb.NewAddressRepository(),
	}
}

func (u *AddressQueryService) FindById(ctx context.Context, tenantId, id string) (*projection.AddressView, bool, error) {
	return u.repos.FindById(ctx, tenantId, id)
}

func (u *AddressQueryService) Create(ctx context.Context, user *projection.AddressView) error {
	_, err := u.repos.CreateById(ctx, user)
	return err
}

func (u *AddressQueryService) Update(ctx context.Context, user *projection.AddressView) error {
	_, err := u.repos.UpdateById(ctx, user)
	return err
}

func (u *AddressQueryService) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id)
}

func (u *AddressQueryService) FindPagingData(ctx context.Context, query *ddd_repository.PagingQuery) (*ddd_repository.PagingData, bool, error) {
	return u.repos.FindPaging(ctx, query)
}
