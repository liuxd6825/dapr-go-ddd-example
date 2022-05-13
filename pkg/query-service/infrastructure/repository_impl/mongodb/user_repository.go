package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserRepository struct {
	base *BaseRepository[*projection.UserView]
}

func NewUserRepository(opts ...*RepositoryOptions) repository.UserViewRepository {
	newFunc := func() *projection.UserView {
		return &projection.UserView{}
	}
	return &UserRepository{
		base: NewBaseRepository[*projection.UserView](newFunc, "users", opts...),
	}
}

func (u *UserRepository) CreateById(ctx context.Context, user *projection.UserView) (*projection.UserView, error) {
	return u.base.CreateById(ctx, user)
}

func (u UserRepository) UpdateById(ctx context.Context, user *projection.UserView) (*projection.UserView, error) {
	return u.base.UpdateById(ctx, user)
}

func (u UserRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.base.DeleteById(ctx, tenantId, id)
}

func (u UserRepository) FindById(ctx context.Context, tenantId string, id string) (*projection.UserView, bool, error) {
	return u.base.FindById(ctx, tenantId, id)
}

func (u UserRepository) FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*projection.UserView] {
	return u.base.FindPaging(ctx, query)
}
