package repository

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
	dr "github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserViewRepository interface {
	CreateById(ctx context.Context, user *projection.UserView) (*projection.UserView, error)
	UpdateById(ctx context.Context, user *projection.UserView) (*projection.UserView, error)
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindById(ctx context.Context, tenantId string, id string) (*projection.UserView, bool, error)
	Search(ctx context.Context, searchQuery *dr.SearchQuery) (res *[]projection.UserView, ok bool, err error)
}
