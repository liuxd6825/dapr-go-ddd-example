package mongodb

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/repository"
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

func (r *InventoryViewRepositoryImpl) Create(ctx context.Context, view *view.InventoryView, opt repository.Options) error {
	return r.dao.Insert(ctx, view)
}

func (r *InventoryViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.InventoryView, opt repository.Options) error {
	return r.dao.InsertMany(ctx, views)
}

func (r *InventoryViewRepositoryImpl) Update(ctx context.Context, view *view.InventoryView, opt repository.Options) error {
	return r.dao.Update(ctx, view)
}

func (r *InventoryViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.InventoryView, opt repository.Options) error {
	return r.dao.UpdateMany(ctx, views)
}

func (r *InventoryViewRepositoryImpl) Delete(ctx context.Context, view *view.InventoryView, opt repository.Options) error {
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId())
}

func (r *InventoryViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.InventoryView, opt repository.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opt)
}

func (r *InventoryViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opt repository.Options) error {
	return r.dao.DeleteById(ctx, tenantId, id)
}

func (r *InventoryViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) error {
	return r.dao.DeleteByIds(ctx, tenantId, ids)
}

func (r *InventoryViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opt repository.Options) error {
	return r.dao.DeleteAll(ctx, tenantId)
}

func (r *InventoryViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opt repository.Options) error {
	return r.dao.DeleteByFilter(ctx, tenantId, filter)
}

func (r *InventoryViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opt repository.Options) (*view.InventoryView, bool, error) {
	return r.dao.FindById(ctx, tenantId, id)
}

func (r *InventoryViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) ([]*view.InventoryView, bool, error) {
	return r.dao.FindByIds(ctx, tenantId, ids)
}

func (r *InventoryViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opt repository.Options) ([]*view.InventoryView, bool, error) {
	return r.dao.FindAll(ctx, tenantId).Result()
}

func (r *InventoryViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opt repository.Options) (*ddd_repository.FindPagingResult[*view.InventoryView], bool, error) {
	return r.dao.FindPaging(ctx, query).Result()
}
