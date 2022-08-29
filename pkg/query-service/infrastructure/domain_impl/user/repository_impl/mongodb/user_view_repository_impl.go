package mongodb

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
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

func (r *UserViewRepositoryImpl) Create(ctx context.Context, view *view.UserView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Insert(ctx, view, ops...)
}

func (r *UserViewRepositoryImpl) CreateMany(ctx context.Context, views []*view.UserView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.InsertMany(ctx, views, ops...)
}

func (r *UserViewRepositoryImpl) Update(ctx context.Context, view *view.UserView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.Update(ctx, view, ops...)
}

func (r *UserViewRepositoryImpl) UpdateMany(ctx context.Context, views []*view.UserView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.UpdateMany(ctx, views, ops...)
}

func (r *UserViewRepositoryImpl) Delete(ctx context.Context, view *view.UserView, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, view.GetTenantId(), view.GetId(), ops...)
}

func (r *UserViewRepositoryImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.UserView, opts ...service.Options) error {
	ids, err := ddd_repository.NewIds(ctx, views)
	if err != nil {
		return err
	}
	return r.DeleteByIds(ctx, tenantId, ids, opts...)
}

func (r *UserViewRepositoryImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteById(ctx, tenantId, id, ops...)
}

func (r *UserViewRepositoryImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByIds(ctx, tenantId, ids, ops...)
}

func (r *UserViewRepositoryImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteAll(ctx, tenantId, ops...)
}

func (r *UserViewRepositoryImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	ops := newOptions(opts...)
	return r.dao.DeleteByFilter(ctx, tenantId, filter, ops...)
}

func (r *UserViewRepositoryImpl) FindById(ctx context.Context, tenantId string, id string, opts ...service.Options) (*view.UserView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindById(ctx, tenantId, id, ops...)
}

func (r *UserViewRepositoryImpl) FindByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) ([]*view.UserView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindByIds(ctx, tenantId, ids, ops...)
}

func (r *UserViewRepositoryImpl) FindAll(ctx context.Context, tenantId string, opts ...service.Options) ([]*view.UserView, bool, error) {
	ops := newOptions(opts...)
	return r.dao.FindAll(ctx, tenantId, ops...).Result()
}

func (r *UserViewRepositoryImpl) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...service.Options) (*ddd_repository.FindPagingResult[*view.UserView], bool, error) {
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
