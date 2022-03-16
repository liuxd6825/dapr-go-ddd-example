package mongodb

import (
	"context"
	"errors"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/repository"
	dr "github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

type UserRepository struct {
	dr.BaseRepository
}

func NewUserRepository() repository.UserViewRepository {
	coll := GetMongoDB().NewCollection("users")
	return &UserRepository{
		BaseRepository: ddd_mongodb.NewBaseRepository(&userViewBuilder{}, coll),
	}
}

func (u *UserRepository) CreateById(ctx context.Context, user *projection.UserView) (resUser *projection.UserView, resErr error) {
	resErr = u.BaseRepository.BaseCreate(ctx, user).OnSuccess(func(data interface{}) error {
		resUser = user
		return nil
	}).GetError()
	return
}

func (u *UserRepository) UpdateById(ctx context.Context, user *projection.UserView) (resUser *projection.UserView, resErr error) {
	resErr = u.BaseRepository.BaseUpdate(ctx, user).OnSuccess(func(data interface{}) error {
		resUser = user
		return nil
	}).GetError()
	return
}

func (u *UserRepository) FindById(ctx context.Context, tenantId string, id string) (resUser *projection.UserView, resErr error) {
	resErr = u.BaseRepository.BaseFindById(ctx, tenantId, id).OnSuccess(func(data interface{}) error {
		var ok bool
		resUser, ok = data.(*projection.UserView)
		if !ok {
			return errors.New("find result not is *projection.UserView")
		}
		return nil
	}).GetError()
	return
}

func (u *UserRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.BaseRepository.BaseDeleteById(ctx, tenantId, id).GetError()
}

type userViewBuilder struct {
}

func (b *userViewBuilder) New() interface{} {
	return &projection.UserView{}
}
func (b *userViewBuilder) NewList() interface{} {
	return &[]projection.UserView{}
}
