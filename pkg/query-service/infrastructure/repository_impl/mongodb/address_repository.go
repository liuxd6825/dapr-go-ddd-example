package mongodb

import (
	"context"
	"errors"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/repository"
	dr "github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository/ddd_mongodb"
)

type AddressRepository struct {
	repos dr.Repository
}

func NewAddressRepository() repository.AddressViewRepository {
	coll := GetMongoDB().GetCollection("address")
	return &AddressRepository{
		repos: ddd_mongodb.NewRepository(&addressViewBuilder{}, GetMongoDB(), coll),
	}
}

func (u *AddressRepository) CreateById(ctx context.Context, addr *projection.AddressView) (res *projection.AddressView, resErr error) {
	resErr = u.repos.Insert(ctx, addr).OnSuccess(func(data interface{}) error {
		res = addr
		return nil
	}).GetError()
	return
}

func (u *AddressRepository) UpdateById(ctx context.Context, addr *projection.AddressView) (res *projection.AddressView, resErr error) {
	resErr = u.repos.Update(ctx, addr).OnSuccess(func(data interface{}) error {
		res = addr
		return nil
	}).GetError()
	return
}

func (u *AddressRepository) FindById(ctx context.Context, tenantId string, id string) (resUser *projection.AddressView, ok bool, resErr error) {
	resErr = u.repos.FindById(ctx, tenantId, id).OnSuccess(func(data interface{}) error {
		resUser, ok = data.(*projection.AddressView)
		if !ok {
			return errors.New("find result not is *projection.AddressView")
		}
		return nil
	}).OnNotFond(func() error {
		ok = false
		return nil
	}).GetError()
	return
}

func (u *AddressRepository) DeleteById(ctx context.Context, tenantId string, id string) error {
	return u.repos.DeleteById(ctx, tenantId, id).GetError()
}

func (u *AddressRepository) FindPaging(ctx context.Context, query *dr.PagingQuery) (res *dr.PagingData, isFound bool, err error) {
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

type addressViewBuilder struct {
}

func (b *addressViewBuilder) NewOne() interface{} {
	return &projection.AddressView{}
}

func (b *addressViewBuilder) NewList() interface{} {
	return &[]projection.AddressView{}
}
