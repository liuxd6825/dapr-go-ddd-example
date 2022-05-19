package queryappservice

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
)

type UserQueryAppService interface {
	QueryAppService[*projection.UserView]
	GetById(ctx context.Context, tenantId, id string) (data *projection.UserView, isFound bool, err error)
}

type userQueryAppService struct {
	BaseQueryAppService[*projection.UserView]
}

var _userQueryAppService UserQueryAppService

func init() {
	_userQueryAppService = newUserQueryAppService()
}

func GetUserQueryAppService() UserQueryAppService {
	return _userQueryAppService
}

func newUserQueryAppService() UserQueryAppService {
	res := &userQueryAppService{}
	res.appId = "query-service"
	res.resourceName = "users"
	res.apiVersion = "v1.0"
	return res
}

func (s *userQueryAppService) GetById(ctx context.Context, tenantId, id string) (data *projection.UserView, isFound bool, err error) {
	data = &projection.UserView{}
	isFound, err = s.GetData(ctx, tenantId, id, data)
	return
}
