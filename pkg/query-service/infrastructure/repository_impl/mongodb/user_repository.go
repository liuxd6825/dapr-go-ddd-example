package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

type UserRepository struct {
	repos *ddd_mongodb.Repository[*projection.UserView]
}

func NewUserRepository() repository.UserViewRepository {
	coll := GetMongoDB().GetCollection("users")
	builder := ddd_repository.NewEntityBuilder(func() *projection.UserView {
		return &projection.UserView{}
	}, func() *[]*projection.UserView {
		return &[]*projection.UserView{}
	})
	return &UserRepository{
		repos: ddd_mongodb.NewRepository[*projection.UserView](builder, GetMongoDB(), coll),
	}
}

func (u *UserRepository) CreateById(ctx context.Context, user *projection.UserView) (resUser *projection.UserView, resErr error) {
	return u.repos.Insert(ctx, user).Result()
}

func (u *UserRepository) UpdateById(ctx context.Context, user *projection.UserView) (resUser *projection.UserView, resErr error) {
	return u.repos.Update(ctx, user).Result()
}

func (u *UserRepository) FindById(ctx context.Context, tenantId string, id string) (resUser *projection.UserView, ok bool, resErr error) {
	return u.repos.FindById(ctx, tenantId, id).Result()
}

func (u *UserRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id).GetError()
}

func (u *UserRepository) FindPaging(ctx context.Context, query *ddd_repository.PagingQuery) *ddd_repository.FindPagingResult[*projection.UserView] {
	return u.repos.FindPaging(ctx, query)
}
