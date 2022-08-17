package mongodb

import (
	"context"
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/dao/mongo_dao"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type SaleItemViewRepositoryImpl struct {
	dao *mongo_dao.Dao[*view.SaleItemView]
}

func NewSaleItemViewRepository() repository.SaleItemViewRepository {
	return &SaleItemViewRepositoryImpl{
		dao: mongo_dao.NewDao[*view.SaleItemView]("sale_bill"),
	}
}

func (r *SaleItemViewRepositoryImpl) Create(ctx context.Context, view *view.SaleItemView, opt repository.Options) error {
	return r.dao.Insert(ctx, view)
}

func (r *SaleItemViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.SaleItemView, opt repository.Options) error {
	return r.dao.InsertMany(ctx, views)
}

func (r *SaleItemViewRepositoryImpl) Update(ctx context.Context, view *view.SaleItemView, opt repository.Options) error {
	return r.dao.Update(ctx, view)
}

func (r *SaleItemViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.SaleItemView, opt repository.Options) error {
	return r.dao.UpdateMany(ctx, views)
}

func (r *SaleItemViewRepositoryImpl) Delete(ctx context.Context, view *view.SaleItemView, opt repository.Options) error {
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId())
}

func (r *SaleItemViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.SaleItemView, opt repository.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opt)
}

func (r *SaleItemViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opt repository.Options) error {
	return r.dao.DeleteById(ctx, tenantId, id)
}

func (r *SaleItemViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) error {
	return r.dao.DeleteByIds(ctx, tenantId, ids)
}

func (r *SaleItemViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opt repository.Options) error {
	return r.dao.DeleteAll(ctx, tenantId)
}

func (r *SaleItemViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opt repository.Options) error {
	return r.dao.DeleteByFilter(ctx, tenantId, filter)
}
func (r *SaleItemViewRepositoryImpl) DeleteBySaleBillId(ctx context.Context, tenantId string, saleBillId string, opt repository.Options) error {
	filter := fmt.Sprintf("saleBillId = %s", saleBillId)
	return r.dao.DeleteByFilter(ctx, tenantId, filter)
}

func (r *SaleItemViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opt repository.Options) (*view.SaleItemView, bool, error) {
	return r.dao.FindById(ctx, tenantId, id)
}

func (r *SaleItemViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) ([]*view.SaleItemView, bool, error) {
	return r.dao.FindByIds(ctx, tenantId, ids)
}

func (r *SaleItemViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opt repository.Options) ([]*view.SaleItemView, bool, error) {
	return r.dao.FindAll(ctx, tenantId).Result()
}

func (r *SaleItemViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opt repository.Options) (*ddd_repository.FindPagingResult[*view.SaleItemView], bool, error) {
	return r.dao.FindPaging(ctx, query).Result()
}
func (r *SaleItemViewRepositoryImpl) FindBySaleBillId(ctx context.Context, tenantId string, saleBillId string, opt repository.Options) ([]*view.SaleItemView, bool, error) {
	filterMap := map[string]interface{}{
		"saleBillId": saleBillId,
	}
	return r.dao.FindListByMap(ctx, tenantId, filterMap).Result()
}
