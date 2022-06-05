package queryservice_impl

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/repository"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/infrastructure/domain/repository_impl/mongodb"
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

func (u *UserQueryService) FindById(ctx context.Context, tenantId, userId string) (*projection.UserView, bool, error) {
	return u.repos.FindById(ctx, tenantId, userId)
}

func (u *UserQueryService) FindAll(ctx context.Context, tenantId string) (*[]*projection.UserView, bool, error) {
	return u.repos.FindAll(ctx, tenantId)
}

func (u *UserQueryService) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*projection.UserView], bool, error) {
	return u.repos.FindPaging(ctx, query).Result()
}
