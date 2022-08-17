package mongo_dao

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type Dao[T ddd.Entity] struct {
	dao *ddd_mongodb.Dao[T]
}

func NewDao[T ddd.Entity](collectionName string, opts ...*RepositoryOptions) *Dao[T] {
	options := NewRepositoryOptions()
	options.Merge(opts...)
	coll := options.mongoDB.GetCollection(collectionName)
	return &Dao[T]{
		dao: ddd_mongodb.NewDao[T](options.mongoDB, coll),
	}
}

func (d *Dao[T]) Insert(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) error {
	return d.dao.Insert(ctx, entity, opts...).GetError()
}

func (d *Dao[T]) InsertMany(ctx context.Context, entity []T, opts ...*ddd_repository.SetOptions) error {
	return d.dao.InsertMany(ctx, entity, opts...).GetError()
}

func (d *Dao[T]) Update(ctx context.Context, entity T, opts ...*ddd_repository.SetOptions) error {
	return d.dao.Update(ctx, entity, opts...).GetError()
}

func (d *Dao[T]) UpdateMany(ctx context.Context, entity []T, opts ...*ddd_repository.SetOptions) error {
	return d.dao.UpdateManyById(ctx, entity, opts...).GetError()
}

func (d *Dao[T]) UpdateManyByFilter(ctx context.Context, tenantId, filter string, data interface{}, opts ...*ddd_repository.SetOptions) error {
	return d.dao.UpdateManyByFilter(ctx, tenantId, filter, data, opts...).GetError()
}

func (d *Dao[T]) DeleteById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteById(ctx, tenantId, id, opts...).GetError()
}

func (d *Dao[T]) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteByIds(ctx, tenantId, ids, opts...)
}

func (d *Dao[T]) DeleteAll(ctx context.Context, tenantId string, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteAll(ctx, tenantId, opts...).GetError()
}

func (d *Dao[T]) DeleteByFilter(ctx context.Context, tenantId string, filter string, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteByFilter(ctx, tenantId, filter, opts...)
}

func (d *Dao[T]) DeleteByMap(ctx context.Context, tenantId string, filterMap map[string]interface{}, opts ...*ddd_repository.SetOptions) error {
	return d.dao.DeleteByMap(ctx, tenantId, filterMap, opts...).GetError()
}

func (d *Dao[T]) FindById(ctx context.Context, tenantId string, id string, opts ...*ddd_repository.FindOptions) (T, bool, error) {
	return d.dao.FindById(ctx, tenantId, id, opts...).Result()
}

func (d *Dao[T]) FindByIds(ctx context.Context, tenantId string, ids []string) ([]T, bool, error) {
	return d.dao.FindByIds(ctx, tenantId, ids).Result()
}

func (d *Dao[T]) FindAll(ctx context.Context, tenantId string, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	return d.dao.FindAll(ctx, tenantId, opts...)
}

func (d *Dao[T]) FindListByMap(ctx context.Context, tenantId string, filterMap map[string]interface{}, opts ...*ddd_repository.FindOptions) *ddd_repository.FindListResult[T] {
	return d.dao.FindListByMap(ctx, tenantId, filterMap, opts...)
}

func (d *Dao[T]) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery, opts ...*ddd_repository.FindOptions) *ddd_repository.FindPagingResult[T] {
	return d.dao.FindPaging(ctx, query, opts...)
}

type RepositoryOptions struct {
	mongoDB *ddd_mongodb.MongoDB
}

func NewRepositoryOptions() *RepositoryOptions {
	return &RepositoryOptions{}
}

func (o *RepositoryOptions) SetMongoDB(mongoDB *ddd_mongodb.MongoDB) *RepositoryOptions {
	o.mongoDB = mongoDB
	return o
}

func (o *RepositoryOptions) Merge(opts ...*RepositoryOptions) *RepositoryOptions {
	if opts != nil {
		for _, item := range opts {
			if item.mongoDB != nil {
				o.mongoDB = item.mongoDB
			}
		}
	}
	if o.mongoDB == nil {
		o.mongoDB = GetMongoDB()
	}
	return o
}

func GetMongoDB() *ddd_mongodb.MongoDB {
	return restapp.GetMongoDB()
}

func NewSession(isWrite bool) ddd_repository.Session {
	return ddd_mongodb.NewSession(isWrite, GetMongoDB())
}
