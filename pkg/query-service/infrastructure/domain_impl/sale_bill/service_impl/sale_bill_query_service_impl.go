package service_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	repos_impl "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/sale_bill/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"sync"
)

//
// SaleBillQueryDomainServiceImpl
// @Description: 查询领域服务实现类
//
type SaleBillQueryDomainServiceImpl struct {
	repos repository.SaleBillViewRepository
}

// 单例应用服务
var saleBillQueryDomainService service.SaleBillQueryDomainService

// 并发安全
var onceSaleBill sync.Once

//
// GetSaleBillQueryDomainService
// @Description: 获取单例领域服务
// @return service.SaleBillQueryDomainService
//
func GetSaleBillQueryDomainService() service.SaleBillQueryDomainService {
	onceSaleBill.Do(func() {
		saleBillQueryDomainService = newSaleBillQueryDomainService()
	})
	return saleBillQueryDomainService
}

func newSaleBillQueryDomainService() service.SaleBillQueryDomainService {
	return &SaleBillQueryDomainServiceImpl{
		repos: repos_impl.NewSaleBillViewRepository(),
	}
}

func (s *SaleBillQueryDomainServiceImpl) Create(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Create(ctx, view, opt)
}

func (s *SaleBillQueryDomainServiceImpl) CreateMany(ctx context.Context, views []*view.SaleBillView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.CreateMany(ctx, views, opt)
}

func (s *SaleBillQueryDomainServiceImpl) Update(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Update(ctx, view, opt)
}

func (s *SaleBillQueryDomainServiceImpl) UpdateMany(ctx context.Context, views []*view.SaleBillView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.UpdateMany(ctx, views, opt)
}

func (s *SaleBillQueryDomainServiceImpl) Delete(ctx context.Context, view *view.SaleBillView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Delete(ctx, view, opt)
}

func (s *SaleBillQueryDomainServiceImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.SaleBillView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteMany(ctx, tenantId, views, opt)
}

func (s *SaleBillQueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteById(ctx, tenantId, id, opt)
}

func (s *SaleBillQueryDomainServiceImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByIds(ctx, tenantId, ids, opt)
}

func (s *SaleBillQueryDomainServiceImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByFilter(ctx, tenantId, filter, opt)
}

func (s *SaleBillQueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteAll(ctx, tenantId, opt)
}

func (s *SaleBillQueryDomainServiceImpl) FindById(ctx context.Context, qry *query.SaleBillFindByIdQuery, opts ...service.Options) (*view.SaleBillView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindById(ctx, qry.TenantId, qry.Id, opt)
}

func (s *SaleBillQueryDomainServiceImpl) FindByIds(ctx context.Context, qry *query.SaleBillFindByIdsQuery, opts ...service.Options) ([]*view.SaleBillView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindByIds(ctx, qry.TenantId, qry.Ids, opt)
}

func (s *SaleBillQueryDomainServiceImpl) FindAll(ctx context.Context, qry *query.SaleBillFindAllQuery, opts ...service.Options) ([]*view.SaleBillView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindAll(ctx, qry.TenantId, opt)
}

func (s *SaleBillQueryDomainServiceImpl) FindPaging(ctx context.Context, qry *query.SaleBillFindPagingQuery, opts ...service.Options) (*query.SaleBillFindPagingResult, bool, error) {
	q := ddd_repository.NewFindPagingQuery()
	q.SetTenantId(qry.TenantId)
	q.SetFilter(qry.Filter)
	q.SetPageNum(qry.PageNum)
	q.SetFields(qry.Fields)
	q.SetSort(qry.Sort)
	q.SetPageSize(qry.PageSize)

	opt := MergeOptions(opts...)
	data, _, _ := s.repos.FindPaging(ctx, q, opt)

	res := query.NewSaleBillFindPagingResult()
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
