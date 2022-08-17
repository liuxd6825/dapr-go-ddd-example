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
// SaleItemQueryDomainServiceImpl
// @Description: 查询领域服务实现类
//
type SaleItemQueryDomainServiceImpl struct {
	repos repository.SaleItemViewRepository
}

// 单例应用服务
var saleItemQueryDomainService service.SaleItemQueryDomainService

// 并发安全
var onceSaleItem sync.Once

//
// GetSaleItemQueryDomainService
// @Description: 获取单例领域服务
// @return service.SaleItemQueryDomainService
//
func GetSaleItemQueryDomainService() service.SaleItemQueryDomainService {
	onceSaleItem.Do(func() {
		saleItemQueryDomainService = newSaleItemQueryDomainService()
	})
	return saleItemQueryDomainService
}

func newSaleItemQueryDomainService() service.SaleItemQueryDomainService {
	return &SaleItemQueryDomainServiceImpl{
		repos: repos_impl.NewSaleItemViewRepository(),
	}
}

func (s *SaleItemQueryDomainServiceImpl) Create(ctx context.Context, view *view.SaleItemView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Create(ctx, view, opt)
}

func (s *SaleItemQueryDomainServiceImpl) CreateMany(ctx context.Context, views []*view.SaleItemView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.CreateMany(ctx, views, opt)
}

func (s *SaleItemQueryDomainServiceImpl) Update(ctx context.Context, view *view.SaleItemView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Update(ctx, view, opt)
}

func (s *SaleItemQueryDomainServiceImpl) UpdateMany(ctx context.Context, views []*view.SaleItemView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.UpdateMany(ctx, views, opt)
}

func (s *SaleItemQueryDomainServiceImpl) Delete(ctx context.Context, view *view.SaleItemView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Delete(ctx, view, opt)
}

func (s *SaleItemQueryDomainServiceImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.SaleItemView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteMany(ctx, tenantId, views, opt)
}

func (s *SaleItemQueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteById(ctx, tenantId, id, opt)
}

func (s *SaleItemQueryDomainServiceImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByIds(ctx, tenantId, ids, opt)
}
func (s *SaleItemQueryDomainServiceImpl) DeleteBySaleBillId(ctx context.Context, tenantId string, saleBillId string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteBySaleBillId(ctx, tenantId, saleBillId, opt)
}

func (s *SaleItemQueryDomainServiceImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByFilter(ctx, tenantId, filter, opt)
}

func (s *SaleItemQueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteAll(ctx, tenantId, opt)
}

func (s *SaleItemQueryDomainServiceImpl) FindById(ctx context.Context, qry *query.SaleItemFindByIdQuery, opts ...service.Options) (*view.SaleItemView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindById(ctx, qry.TenantId, qry.Id, opt)
}

func (s *SaleItemQueryDomainServiceImpl) FindByIds(ctx context.Context, qry *query.SaleItemFindByIdsQuery, opts ...service.Options) ([]*view.SaleItemView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindByIds(ctx, qry.TenantId, qry.Ids, opt)
}

func (s *SaleItemQueryDomainServiceImpl) FindAll(ctx context.Context, qry *query.SaleItemFindAllQuery, opts ...service.Options) ([]*view.SaleItemView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindAll(ctx, qry.TenantId, opt)
}

func (s *SaleItemQueryDomainServiceImpl) FindPaging(ctx context.Context, qry *query.SaleItemFindPagingQuery, opts ...service.Options) (*query.SaleItemFindPagingResult, bool, error) {
	q := ddd_repository.NewFindPagingQuery()
	q.SetTenantId(qry.TenantId)
	q.SetFilter(qry.Filter)
	q.SetPageNum(qry.PageNum)
	q.SetFields(qry.Fields)
	q.SetSort(qry.Sort)
	q.SetPageSize(qry.PageSize)

	opt := MergeOptions(opts...)
	data, _, _ := s.repos.FindPaging(ctx, q, opt)

	res := query.NewSaleItemFindPagingResult()
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
func (s *SaleItemQueryDomainServiceImpl) FindBySaleBillId(ctx context.Context, qry *query.SaleItemFindBySaleBillIdQuery, opts ...service.Options) ([]*view.SaleItemView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindBySaleBillId(ctx, qry.TenantId, qry.SaleBillId, opt)
}
