package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/appquery"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/assembler"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/executor"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/application/internals/sale_bill/executor/sale_bill_impl"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"sync"
)

//
// SaleBillQueryAppService
// @Description: <no value>查询应用服务类
//
type SaleBillQueryAppService struct {
	saleBillCreateExecutor     executor.SaleBillCreateExecutor
	saleBillCreateManyExecutor executor.SaleBillCreateManyExecutor

	saleBillUpdateExecutor     executor.SaleBillUpdateExecutor
	saleBillUpdateManyExecutor executor.SaleBillUpdateManyExecutor

	saleBillDeleteByIdExecutor executor.SaleBillDeleteByIdExecutor
	saleBillDeleteManyExecutor executor.SaleBillDeleteManyExecutor
	saleBillDeleteAllExecutor  executor.SaleBillDeleteAllExecutor

	saleBillFindAllExecutor    executor.SaleBillFindAllExecutor
	saleBillFindByIdExecutor   executor.SaleBillFindByIdExecutor
	saleBillFindByIdsExecutor  executor.SaleBillFindByIdsExecutor
	saleBillFindPagingExecutor executor.SaleBillFindPagingExecutor
}

// 单例应用服务
var saleBillQueryAppService *SaleBillQueryAppService

// 并发安全
var onceSaleBill sync.Once

//
// GetSaleBillQueryAppService
// @Description: 获取单例应用服务
// @return *SaleBillQueryAppService
//
func GetSaleBillQueryAppService() *SaleBillQueryAppService {
	onceSaleBill.Do(func() {
		saleBillQueryAppService = newSaleBillQueryAppService()
	})
	return saleBillQueryAppService
}

//
// NewSaleBillQueryAppService
// @Description: 创建SaleBill查询应用服务
// @return *SaleBillQueryAppService
//
func newSaleBillQueryAppService() *SaleBillQueryAppService {
	return &SaleBillQueryAppService{
		saleBillCreateExecutor:     sale_bill_impl.GetSaleBillCreateExecutor(),
		saleBillCreateManyExecutor: sale_bill_impl.GetSaleBillCreateManyExecutor(),

		saleBillUpdateExecutor:     sale_bill_impl.GetSaleBillUpdateExecutor(),
		saleBillUpdateManyExecutor: sale_bill_impl.GetSaleBillUpdateManyExecutor(),

		saleBillDeleteByIdExecutor: sale_bill_impl.GetSaleBillDeleteByIdExecutor(),
		saleBillDeleteManyExecutor: sale_bill_impl.GetSaleBillDeleteManyExecutor(),
		saleBillDeleteAllExecutor:  sale_bill_impl.GetSaleBillDeleteAllExecutor(),

		saleBillFindAllExecutor:    sale_bill_impl.GetSaleBillFindAllExecutor(),
		saleBillFindByIdExecutor:   sale_bill_impl.GetSaleBillFindByIdExecutor(),
		saleBillFindByIdsExecutor:  sale_bill_impl.GetSaleBillFindByIdsExecutor(),
		saleBillFindPagingExecutor: sale_bill_impl.GetSaleBillFindPagingExecutor(),
	}
}

//
// Create
// @Description: 创建SaleBillView
// @param ctx 上下文
// @param *view.SaleBillView SaleBill实体对象
// @return error 错误
//
func (a *SaleBillQueryAppService) Create(ctx context.Context, v *view.SaleBillView) error {
	return a.saleBillCreateExecutor.Execute(ctx, v)
}

//
// CreateMany
// @Description: 创建SaleBillView
// @param ctx
// @return []*view.SaleBillView  SaleBill实体对象切片
// @return error 错误
//
func (a *SaleBillQueryAppService) CreateMany(ctx context.Context, vList []*view.SaleBillView) error {
	return a.saleBillCreateManyExecutor.Execute(ctx, vList)
}

//
// Update
// @Description: 按ID更新SaleBillView
// @receiver a
// @param ctx
// @param v  *view.SaleBillView
// @return error 错误
//
func (a *SaleBillQueryAppService) Update(ctx context.Context, v *view.SaleBillView) error {
	return a.saleBillUpdateExecutor.Execute(ctx, v)
}

//
// UpdateMany
// @Description:  创建SaleBillView
// @param ctx
// @return []*view.SaleBillView  SaleBill实体对象切片
// @return error 错误
//
func (a *SaleBillQueryAppService) UpdateMany(ctx context.Context, vList []*view.SaleBillView) error {
	return a.saleBillUpdateManyExecutor.Execute(ctx, vList)
}

//
// DeleteById
// @Description: 按ID删除SaleBillView
// @param ctx
// @param tenantId 租户ID
// @param id 视图ID
// @return error 错误
//
func (a *SaleBillQueryAppService) DeleteById(ctx context.Context, tenantId, id string) error {
	return a.saleBillDeleteByIdExecutor.Execute(ctx, tenantId, id)
}

//
// DeleteMany
// @Description: 删除多个SaleBillView
// @param ctx
// @param tenantId 租户ID
// @param []*view.SaleBillView  SaleBill实体对象切片
// @return error 错误
//
func (a *SaleBillQueryAppService) DeleteMany(ctx context.Context, tenantId string, vList []*view.SaleBillView) error {
	return a.saleBillDeleteManyExecutor.Execute(ctx, tenantId, vList)
}

//
// DeleteAll
// @Description:  删除所有
// @receiver a
// @param ctx
// @param tenantId 租户ID
// @return error
//
func (a *SaleBillQueryAppService) DeleteAll(ctx context.Context, tenantId string) error {
	return a.saleBillDeleteAllExecutor.Execute(ctx, tenantId)

}

//
// FindById
// @Description:  按ID查询SaleBillView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.SaleBillView
// @return bool 是否查询到数据
// @return error
//
func (a *SaleBillQueryAppService) FindById(ctx context.Context, tenantId string, id string) (*view.SaleBillView, bool, error) {
	qry := assembler.SaleBill.AssFindByIdAppQuery(tenantId, id)
	return a.saleBillFindByIdExecutor.Execute(ctx, qry)
}

//
// FindByIds
// @Description:  按多个ID查询SaleBillView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return *view.SaleBillView
// @return bool 是否查询到数据
// @return error
//
func (a *SaleBillQueryAppService) FindByIds(ctx context.Context, tenantId string, ids []string) ([]*view.SaleBillView, bool, error) {
	qry := assembler.SaleBill.AssFindByIdsAppQuery(tenantId, ids)
	return a.saleBillFindByIdsExecutor.Execute(ctx, qry)
}

//
// FindAll
// @Description: 查询所有view.SaleBillView
// @receiver a
// @param ctx
// @param qry 查询命令
// @return []*view.SaleBillView
// @return bool 是否查询到数据
// @return error 错误
//
func (a *SaleBillQueryAppService) FindAll(ctx context.Context, tenantId string) ([]*view.SaleBillView, bool, error) {
	qry := assembler.SaleBill.AssFindAllAppQuery(tenantId)
	return a.saleBillFindAllExecutor.Execute(ctx, qry)
}

//
// FindPaging
// @Description: 分页查询
// @receiver a
// @param ctx 上下文
// @param qry 分页查询条件
// @return *appquery.SaleBillFindPagingResult 分页数据
// @return bool 是否查询到数据
// @return error 错误
//
func (a *SaleBillQueryAppService) FindPaging(ctx context.Context, aq *appquery.SaleBillFindPagingAppQuery) (*appquery.SaleBillFindPagingResult, bool, error) {
	return a.saleBillFindPagingExecutor.Execute(ctx, aq)
}
