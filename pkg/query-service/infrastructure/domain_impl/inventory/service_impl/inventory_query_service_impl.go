package service_impl

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/query"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/repository"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/service"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	repos_impl "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/domain_impl/inventory/repository_impl/mongodb"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_repository"
	"sync"
)

//
// InventoryQueryDomainServiceImpl
// @Description: 查询领域服务实现类
//
type InventoryQueryDomainServiceImpl struct {
	repos repository.InventoryViewRepository
}

// 单例应用服务
var inventoryQueryDomainService service.InventoryQueryDomainService

// 并发安全
var onceInventory sync.Once

//
// GetInventoryQueryDomainService
// @Description: 获取单例领域服务
// @return service.InventoryQueryDomainService
//
func GetInventoryQueryDomainService() service.InventoryQueryDomainService {
	onceInventory.Do(func() {
		inventoryQueryDomainService = newInventoryQueryDomainService()
	})
	return inventoryQueryDomainService
}

func newInventoryQueryDomainService() service.InventoryQueryDomainService {
	return &InventoryQueryDomainServiceImpl{
		repos: repos_impl.NewInventoryViewRepository(),
	}
}

func (s *InventoryQueryDomainServiceImpl) Create(ctx context.Context, view *view.InventoryView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Create(ctx, view, opt)
}

func (s *InventoryQueryDomainServiceImpl) CreateMany(ctx context.Context, views []*view.InventoryView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.CreateMany(ctx, views, opt)
}

func (s *InventoryQueryDomainServiceImpl) Update(ctx context.Context, view *view.InventoryView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Update(ctx, view, opt)
}

func (s *InventoryQueryDomainServiceImpl) UpdateMany(ctx context.Context, views []*view.InventoryView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.UpdateMany(ctx, views, opt)
}

func (s *InventoryQueryDomainServiceImpl) Delete(ctx context.Context, view *view.InventoryView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.Delete(ctx, view, opt)
}

func (s *InventoryQueryDomainServiceImpl) DeleteMany(ctx context.Context, tenantId string, views []*view.InventoryView, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteMany(ctx, tenantId, views, opt)
}

func (s *InventoryQueryDomainServiceImpl) DeleteById(ctx context.Context, tenantId string, id string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteById(ctx, tenantId, id, opt)
}

func (s *InventoryQueryDomainServiceImpl) DeleteByIds(ctx context.Context, tenantId string, ids []string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByIds(ctx, tenantId, ids, opt)
}

func (s *InventoryQueryDomainServiceImpl) DeleteByFilter(ctx context.Context, tenantId, filter string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteByFilter(ctx, tenantId, filter, opt)
}

func (s *InventoryQueryDomainServiceImpl) DeleteAll(ctx context.Context, tenantId string, opts ...service.Options) error {
	opt := MergeOptions(opts...)
	return s.repos.DeleteAll(ctx, tenantId, opt)
}

func (s *InventoryQueryDomainServiceImpl) FindById(ctx context.Context, qry *query.InventoryFindByIdQuery, opts ...service.Options) (*view.InventoryView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindById(ctx, qry.TenantId, qry.Id, opt)
}

func (s *InventoryQueryDomainServiceImpl) FindByIds(ctx context.Context, qry *query.InventoryFindByIdsQuery, opts ...service.Options) ([]*view.InventoryView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindByIds(ctx, qry.TenantId, qry.Ids, opt)
}

func (s *InventoryQueryDomainServiceImpl) FindAll(ctx context.Context, qry *query.InventoryFindAllQuery, opts ...service.Options) ([]*view.InventoryView, bool, error) {
	opt := MergeOptions(opts...)
	return s.repos.FindAll(ctx, qry.TenantId, opt)
}

func (s *InventoryQueryDomainServiceImpl) FindPaging(ctx context.Context, qry *query.InventoryFindPagingQuery, opts ...service.Options) (*query.InventoryFindPagingResult, bool, error) {
	q := ddd_repository.NewFindPagingQuery()
	q.SetTenantId(qry.TenantId)
	q.SetFilter(qry.Filter)
	q.SetPageNum(qry.PageNum)
	q.SetFields(qry.Fields)
	q.SetSort(qry.Sort)
	q.SetPageSize(qry.PageSize)

	opt := MergeOptions(opts...)
	data, _, _ := s.repos.FindPaging(ctx, q, opt)

	res := query.NewInventoryFindPagingResult()
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
