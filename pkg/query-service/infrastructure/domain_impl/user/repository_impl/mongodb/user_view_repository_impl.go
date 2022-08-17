package mongodb

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/db/dao/mongo_dao"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type UserViewRepositoryImpl struct {
	dao *mongo_dao.Dao[*view.UserView]
}

func NewUserViewRepository() repository.UserViewRepository {
	return &UserViewRepositoryImpl{
		dao: mongo_dao.NewDao[*view.UserView]("user"),
	}
}

func (r *UserViewRepositoryImpl) Create(ctx context.Context, view *view.UserView, opt repository.Options) error {
	return r.dao.Insert(ctx, view)
}

func (r *UserViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.UserView, opt repository.Options) error {
	return r.dao.InsertMany(ctx, views)
}

func (r *UserViewRepositoryImpl) Update(ctx context.Context, view *view.UserView, opt repository.Options) error {
	return r.dao.Update(ctx, view)
}

func (r *UserViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.UserView, opt repository.Options) error {
	return r.dao.UpdateMany(ctx, views)
}

func (r *UserViewRepositoryImpl) Delete(ctx context.Context, view *view.UserView, opt repository.Options) error {
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId())
}

func (r *UserViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.UserView, opt repository.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opt)
}

func (r *UserViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opt repository.Options) error {
	return r.dao.DeleteById(ctx, tenantId, id)
}

func (r *UserViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) error {
	return r.dao.DeleteByIds(ctx, tenantId, ids)
}

func (r *UserViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opt repository.Options) error {
	return r.dao.DeleteAll(ctx, tenantId)
}

func (r *UserViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opt repository.Options) error {
	return r.dao.DeleteByFilter(ctx, tenantId, filter)
}

func (r *UserViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opt repository.Options) (*view.UserView, bool, error) {
	return r.dao.FindById(ctx, tenantId, id)
}

func (r *UserViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opt repository.Options) ([]*view.UserView, bool, error) {
	return r.dao.FindByIds(ctx, tenantId, ids)
}

func (r *UserViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opt repository.Options) ([]*view.UserView, bool, error) {
	return r.dao.FindAll(ctx, tenantId).Result()
}

func (r *UserViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opt repository.Options) (*ddd_repository.FindPagingResult[*view.UserView], bool, error) {
	return r.dao.FindPaging(ctx, query).Result()
}
