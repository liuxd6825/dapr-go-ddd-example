package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"time"
)

type Options interface {
	GetTimeout() *time.Duration
	SetTimeout(v *time.Duration) Options
	GetUpdateMask() *[]string
	SetUpdateMask(*[]string) Options
	Merge(opts ...Options) Options
}

type SaleBillQueryDomainService interface {
	Create(ctx context.Context, view *view.SaleBillView, opts ...Options) error
	CreateMany(ctx context.Context, views []*view.SaleBillView, opts ...Options) error

	Update(ctx context.Context, view *view.SaleBillView, opts ...Options) error
	UpdateMany(ctx context.Context, views []*view.SaleBillView, opts ...Options) error

	Delete(ctx context.Context, view *view.SaleBillView, opts ...Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.SaleBillView, opts ...Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...Options) error

	FindById(ctx context.Context, qry *query.SaleBillFindByIdQuery, opts ...Options) (*view.SaleBillView, bool, error)
	FindByIds(ctx context.Context, qry *query.SaleBillFindByIdsQuery, opts ...Options) ([]*view.SaleBillView, bool, error)
	FindAll(ctx context.Context, qry *query.SaleBillFindAllQuery, opts ...Options) ([]*view.SaleBillView, bool, error)
	FindPaging(ctx context.Context, qry *query.SaleBillFindPagingQuery, opts ...Options) (*query.SaleBillFindPagingResult, bool, error)
}
