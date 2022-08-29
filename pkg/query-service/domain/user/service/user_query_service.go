package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	"time"
)

type Options interface {
	GetTimeout() *time.Duration
	SetTimeout(v *time.Duration) Options
	GetUpdateMask() *[]string
	SetUpdateMask(*[]string) Options
	Merge(opts ...Options) Options
}

type UserQueryDomainService interface {
	Create(ctx context.Context, view *view.UserView, opts ...Options) error
	CreateMany(ctx context.Context, views []*view.UserView, opts ...Options) error

	Update(ctx context.Context, view *view.UserView, opts ...Options) error
	UpdateMany(ctx context.Context, views []*view.UserView, opts ...Options) error

	Delete(ctx context.Context, view *view.UserView, opts ...Options) error
	DeleteMany(ctx context.Context, tenantId string, views []*view.UserView, opts ...Options) error
	DeleteById(ctx context.Context, tenantId string, id string, opts ...Options) error
	DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...Options) error
	DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...Options) error
	DeleteAll(ctx context.Context, tenantId string, opts ...Options) error

	FindById(ctx context.Context, qry *query.UserFindByIdQuery, opts ...Options) (*view.UserView, bool, error)
	FindByIds(ctx context.Context, qry *query.UserFindByIdsQuery, opts ...Options) ([]*view.UserView, bool, error)
	FindAll(ctx context.Context, qry *query.UserFindAllQuery, opts ...Options) ([]*view.UserView, bool, error)
	FindPaging(ctx context.Context, qry *query.UserFindPagingQuery, opts ...Options) (*query.UserFindPagingResult, bool, error)
}
