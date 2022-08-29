package mongodb

import (
	"context"
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/dao/mongo_dao"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type SaleItemViewRepositoryImpl struct {
	dao *mongo_dao.Dao[*view.SaleItemView]
}

func NewSaleItemViewRepository() repository.SaleItemViewRepository {
	return &SaleItemViewRepositoryImpl{
		dao: mongo_dao.NewDao[*view.SaleItemView]("sale_item"),
	}
}

func (r *SaleItemViewRepositoryImpl) Create(ctx context.Context, view *view.SaleItemView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Insert(ctx, view, ops...)
}

func (r *SaleItemViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.SaleItemView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.InsertMany(ctx, views, ops...)
}

func (r *SaleItemViewRepositoryImpl) Update(ctx context.Context, view *view.SaleItemView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Update(ctx, view, ops...)
}

func (r *SaleItemViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.SaleItemView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.UpdateMany(ctx, views, ops...)
}

func (r *SaleItemViewRepositoryImpl) Delete(ctx context.Context, view *view.SaleItemView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId(), ops...)
}

func (r *SaleItemViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.SaleItemView, opts ...service.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opts...)
}

func (r *SaleItemViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, tenantId, id, ops...)
}

func (r *SaleItemViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByIds(ctx, tenantId, ids, ops...)
}

func (r *SaleItemViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteAll(ctx, tenantId, ops...)
}

func (r *SaleItemViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByFilter(ctx, tenantId, filter, ops...)
}
func (r *SaleItemViewRepositoryImpl) DeleteBySaleBillId(ctx context.Context, tenantId string, saleBillId string, opts ...service.Options) error {
	ops := newOptions(opts...)
	filter := fmt.Sprintf(`saleBillId == "%s"`, saleBillId)
	return r.dao.DeleteByFilter(ctx, tenantId, filter, ops...)
}

func (r *SaleItemViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opts ...service.Options) (*view.SaleItemView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindById(ctx, tenantId, id, ops...)
}

func (r *SaleItemViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) ([]*view.SaleItemView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindByIds(ctx, tenantId, ids, ops...)
}

func (r *SaleItemViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opts ...service.Options) ([]*view.SaleItemView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindAll(ctx, tenantId, ops...).Result()
}

func (r *SaleItemViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...service.Options) (*ddd_repository.FindPagingResult[*view.SaleItemView], bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindPaging(ctx, query, ops...).Result()
}
func (r *SaleItemViewRepositoryImpl) FindBySaleBillId(ctx context.Context, tenantId string, saleBillId string, opts ...service.Options) ([]*view.SaleItemView, bool, error) {
	filterMap := map[string]interface{}{
		"saleBillId": saleBillId,
	}
	ops := newOptions(opts...)
	return r.dao.FindListByMap(ctx, tenantId, filterMap, ops...).Result()
}
