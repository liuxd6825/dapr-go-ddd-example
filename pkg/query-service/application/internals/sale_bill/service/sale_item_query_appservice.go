package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/executor"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/executor/sale_item_impl"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"sync"
)

//
// SaleItemQueryAppService
// @Description: 销售明细项查询应用服务类
//
type SaleItemQueryAppService struct {
	saleItemCreateExecutor     executor.SaleItemCreateExecutor
	saleItemCreateManyExecutor executor.SaleItemCreateManyExecutor

	saleItemUpdateExecutor     executor.SaleItemUpdateExecutor
	saleItemUpdateManyExecutor executor.SaleItemUpdateManyExecutor

	saleItemDeleteByIdExecutor         executor.SaleItemDeleteByIdExecutor
	saleItemDeleteManyExecutor         executor.SaleItemDeleteManyExecutor
	saleItemDeleteAllExecutor          executor.SaleItemDeleteAllExecutor
	saleItemDeleteBySaleBillIdExecutor executor.SaleItemDeleteBySaleBillIdExecutor

	saleItemFindAllExecutor          executor.SaleItemFindAllExecutor
	saleItemFindByIdExecutor         executor.SaleItemFindByIdExecutor
	saleItemFindByIdsExecutor        executor.SaleItemFindByIdsExecutor
	saleItemFindPagingExecutor       executor.SaleItemFindPagingExecutor
	saleItemFindBySaleBillIdExecutor executor.SaleItemFindBySaleBillIdExecutor
}

// 单例应用服务
var saleItemQueryAppService *SaleItemQueryAppService

// 并发安全
var onceSaleItem sync.Once

//
// GetSaleItemQueryAppService
// @Description: 获取单例应用服务
// @return *SaleItemQueryAppService
//
func GetSaleItemQueryAppService() *SaleItemQueryAppService {
	onceSaleItem.Do(func() {
		saleItemQueryAppService = newSaleItemQueryAppService()
	})
	return saleItemQueryAppService
}

//
// NewSaleItemQueryAppService
// @Description: 创建SaleItem查询应用服务
// @return *SaleItemQueryAppService
//
func newSaleItemQueryAppService() *SaleItemQueryAppService {
	return &SaleItemQueryAppService{
		saleItemCreateExecutor:     sale_item_impl.GetSaleItemCreateExecutor(),
		saleItemCreateManyExecutor: sale_item_impl.GetSaleItemCreateManyExecutor(),

		saleItemUpdateExecutor:     sale_item_impl.GetSaleItemUpdateExecutor(),
		saleItemUpdateManyExecutor: sale_item_impl.GetSaleItemUpdateManyExecutor(),

		saleItemDeleteByIdExecutor:         sale_item_impl.GetSaleItemDeleteByIdExecutor(),
		saleItemDeleteManyExecutor:         sale_item_impl.GetSaleItemDeleteManyExecutor(),
		saleItemDeleteAllExecutor:          sale_item_impl.GetSaleItemDeleteAllExecutor(),
		saleItemDeleteBySaleBillIdExecutor: sale_item_impl.GetSaleItemDeleteBySaleBillIdExecutor(),

		saleItemFindAllExecutor:          sale_item_impl.GetSaleItemFindAllExecutor(),
		saleItemFindByIdExecutor:         sale_item_impl.GetSaleItemFindByIdExecutor(),
		saleItemFindByIdsExecutor:        sale_item_impl.GetSaleItemFindByIdsExecutor(),
		saleItemFindPagingExecutor:       sale_item_impl.GetSaleItemFindPagingExecutor(),
		saleItemFindBySaleBillIdExecutor: sale_item_impl.GetSaleItemFindBySaleBillIdExecutor(),
	}
}

//
// Create
// @Description: 创建SaleItemView
// @param ctx 上下文
// @param *view.SaleItemView SaleItem实体对象
// @return error 错误
//
func (a *SaleItemQueryAppService) Create(ctx context.Context, v *view.SaleItemView) error {
	return a.saleItemCreateExecutor.Execute(ctx, v)
}

//
// CreateMany
// @Description: 创建SaleItemView
// @param ctx
// @return []*view.SaleItemView  SaleItem实体对象切片
// @return error 错误
//
func (a *SaleItemQueryAppService) CreateMany(ctx context.Context, vList []*view.SaleItemView) error {
	return a.saleItemCreateManyExecutor.Execute(ctx, vList)
}

