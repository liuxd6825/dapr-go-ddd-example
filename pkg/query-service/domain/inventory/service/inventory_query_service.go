package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
)

type Options interface {
	repository.Options
}

type InventoryQueryDomainService interface {
	Create(ctx context.Context, view *view.InventoryView, opts ...Options) error
	CreateMany(ctx context.Context, views []*view.InventoryView, opts ...Options) error

	Update(ctx context.Context, view *view.InventoryView, opts ...Options) error
	UpdateMany(ctx context.Context, views []*view.InventoryView, opts ...Options) error

	Delete(ctx context.Context, view *view.InventoryView, opts ...Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.InventoryView, opts ...Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...Options) error

	FindById(ctx context.Context, qry *query.InventoryFindByIdQuery, opts ...Options) (*view.InventoryView, bool, error)
	FindByIds(ctx context.Context, qry *query.InventoryFindByIdsQuery, opts ...Options) ([]*view.InventoryView, bool, error)
	FindAll(ctx context.Context, qry *query.InventoryFindAllQuery, opts ...Options) ([]*view.InventoryView, bool, error)
	FindPaging(ctx context.Context, qry *query.InventoryFindPagingQuery, opts ...Options) (*query.InventoryFindPagingResult, bool, error)
}
