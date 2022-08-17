package executor

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
)

//
// SaleBillCreateExecutor
// @Description: 新建
//
type SaleBillCreateExecutor interface {
	Execute(context.Context, *view.SaleBillView) error
}

//
// SaleBillCreateManyExecutor
// @Description: 新建多个
//
type SaleBillCreateManyExecutor interface {
	Execute(context.Context, []*view.SaleBillView) error
}

//
// SaleBillDeleteByIdExecutor
// @Description: 按Id删除
//
type SaleBillDeleteByIdExecutor interface {
	Execute(ctx context.Context, tenantId string, id string) error
}

//
// SaleBillDeleteManyExecutor
// @Description: 删除多个
//
type SaleBillDeleteManyExecutor interface {
	Execute(context.Context, string, []*view.SaleBillView) error
}

//
// SaleBillDeleteAllExecutor
// @Description: 删除所有
//
type SaleBillDeleteAllExecutor interface {
	Execute(ctx context.Context, tenantId string) error
}

//
// SaleBillUpdateExecutor
// @Description: 更新
//
type SaleBillUpdateExecutor interface {
	Execute(context.Context, *view.SaleBillView) error
}

//
// SaleBillUpdateManyExecutor
// @Description: 更新多个
//
type SaleBillUpdateManyExecutor interface {
	Execute(context.Context, []*view.SaleBillView) error
}

//
// SaleBillFindAllExecutor
// @Description: 查询所有
//
type SaleBillFindAllExecutor interface {
	Execute(context.Context, *appquery.SaleBillFindAllAppQuery) ([]*view.SaleBillView, bool, error)
}

//
// SaleBillFindByIdExecutor
// @Description: 按Id查询
//
type SaleBillFindByIdExecutor interface {
	Execute(context.Context, *appquery.SaleBillFindByIdAppQuery) (*view.SaleBillView, bool, error)
}

//
// SaleBillFindByIdsExecutor
// @Description: 按Id列表查询多个
//
type SaleBillFindByIdsExecutor interface {
	Execute(context.Context, *appquery.SaleBillFindByIdsAppQuery) ([]*view.SaleBillView, bool, error)
}

//
// SaleBillFindPagingExecutor
// @Description: 分页查询
//
type SaleBillFindPagingExecutor interface {
	Execute(context.Context, *appquery.SaleBillFindPagingAppQuery) (*appquery.SaleBillFindPagingResult, bool, error)
}
