package repository

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserViewRepository interface {
	CreateById(ctx context.Context, user *view.UserView) (*view.UserView, error)
	UpdateById(ctx context.Context, user *view.UserView) (*view.UserView, error)
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindById(ctx context.Context, tenantId string, id string) (*view.UserView, bool, error)
	FindAll(ctx context.Context, tenantId string) (*[]*view.UserView, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[*view.UserView]
}
