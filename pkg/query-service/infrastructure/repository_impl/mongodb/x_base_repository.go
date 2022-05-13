package mongodb

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func GetMongoDB() *ddd_mongodb.MongoDB {
	return restapp.GetMongoDB()
}

type BaseRepository[T ddd.Entity] struct {
	super *ddd_mongodb.Repository[T]
}

func NewBaseRepository[T ddd.Entity](newFunc func() T, collectionName string, opts ...*RepositoryOptions) *BaseRepository[T] {
	options := NewRepositoryOptions()
	options.Merge(opts...)
	coll := options.mongoDB.GetCollection(collectionName)
	return &BaseRepository[T]{
		super: ddd_mongodb.NewRepository[T](newFunc, options.mongoDB, coll),
	}
}

func (u *BaseRepository[T]) CreateById(ctx context.Context, entity T) (T, error) {
	return u.super.Insert(ctx, entity).Result()
}

func (u *BaseRepository[T]) UpdateById(ctx context.Context, entity T) (T, error) {
	return u.super.Update(ctx, entity).Result()
}

func (u *BaseRepository[T]) FindById(ctx context.Context, tenantId string, id string) (T, bool, error) {
	return u.super.FindById(ctx, tenantId, id).Result()
}

func (u *BaseRepository[T]) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.super.DeleteById(ctx, tenantId, id).GetError()
}

func (u *BaseRepository[T]) FindPaging(ctx context.Context, query *ddd_repository.FindPagingQuery) *ddd_repository.FindPagingResult[T] {
	return u.super.FindPaging(ctx, query)
}
