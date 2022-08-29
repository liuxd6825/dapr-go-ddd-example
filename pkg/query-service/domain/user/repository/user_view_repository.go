package repository

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserViewRepository interface {
	Create(ctx context.Context, view *view.UserView, opts ...service.Options) error
	CreateMany(ctx context.Context, views []*view.UserView, opts ...service.Options) error

	Update(ctx context.Context, view *view.UserView, opts ...service.Options) error
	UpdateMany(ctx context.Context, views []*view.UserView, opts ...service.Options) error

	Delete(ctx context.Context, view *view.UserView, opts ...service.Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.UserView, opts ...service.Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error

	FindById(ctx context.Context, tenantId string, id string, opts ...service.Options) (*view.UserView, bool, error)
	FindByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) ([]*view.UserView, bool, error)
	FindAll(ctx context.Context, tenantId string, opts ...service.Options) ([]*view.UserView, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...service.Options) (*ddd_repository.FindPagingResult[*view.UserView], bool, error)
}
