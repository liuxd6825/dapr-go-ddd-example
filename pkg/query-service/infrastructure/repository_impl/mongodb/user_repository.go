package mongodb

import (
	"context"
	"errors"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	dr "github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

type UserRepository struct {
	repos dr.Repository
}

func NewUserRepository() repository.UserViewRepository {
	coll := GetMongoDB().GetCollection("users")
	return &UserRepository{
		repos: ddd_mongodb.NewRepository(&userViewBuilder{}, GetMongoDB(), coll),
	}
}

func (u *UserRepository) CreateById(ctx context.Context, user *projection.UserView) (resUser *projection.UserView, resErr error) {
	resErr = u.repos.Insert(ctx, user).OnSuccess(func(data interface{}) error {
		resUser = user
		return nil
	}).GetError()
	return
}

func (u *UserRepository) UpdateById(ctx context.Context, user *projection.UserView) (resUser *projection.UserView, resErr error) {
	resErr = u.repos.Update(ctx, user).OnSuccess(func(data interface{}) error {
		resUser = user
		return nil
	}).GetError()
	return
}

func (u *UserRepository) FindById(ctx context.Context, tenantId string, id string) (resUser *projection.UserView, ok bool, resErr error) {
	resErr = u.repos.FindById(ctx, tenantId, id).OnSuccess(func(data interface{}) error {
		resUser, ok = data.(*projection.UserView)
		if !ok {
			return errors.New("find result not is *projection.UserView")
		}
		return nil
	}).OnNotFond(func() error {
		ok = false
		return nil
	}).GetError()
	return
}

func (u *UserRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id).GetError()
}

func (u *UserRepository) FindPaging(ctx context.Context, query *dr.PagingQuery) (res *dr.PagingData, isFound bool, err error) {
	err = u.repos.FindPaging(ctx, query).OnSuccess(func(data *dr.PagingData) error {
		res = data
		isFound = true
		return nil
	}).OnNotFond(func() error {
		isFound = false
		return nil
	}).GetError()
	return
}

type userViewBuilder struct {
}

func (b *userViewBuilder) NewOne() interface{} {
	return &projection.UserView{}
}

func (b *userViewBuilder) NewList() interface{} {
	return &[]projection.UserView{}
}
