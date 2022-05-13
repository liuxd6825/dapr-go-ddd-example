package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type AddressRepository struct {
	base *BaseRepository[*projection.AddressView]
}

func NewAddressRepository(opts ...*RepositoryOptions) repository.AddressViewRepository {
	newFunc := func() *projection.AddressView {
		return &projection.AddressView{}
	}
	return &AddressRepository{
		base: NewBaseRepository[*projection.AddressView](newFunc, "address", opts...),
	}
}

func (u *AddressRepository) CreateById(ctx context.Context, addr *projection.AddressView) (res *projection.AddressView, resErr error) {
	return u.base.CreateById(ctx, addr)
}

func (u *AddressRepository) UpdateById(ctx context.Context, addr *projection.AddressView) (res *projection.AddressView, resErr error) {
	return u.base.UpdateById(ctx, addr)
}

func (u *AddressRepository) FindById(ctx context.Context, tenantId string, id string) (resUser *projection.AddressView, ok bool, resErr error) {
	return u.base.FindById(ctx, tenantId, id)
}

func (u *AddressRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.base.DeleteById(ctx, tenantId, id)
}

func (u *AddressRepository) FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*projection.AddressView] {
	return u.base.FindPaging(ctx, query)
}
