package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/repository"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/base/domain/repository/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserRepository struct {
	base *mongodb.BaseRepository[*view.UserView]
}

func NewUserRepository(opts ...*mongodb.RepositoryOptions) repository.UserViewRepository {
	newFunc := func() *view.UserView {
		return &view.UserView{}
	}
	return &UserRepository{
		base: mongodb.NewBaseRepository[*view.UserView](newFunc, "users", opts...),
	}
}

func (u *UserRepository) CreateById(ctx context.Context, user *view.UserView) (*view.UserView, error) {
	return u.base.CreateById(ctx, user)
}

func (u UserRepository) UpdateById(ctx context.Context, user *view.UserView) (*view.UserView, error) {
	return u.base.UpdateById(ctx, user)
}

func (u UserRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.base.DeleteById(ctx, tenantId, id)
}

func (u UserRepository) FindById(ctx context.Context, tenantId string, id string) (*view.UserView, bool, error) {
	return u.base.FindById(ctx, tenantId, id)
}

func (u UserRepository) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*view.UserView] {
	return u.base.FindPaging(ctx, query)
}

func (u *UserRepository) FindAll(ctx context.Context, tenantId string) (*[]*view.UserView, bool, error) {
	return u.base.FindAll(ctx, tenantId).Result()
}
