package queryservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserQueryService struct {
	repos repository.UserViewRepository
}

func NewUserQueryService() *UserQueryService {
	return &UserQueryService{
		repos: mongodb.NewUserRepository(),
	}
}

func (u *UserQueryService) FindById(ctx context.Context, tenantId, userId string) (*projection.UserView, bool, error) {
	return u.repos.FindById(ctx, tenantId, userId)
}

func (u *UserQueryService) Create(ctx context.Context, user *projection.UserView) error {
	_, err := u.repos.CreateById(ctx, user)
	return err
}

func (u *UserQueryService) Update(ctx context.Context, user *projection.UserView) error {
	_, err := u.repos.UpdateById(ctx, user)
	return err
}

func (u *UserQueryService) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id)
}

func (u *UserQueryService) FindPagingData(ctx context.Context, query *ddd_repository.PagingQuery) (*ddd_repository.PagingData, bool, error) {
	return u.repos.FindPaging(ctx, query)
}
