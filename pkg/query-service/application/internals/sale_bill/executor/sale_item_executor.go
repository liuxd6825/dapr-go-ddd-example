package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
)

//
// SaleItemCreateExecutor
// @Description: 新建
//
type SaleItemCreateExecutor interface {
	Execute(context.Context, *view.SaleItemView) error
}

//
// SaleItemCreateManyExecutor
// @Description: 新建多个
//
type SaleItemCreateManyExecutor interface {
	Execute(context.Context, []*view.SaleItemView) error
}

//
// SaleItemDeleteByIdExecutor
// @Description: 按Id删除
//
type SaleItemDeleteByIdExecutor interface {
	Execute(ctx context.Context, tenantId string, id string) error
}

//
// SaleItemDeleteBySaleBillIdExecutor
// @Description: 按聚合根Id删除
//
type SaleItemDeleteBySaleBillIdExecutor interface {
	Execute(ctx context.Context, tenantId string, saleBillId string) error
}

//
// SaleItemDeleteManyExecutor
// @Description: 删除多个
//
type SaleItemDeleteManyExecutor interface {
	Execute(context.Context, string, []*view.SaleItemView) error
}

//
// SaleItemDeleteAllExecutor
// @Description: 删除所有
//
type SaleItemDeleteAllExecutor interface {
	Execute(ctx context.Context, tenantId string) error
}

//
// SaleItemUpdateExecutor
// @Description: 更新
//
type SaleItemUpdateExecutor interface {
	Execute(context.Context, *view.SaleItemView) error
}

//
// SaleItemUpdateManyExecutor
// @Description: 更新多个
//
type SaleItemUpdateManyExecutor interface {
	Execute(context.Context, []*view.SaleItemView) error
}

//
// SaleItemFindAllExecutor
// @Description: 查询所有
//
type SaleItemFindAllExecutor interface {
	Execute(context.Context, *appquery.SaleItemFindAllAppQuery) ([]*view.SaleItemView, bool, error)
}

//
// SaleItemFindByIdExecutor
// @Description: 按Id查询
//
type SaleItemFindByIdExecutor interface {
	Execute(context.Context, *appquery.SaleItemFindByIdAppQuery) (*view.SaleItemView, bool, error)
}

//
// SaleItemFindByIdsExecutor
// @Description: 按Id列表查询多个
//
type SaleItemFindByIdsExecutor interface {
	Execute(context.Context, *appquery.SaleItemFindByIdsAppQuery) ([]*view.SaleItemView, bool, error)
}

//
// SaleItemFindPagingExecutor
// @Description: 分页查询
//
type SaleItemFindPagingExecutor interface {
	Execute(context.Context, *appquery.SaleItemFindPagingAppQuery) (*appquery.SaleItemFindPagingResult, bool, error)
}

//
// SaleItemFindBySaleBillIdExecutor
// @Description: 按聚合Id查询
//
type SaleItemFindBySaleBillIdExecutor interface {
	Execute(context.Context, *appquery.SaleItemFindBySaleBillIdAppQuery) ([]*view.SaleItemView, bool, error)
}
