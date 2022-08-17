package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
)

type SaleItemQueryDomainService interface {
	Create(ctx context.Context, view *view.SaleItemView, opts ...Options) error
	CreateMany(ctx context.Context, views []*view.SaleItemView, opts ...Options) error

	Update(ctx context.Context, view *view.SaleItemView, opts ...Options) error
	UpdateMany(ctx context.Context, views []*view.SaleItemView, opts ...Options) error

	Delete(ctx context.Context, view *view.SaleItemView, opts ...Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.SaleItemView, opts ...Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...Options) error
	DeleteBySaleBillId(ctx context.Context, tenantId string, saleBillId string, opts ...Options) error

	FindById(ctx context.Context, qry *query.SaleItemFindByIdQuery, opts ...Options) (*view.SaleItemView, bool, error)
	FindByIds(ctx context.Context, qry *query.SaleItemFindByIdsQuery, opts ...Options) ([]*view.SaleItemView, bool, error)
	FindAll(ctx context.Context, qry *query.SaleItemFindAllQuery, opts ...Options) ([]*view.SaleItemView, bool, error)
	FindPaging(ctx context.Context, qry *query.SaleItemFindPagingQuery, opts ...Options) (*query.SaleItemFindPagingResult, bool, error)
	FindBySaleBillId(ctx context.Context, qry *query.SaleItemFindBySaleBillIdQuery, opts ...Options) ([]*view.SaleItemView, bool, error)
}
