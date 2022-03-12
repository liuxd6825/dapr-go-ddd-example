package queryservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/repository"
)

type UserQueryService struct {
	repos repository.UserViewRepository
}

func NewUserQueryService(repos repository.UserViewRepository) *UserQueryService {
	return &UserQueryService{
		repos: repos,
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
