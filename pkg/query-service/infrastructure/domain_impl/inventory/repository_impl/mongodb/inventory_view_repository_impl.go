package mongodb

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/dao/mongo_dao"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type InventoryViewRepositoryImpl struct {
	dao *mongo_dao.Dao[*view.InventoryView]
}

func NewInventoryViewRepository() repository.InventoryViewRepository {
	return &InventoryViewRepositoryImpl{
		dao: mongo_dao.NewDao[*view.InventoryView]("inventory"),
	}
}

func (r *InventoryViewRepositoryImpl) Create(ctx context.Context, view *view.InventoryView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Insert(ctx, view, ops...)
}

func (r *InventoryViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.InventoryView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.InsertMany(ctx, views, ops...)
}

func (r *InventoryViewRepositoryImpl) Update(ctx context.Context, view *view.InventoryView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Update(ctx, view, ops...)
}

func (r *InventoryViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.InventoryView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.UpdateMany(ctx, views, ops...)
}

func (r *InventoryViewRepositoryImpl) Delete(ctx context.Context, view *view.InventoryView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId(), ops...)
}

func (r *InventoryViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.InventoryView, opts ...service.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opts...)
}

func (r *InventoryViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, tenantId, id, ops...)
}

func (r *InventoryViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByIds(ctx, tenantId, ids, ops...)
}

func (r *InventoryViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteAll(ctx, tenantId, ops...)
}

func (r *InventoryViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByFilter(ctx, tenantId, filter, ops...)
}

func (r *InventoryViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opts ...service.Options) (*view.InventoryView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindById(ctx, tenantId, id, ops...)
}

func (r *InventoryViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) ([]*view.InventoryView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindByIds(ctx, tenantId, ids, ops...)
}

func (r *InventoryViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opts ...service.Options) ([]*view.InventoryView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindAll(ctx, tenantId, ops...).Result()
}

func (r *InventoryViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...service.Options) (*ddd_repository.FindPagingResult[*view.InventoryView], bool, error) {
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
