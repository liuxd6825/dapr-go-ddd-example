package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

type AddressRepository struct {
	repos *ddd_mongodb.Repository[*projection.AddressView]
}

func NewAddressRepository() repository.AddressViewRepository {
	coll := GetMongoDB().GetCollection("address")
	builder := ddd_repository.NewEntityBuilder(func() *projection.AddressView {
		return &projection.AddressView{}
	}, func() *[]*projection.AddressView {
		return &[]*projection.AddressView{}
	})
	return &AddressRepository{
		repos: ddd_mongodb.NewRepository(builder, GetMongoDB(), coll),
	}
}

func (u *AddressRepository) CreateById(ctx context.Context, addr *projection.AddressView) (res *projection.AddressView, resErr error) {
	return u.repos.Insert(ctx, addr).Result()
}

func (u *AddressRepository) UpdateById(ctx context.Context, addr *projection.AddressView) (res *projection.AddressView, resErr error) {
	return u.repos.Update(ctx, addr).Result()
}

func (u *AddressRepository) FindById(ctx context.Context, tenantId string, id string) (resUser *projection.AddressView, ok bool, resErr error) {
	return u.repos.FindById(ctx, tenantId, id).Result()
}

func (u *AddressRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id).GetError()
}

func (u *AddressRepository) FindPaging(ctx context.Context, query *ddd_repository.PagingQuery) *ddd_repository.FindPagingResult[*projection.AddressView] {
	return u.repos.FindPaging(ctx, query)
}
