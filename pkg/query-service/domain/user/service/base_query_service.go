package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
)

type BaseQueryDomainService[T ddd.Entity] interface {
	FindById(ctx context.Context, tenantId, userId string) (T, bool, error)
	Create(ctx context.Context, user T) error
	Update(ctx context.Context, user T) error
	DeleteById(ctx context.Context, tenantId string, id string) error
	FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[T], bool, error)
}

type BaseQueryService2Impl[T ddd.Entity] struct {
	emptyEntity T
}

func (s *BaseQueryService2Impl[T]) FindById(ctx context.Context, tenantId, userId string) (T, bool, error) {

	return s.emptyEntity, false, nil
}
func (s *BaseQueryService2Impl[T]) Create(ctx context.Context, user T) error {
	return nil
}
func (s *BaseQueryService2Impl[T]) Update(ctx context.Context, user T) error {
	return nil
}
func (s *BaseQueryService2Impl[T]) DeleteById(ctx context.Context, tenantId string, id string) error {
	return nil
}
func (s *BaseQueryService2Impl[T]) FindPagingData(ctx context.Context, query *ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[T], bool, error) {
	return nil, false, nil
}
