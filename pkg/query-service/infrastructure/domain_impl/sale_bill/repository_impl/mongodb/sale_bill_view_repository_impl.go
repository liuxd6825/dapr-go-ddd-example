package mongodb

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/repository"
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

func (r *SaleBillViewRepositoryImpl) Create(ctx context.Context, view *view.SaleBillView, opt repository.Options) error {
	return r.dao.Insert(ctx, view)
}

func (r *SaleBillViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.SaleBillView, opt repository.Options) error {
	return r.dao.InsertMany(ctx, views)
}

func (r *SaleBillViewRepositoryImpl) Update(ctx context.Context, view *view.SaleBillView, opt repository.Options) error {
	return r.dao.Update(ctx, view)
}

func (r *SaleBillViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.SaleBillView, opt repository.Options) error {
	return r.dao.UpdateMany(ctx, views)
}

func (r *SaleBillViewRepositoryImpl) Delete(ctx context.Context, view *view.SaleBillView, opt repository.Options) error {
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId())
}

func (r *SaleBillViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.SaleBillView, opt repository.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opt)
}

func (r *SaleBillViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opt repository.Options) error {
	return r.dao.DeleteById(ctx, tenantId, id)
}

func (r *SaleBillViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) error {
	return r.dao.DeleteByIds(ctx, tenantId, ids)
}

func (r *SaleBillViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opt repository.Options) error {
	return r.dao.DeleteAll(ctx, tenantId)
}

func (r *SaleBillViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opt repository.Options) error {
	return r.dao.DeleteByFilter(ctx, tenantId, filter)
}

func (r *SaleBillViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opt repository.Options) (*view.SaleBillView, bool, error) {
	return r.dao.FindById(ctx, tenantId, id)
}

func (r *SaleBillViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) ([]*view.SaleBillView, bool, error) {
	return r.dao.FindByIds(ctx, tenantId, ids)
}

func (r *SaleBillViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opt repository.Options) ([]*view.SaleBillView, bool, error) {
	return r.dao.FindAll(ctx, tenantId).Result()
}

func (r *SaleBillViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opt repository.Options) (*ddd_repository.FindPagingResult[*view.SaleBillView], bool, error) {
	return r.dao.FindPaging(ctx, query).Result()
}
