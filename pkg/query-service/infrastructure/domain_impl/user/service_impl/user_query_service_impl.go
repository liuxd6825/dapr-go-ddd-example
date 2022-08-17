package service_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	repos_impl "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/user/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"sync"
)

//
// UserQueryDomainServiceImpl
// @Description: 查询领域服务实现类
//
type UserQueryDomainServiceImpl struct {
	repos repository.UserViewRepository
}

// 单例应用服务
var userQueryDomainService service.UserQueryDomainService

// 并发安全
var onceUser sync.Once

//
// GetUserQueryDomainService
// @Description: 获取单例领域服务
// @return service.UserQueryDomainService
//
func GetUserQueryDomainService() service.UserQueryDomainService {
	onceUser.Do(func() {
		userQueryDomainService = newUserQueryDomainService()
	})
	return userQueryDomainService
}

func newUserQueryDomainService() service.UserQueryDomainService {
	return &UserQueryDomainServiceImpl{
		repos: repos_impl.NewUserViewRepository(),
	}
}

func (s *UserQueryDomainServiceImpl) Create(ctx context.Context, view *view.UserView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Create(ctx, view, opt)
}

func (s *UserQueryDomainServiceImpl) CreateMany(ctx context.Context, views []*view.UserView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.CreateMany(ctx, views, opt)
}

func (s *UserQueryDomainServiceImpl) Update(ctx context.Context, view *view.UserView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Update(ctx, view, opt)
}

func (s *UserQueryDomainServiceImpl) UpdateMany(ctx context.Context, views []*view.UserView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.UpdateMany(ctx, views, opt)
}

func (s *UserQueryDomainServiceImpl) Delete(ctx context.Context, view *view.UserView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Delete(ctx, view, opt)
}

func (s *UserQueryDomainServiceImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.UserView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteMany(ctx, tenantId, views, opt)
}

func (s *UserQueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteById(ctx, tenantId, id, opt)
}

func (s *UserQueryDomainServiceImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByIds(ctx, tenantId, ids, opt)
}

func (s *UserQueryDomainServiceImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByFilter(ctx, tenantId, filter, opt)
}

func (s *UserQueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteAll(ctx, tenantId, opt)
}

func (s *UserQueryDomainServiceImpl) FindById(ctx context.Context, qry *query.UserFindByIdQuery, opts ...service.Options) (*view.UserView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindById(ctx, qry.TenantId, qry.Id, opt)
}

func (s *UserQueryDomainServiceImpl) FindByIds(ctx context.Context, qry *query.UserFindByIdsQuery, opts ...service.Options) ([]*view.UserView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindByIds(ctx, qry.TenantId, qry.Ids, opt)
}

func (s *UserQueryDomainServiceImpl) FindAll(ctx context.Context, qry *query.UserFindAllQuery, opts ...service.Options) ([]*view.UserView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindAll(ctx, qry.TenantId, opt)
}

func (s *UserQueryDomainServiceImpl) FindPaging(ctx context.Context, qry *query.UserFindPagingQuery, opts ...service.Options) (*query.UserFindPagingResult, bool, error) {
	q := ddd_repository.NewFindPagingQuery()
	q.SetTenantId(qry.TenantId)
	q.SetFilter(qry.Filter)
	q.SetPageNum(qry.PageNum)
	q.SetFields(qry.Fields)
	q.SetSort(qry.Sort)
	q.SetPageSize(qry.PageSize)

	opt := MergeOptions(opts...)
	data, _, _ := s.repos.FindPaging(ctx, q, opt)

	res := query.NewUserFindPagingResult()
	res.Data = data.GetData()
	res.TotalRows = data.GetTotalRows()
	res.TotalPages = data.GetTotalPages()
	res.PageNum = data.GetPageNum()
	res.PageSize = data.GetPageSize()
	res.Filter = data.GetFilter()
	res.Sort = data.GetSort()
	res.Error = data.GetError()
	res.IsFound = data.GetIsFound()

	return res, res.IsFound, res.Error
}
