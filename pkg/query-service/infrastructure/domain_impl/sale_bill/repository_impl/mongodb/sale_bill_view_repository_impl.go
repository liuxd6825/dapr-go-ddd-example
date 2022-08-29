package mongodb

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/dao/mongo_dao"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type SaleBillViewRepositoryImpl struct {
	dao *mongo_dao.Dao[*view.SaleBillView]
}

func NewSaleBillViewRepository() repository.SaleBillViewRepository {
	return &SaleBillViewRepositoryImpl{
		dao: mongo_dao.NewDao[*view.SaleBillView]("sale_bill"),
	}
}

func (r *SaleBillViewRepositoryImpl) Create(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Insert(ctx, view, ops...)
}

func (r *SaleBillViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.SaleBillView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.InsertMany(ctx, views, ops...)
}

func (r *SaleBillViewRepositoryImpl) Update(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Update(ctx, view, ops...)
}

func (r *SaleBillViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.SaleBillView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.UpdateMany(ctx, views, ops...)
}

func (r *SaleBillViewRepositoryImpl) Delete(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId(), ops...)
}

func (r *SaleBillViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.SaleBillView, opts ...service.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opts...)
}

func (r *SaleBillViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, tenantId, id, ops...)
}

func (r *SaleBillViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByIds(ctx, tenantId, ids, ops...)
}

func (r *SaleBillViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteAll(ctx, tenantId, ops...)
}

func (r *SaleBillViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByFilter(ctx, tenantId, filter, ops...)
}

func (r *SaleBillViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opts ...service.Options) (*view.SaleBillView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindById(ctx, tenantId, id, ops...)
}

func (r *SaleBillViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) ([]*view.SaleBillView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindByIds(ctx, tenantId, ids, ops...)
}

func (r *SaleBillViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opts ...service.Options) ([]*view.SaleBillView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindAll(ctx, tenantId, ops...).Result()
}

func (r *SaleBillViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...service.Options) (*ddd_repository.FindPagingResult[*view.SaleBillView], bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindPaging(ctx, query, ops...).Result()
}

func newOptions(opts ...service.Options) []ddd_repository.Options {
	var options []ddd_repository.Options
	for _, o := range options {
		options = append(options, o)
	}
	return options
}
