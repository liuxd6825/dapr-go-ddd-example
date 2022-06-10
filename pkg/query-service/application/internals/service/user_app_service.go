package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/service"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"sync"
)

type UserQueryAppService struct {
	service service.UserQueryDomainService
}

var _UserQueryAppService *UserQueryAppService
var onceUser sync.Once

//
// GetUserQueryAppService
// @Description: 获取单实例应用服务
// @return *UserQueryAppService
//
func GetUserQueryAppService() *UserQueryAppService {
	onceUser.Do(func() {
		_UserQueryAppService = newUserQueryAppService()
	})
	return _UserQueryAppService
}

//
// NewUserQueryAppService
// @Description: 创建应用服务
// @return *UserQueryAppService
//
func newUserQueryAppService() *UserQueryAppService {
	return &UserQueryAppService{
		service: service.NewUserQueryService(),
	}
}

func (a *UserQueryAppService) Create(ctx context.Context, v *view.UserView) error {
	return a.service.Create(ctx, v)
}

func (a *UserQueryAppService) Update(ctx context.Context, v *view.UserView) error {
	return a.service.Update(ctx, v)
}

func (a *UserQueryAppService) DeleteById(ctx context.Context, tenantId string, userId string) error {
	return a.service.DeleteById(ctx, tenantId, userId)
}

func (a *UserQueryAppService) FindById(ctx context.Context, tenantId string, userId string) (*view.UserView, bool, error) {
	return a.service.FindById(ctx, tenantId, userId)
}

func (a *UserQueryAppService) FindAll(ctx context.Context, tenantId string) (*[]*view.UserView, bool, error) {
	return a.service.FindAll(ctx, tenantId)
}

func (a *UserQueryAppService) FindPaging(ctx context.Context, query ddd_repository.FindPagingQuery) (*ddd_repository.FindPagingResult[*view.UserView], bool, error) {
	return a.service.FindPaging(ctx, query)
}
