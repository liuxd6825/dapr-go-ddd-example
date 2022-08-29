package repository

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type SaleBillViewRepository interface {
	Create(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error
	CreateMany(ctx context.Context, views []*view.SaleBillView, opts ...service.Options) error

	Update(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error
	UpdateMany(ctx context.Context, views []*view.SaleBillView, opts ...service.Options) error

	Delete(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.SaleBillView, opts ...service.Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error

	FindById(ctx context.Context, tenantId string, id string, opts ...service.Options) (*view.SaleBillView, bool, error)
	FindByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) ([]*view.SaleBillView, bool, error)
	FindAll(ctx context.Context, tenantId string, opts ...service.Options) ([]*view.SaleBillView, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...service.Options) (*ddd_repository.FindPagingResult[*view.SaleBillView], bool, error)
}
