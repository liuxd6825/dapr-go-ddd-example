package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/executor"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/executor/inventory_impl"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	"sync"
)

//
// InventoryQueryAppService
// @Description: <no value>查询应用服务类
//
type InventoryQueryAppService struct {
	inventoryCreateExecutor     executor.InventoryCreateExecutor
	inventoryCreateManyExecutor executor.InventoryCreateManyExecutor

	inventoryUpdateExecutor     executor.InventoryUpdateExecutor
	inventoryUpdateManyExecutor executor.InventoryUpdateManyExecutor

	inventoryDeleteByIdExecutor executor.InventoryDeleteByIdExecutor
	inventoryDeleteManyExecutor executor.InventoryDeleteManyExecutor
	inventoryDeleteAllExecutor  executor.InventoryDeleteAllExecutor

	inventoryFindAllExecutor    executor.InventoryFindAllExecutor
	inventoryFindByIdExecutor   executor.InventoryFindByIdExecutor
	inventoryFindByIdsExecutor  executor.InventoryFindByIdsExecutor
	inventoryFindPagingExecutor executor.InventoryFindPagingExecutor
}

// 单例应用服务
var inventoryQueryAppService *InventoryQueryAppService

// 并发安全
var onceInventory sync.Once

//
// GetInventoryQueryAppService
// @Description: 获取单例应用服务
// @return *InventoryQueryAppService
//
func GetInventoryQueryAppService() *InventoryQueryAppService {
	onceInventory.Do(func() {
		inventoryQueryAppService = newInventoryQueryAppService()
	})
	return inventoryQueryAppService
}

//
// NewInventoryQueryAppService
// @Description: 创建Inventory查询应用服务
// @return *InventoryQueryAppService
//
func newInventoryQueryAppService() *InventoryQueryAppService {
	return &InventoryQueryAppService{
		inventoryCreateExecutor:     inventory_impl.GetInventoryCreateExecutor(),
		inventoryCreateManyExecutor: inventory_impl.GetInventoryCreateManyExecutor(),

		inventoryUpdateExecutor:     inventory_impl.GetInventoryUpdateExecutor(),
		inventoryUpdateManyExecutor: inventory_impl.GetInventoryUpdateManyExecutor(),

		inventoryDeleteByIdExecutor: inventory_impl.GetInventoryDeleteByIdExecutor(),
		inventoryDeleteManyExecutor: inventory_impl.GetInventoryDeleteManyExecutor(),
		inventoryDeleteAllExecutor:  inventory_impl.GetInventoryDeleteAllExecutor(),

		inventoryFindAllExecutor:    inventory_impl.GetInventoryFindAllExecutor(),
		inventoryFindByIdExecutor:   inventory_impl.GetInventoryFindByIdExecutor(),
		inventoryFindByIdsExecutor:  inventory_impl.GetInventoryFindByIdsExecutor(),
		inventoryFindPagingExecutor: inventory_impl.GetInventoryFindPagingExecutor(),
	}
}

//
// Create
// @Description: 创建InventoryView
// @param ctx 上下文
// @param *view.InventoryView Inventory实体对象
// @return error 错误
//
func (a *InventoryQueryAppService) Create(ctx context.Context, v *view.InventoryView) error {
	return a.inventoryCreateExecutor.Execute(ctx, v)
}

//
// CreateMany
// @Description: 创建InventoryView
// @param ctx
// @return []*view.InventoryView  Inventory实体对象切片
// @return error 错误
//
func (a *InventoryQueryAppService) CreateMany(ctx context.Context, vList []*view.InventoryView) error {
	return a.inventoryCreateManyExecutor.Execute(ctx, vList)
}

//
// Update
// @Description: 按ID更新InventoryView
// @receiver a
// @param ctx
// @param v  *view.InventoryView
// @return error 错误
//
func (a *InventoryQueryAppService) Update(ctx context.Context, v *view.InventoryView) error {
	return a.inventoryUpdateExecutor.Execute(ctx, v)
}

//
// UpdateMany
// @Description:  创建InventoryView
// @param ctx
// @return []*view.InventoryView  Inventory实体对象切片
// @return error 错误
//
func (a *InventoryQueryAppService) UpdateMany(ctx context.Context, vList []*view.InventoryView) error {
	return a.inventoryUpdateManyExecutor.Execute(ctx, vList)
}

//
// DeleteById
// @Description: 按ID删除InventoryView
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return error 错误
//
func (a *InventoryQueryAppService) DeleteById(ctx context.Context, tenantId, id string) error {
	return a.inventoryDeleteByIdExecutor.Execute(ctx, tenantId, id)
}

//
// DeleteMany
// @Description: 删除多个InventoryView
// @param ctx
// @param tenantId 租户ID
// @param []*view.InventoryView  Inventory实体对象切片
// @return error 错误
//
func (a *InventoryQueryAppService) DeleteMany(ctx context.Context, tenantId string, vList []*view.InventoryView) error {
	return a.inventoryDeleteManyExecutor.Execute(ctx, tenantId, vList)
}

//
// DeleteAll
// @Description:  删除所有
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return error
//
func (a *InventoryQueryAppService) DeleteAll(ctx context.Context, tenantId string) error {
	return a.inventoryDeleteAllExecutor.Execute(ctx, tenantId)

}

//
// FindById
// @Description:  按ID查询InventoryView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.InventoryView
// @return bool 是否查询到数据
// @return error
//
func (a *InventoryQueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.InventoryView, bool, error) {
	qry := assembler.AssInventoryFindByIdQuery(tenantId, id)
	return a.inventoryFindByIdExecutor.Execute(ctx, qry)
}

//
// FindByIds
// @Description:  按多个ID查询InventoryView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.InventoryView
// @return bool 是否查询到数据
// @return error
//
func (a *InventoryQueryAppService) FindByIds(ctx context.Context, tenantId string, ids []string) ([]*view.InventoryView, bool, error) {
	qry := assembler.AssInventoryFindByIdsQuery(tenantId, ids)
	return a.inventoryFindByIdsExecutor.Execute(ctx, qry)
}

//
// FindAll
// @Description: 查询所有view.InventoryView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return []*view.InventoryView
// @return bool 是否查询到数据
// @return error 错误
//
func (a *InventoryQueryAppService) FindAll(ctx context.Context, tenantId string) ([]*view.InventoryView, bool, error) {
	qry := assembler.AssInventoryFindAllQuery(tenantId)
	return a.inventoryFindAllExecutor.Execute(ctx, qry)
}

//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param qry 分页查询条件
// @return *appquery.InventoryFindPagingResult 分页数据
// @return bool 是否查询到数据
// @return error 错误
//
func (a *InventoryQueryAppService) FindPaging(ctx context.Context, aq *appquery.InventoryFindPagingAppQuery) (*appquery.InventoryFindPagingResult, bool, error) {
	return a.inventoryFindPagingExecutor.Execute(ctx, aq)
}
