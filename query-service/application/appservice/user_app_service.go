package appservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
)

type UserAppQueryService struct {
	service *queryservice.UserQueryService
}

func NewUserAppQueryService() *UserAppQueryService {
	return &UserAppQueryService{
		service: queryservice.NewUserQueryService(),
	}
}

func (a *UserAppQueryService) FindById(ctx context.Context, tenantId string, userId string) (*projection.UserView, bool, error) {
	return a.service.FindById(ctx, tenantId, userId)
}
