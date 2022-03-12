package repository

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
)

type UserViewRepository interface {
	CreateById(ctx context.Context, user *projection.UserView) (*projection.UserView, error)
	UpdateById(ctx context.Context, user *projection.UserView) (*projection.UserView, error)
	DeleteById(ctx context.Context, tenantId string, id string) error
}