//
// Update
// @Description: 按ID更新SaleItemView
// @receiver a
// @param ctx
// @param v  *view.SaleItemView
// @return error 错误
//
func (a *SaleItemQueryAppService) Update(ctx context.Context, v *view.SaleItemView) error {
	return a.saleItemUpdateExecutor.Execute(ctx, v)
}

//
// UpdateMany
// @Description:  创建SaleItemView
// @param ctx
// @return []*view.SaleItemView  SaleItem实体对象切片
// @return error 错误
//
func (a *SaleItemQueryAppService) UpdateMany(ctx context.Context, vList []*view.SaleItemView) error {
	return a.saleItemUpdateManyExecutor.Execute(ctx, vList)
}

//
// DeleteById
// @Description: 按ID删除SaleItemView
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return error 错误
//
func (a *SaleItemQueryAppService) DeleteById(ctx context.Context, tenantId, id string) error {
	return a.saleItemDeleteByIdExecutor.Execute(ctx, tenantId, id)
}

//
// DeleteMany
// @Description: 删除多个SaleItemView
// @param ctx
// @param tenantId 租户ID
// @param []*view.SaleItemView  SaleItem实体对象切片
// @return error 错误
//
func (a *SaleItemQueryAppService) DeleteMany(ctx context.Context, tenantId string, vList []*view.SaleItemView) error {
	return a.saleItemDeleteManyExecutor.Execute(ctx, tenantId, vList)
}

//
// DeleteBySaleBillId
// @Description: 删除多个SaleItemView
// @param ctx
// @param tenantId 租户ID
// @param []*view.SaleItemView  SaleItem实体对象切片
// @return error 错误
//
func (a *SaleItemQueryAppService) DeleteBySaleBillId(ctx context.Context, tenantId string, saleBillId string) error {
	return a.saleItemDeleteBySaleBillIdExecutor.Execute(ctx, tenantId, saleBillId)
}

//
// DeleteAll
// @Description:  删除所有
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return error
//
func (a *SaleItemQueryAppService) DeleteAll(ctx context.Context, tenantId string) error {
	return a.saleItemDeleteAllExecutor.Execute(ctx, tenantId)

}

//
// FindById
// @Description:  按ID查询SaleItemView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.SaleItemView
// @return bool 是否查询到数据
// @return error
//
func (a *SaleItemQueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.SaleItemView, bool, error) {
	qry := assembler.AssSaleItemFindByIdQuery(tenantId, id)
	return a.saleItemFindByIdExecutor.Execute(ctx, qry)
}

//
// FindByIds
// @Description:  按多个ID查询SaleItemView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.SaleItemView
// @return bool 是否查询到数据
// @return error
//
func (a *SaleItemQueryAppService) FindByIds(ctx context.Context, tenantId string, ids []string) ([]*view.SaleItemView, bool, error) {
	qry := assembler.AssSaleItemFindByIdsQuery(tenantId, ids)
	return a.saleItemFindByIdsExecutor.Execute(ctx, qry)
}

//
// FindAll
// @Description: 查询所有view.SaleItemView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return []*view.SaleItemView
// @return bool 是否查询到数据
// @return error 错误
//
func (a *SaleItemQueryAppService) FindAll(ctx context.Context, tenantId string) ([]*view.SaleItemView, bool, error) {
	qry := assembler.AssSaleItemFindAllQuery(tenantId)
	return a.saleItemFindAllExecutor.Execute(ctx, qry)
}

//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param qry 分页查询条件
// @return *appquery.SaleItemFindPagingResult 分页数据
// @return bool 是否查询到数据
// @return error 错误
//
func (a *SaleItemQueryAppService) FindPaging(ctx context.Context, aq *appquery.SaleItemFindPagingAppQuery) (*appquery.SaleItemFindPagingResult, bool, error) {
	return a.saleItemFindPagingExecutor.Execute(ctx, aq)
}

//
// FindBySaleBillId
// @Description: 根据SaleBillId查询
// @receiver a
// @param ctx 上下文
// @param tenantId string 租户ID
// @param saleBillId  SaleBillID
// @return []*view.SaleItemView 实体切片
// @return bool 是否查询到数据
// @return error 错误
//
func (a *SaleItemQueryAppService) FindBySaleBillId(ctx context.Context, tenantId string, saleBillId string) ([]*view.SaleItemView, bool, error) {
	aq := assembler.AssSaleItemFindBySaleBillIdQuery(tenantId, saleBillId)
	return a.saleItemFindBySaleBillIdExecutor.Execute(ctx, aq)
}
