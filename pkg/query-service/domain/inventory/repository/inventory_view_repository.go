package repository

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type Options interface {
	Timeout() *int
	UpdateMask() *[]string
}

type InventoryViewRepository interface {
	Create(ctx context.Context, view *view.InventoryView, opt Options) error
	CreateMany(ctx context.Context, views []*view.InventoryView, opt Options) error

	Update(ctx context.Context, view *view.InventoryView, opt Options) error
	UpdateMany(ctx context.Context, views []*view.InventoryView, opt Options) error

	Delete(ctx context.Context, view *view.InventoryView, opt Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.InventoryView, opt Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opt Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opt Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opt Options) error
	DeleteAll(ctx context.Context, tenantId string, opt Options) error

	FindById(ctx context.Context, tenantId string, id string, opt Options) (*view.InventoryView, bool, error)
	FindByIds(ctx context.Context, tenantId string, ids []string, opt Options) ([]*view.InventoryView, bool, error)
	FindAll(ctx context.Context, tenantId string, opt Options) ([]*view.InventoryView, bool, error)
	FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opt Options) (*ddd_repository.FindPagingResult[*view.InventoryView], bool, error)
}
