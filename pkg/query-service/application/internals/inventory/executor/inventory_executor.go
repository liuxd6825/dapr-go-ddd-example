package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/inventory/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
)

//
// InventoryCreateExecutor
// @Description: 新建
//
type InventoryCreateExecutor interface {
	Execute(context.Context, *view.InventoryView) error
}

//
// InventoryCreateManyExecutor
// @Description: 新建多个
//
type InventoryCreateManyExecutor interface {
	Execute(context.Context, []*view.InventoryView) error
}

//
// InventoryDeleteByIdExecutor
// @Description: 按Id删除
//
type InventoryDeleteByIdExecutor interface {
	Execute(ctx context.Context, tenantId string, id string) error
}

//
// InventoryDeleteManyExecutor
// @Description: 删除多个
//
type InventoryDeleteManyExecutor interface {
	Execute(context.Context, string, []*view.InventoryView) error
}

//
// InventoryDeleteAllExecutor
// @Description: 删除所有
//
type InventoryDeleteAllExecutor interface {
	Execute(ctx context.Context, tenantId string) error
}

//
// InventoryUpdateExecutor
// @Description: 更新
//
type InventoryUpdateExecutor interface {
	Execute(context.Context, *view.InventoryView) error
}

//
// InventoryUpdateManyExecutor
// @Description: 更新多个
//
type InventoryUpdateManyExecutor interface {
	Execute(context.Context, []*view.InventoryView) error
}

//
// InventoryFindAllExecutor
// @Description: 查询所有
//
type InventoryFindAllExecutor interface {
	Execute(context.Context, *appquery.InventoryFindAllAppQuery) ([]*view.InventoryView, bool, error)
}

//
// InventoryFindByIdExecutor
// @Description: 按Id查询
//
type InventoryFindByIdExecutor interface {
	Execute(context.Context, *appquery.InventoryFindByIdAppQuery) (*view.InventoryView, bool, error)
}

//
// InventoryFindByIdsExecutor
// @Description: 按Id列表查询多个
//
type InventoryFindByIdsExecutor interface {
	Execute(context.Context, *appquery.InventoryFindByIdsAppQuery) ([]*view.InventoryView, bool, error)
}

//
// InventoryFindPagingExecutor
// @Description: 分页查询
//
type InventoryFindPagingExecutor interface {
	Execute(context.Context, *appquery.InventoryFindPagingAppQuery) (*appquery.InventoryFindPagingResult, bool, error)
}
