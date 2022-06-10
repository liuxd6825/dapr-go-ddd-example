package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/repository"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/domain/repository/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type AddressRepository struct {
	base *mongodb.BaseRepository[*view.AddressView]
}

func NewAddressRepository(opts ...*mongodb.RepositoryOptions) repository.AddressViewRepository {
	newFunc := func() *view.AddressView {
		return &view.AddressView{}
	}
	return &AddressRepository{
		base: mongodb.NewBaseRepository[*view.AddressView](newFunc, "address", opts...),
	}
}

func (u *AddressRepository) CreateById(ctx context.Context, addr *view.AddressView) (res *view.AddressView, resErr error) {
	return u.base.CreateById(ctx, addr)
}

func (u *AddressRepository) UpdateById(ctx context.Context, addr *view.AddressView) (res *view.AddressView, resErr error) {
	return u.base.UpdateById(ctx, addr)
}

func (u *AddressRepository) FindById(ctx context.Context, tenantId string, id string) (resUser *view.AddressView, ok bool, resErr error) {
	return u.base.FindById(ctx, tenantId, id)
}

func (u *AddressRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.base.DeleteById(ctx, tenantId, id)
}

func (u *AddressRepository) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*view.AddressView] {
	return u.base.FindPaging(ctx, query)
}
