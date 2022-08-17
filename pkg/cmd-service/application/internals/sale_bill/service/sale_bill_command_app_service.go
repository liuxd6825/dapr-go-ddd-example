package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/executor"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/model"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/base/application/service"
	query_dto "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/sale_bill/dto"
)

type SaleBillCommandAppService struct {
	base.BaseQueryAppService
	saleBillConfirmCommandExecutor executor.SaleBillConfirmCommandExecutor
	saleBillCreateCommandExecutor  executor.SaleBillCreateCommandExecutor
	saleBillDeleteCommandExecutor  executor.SaleBillDeleteCommandExecutor
	saleBillUpdateCommandExecutor  executor.SaleBillUpdateCommandExecutor
	saleItemCreateCommandExecutor  executor.SaleItemCreateCommandExecutor
	saleItemDeleteCommandExecutor  executor.SaleItemDeleteCommandExecutor
	saleItemUpdateCommandExecutor  executor.SaleItemUpdateCommandExecutor
	findAggregateByIdExecutor      executor.FindAggregateByIdExecutor
}

//
// NewSaleBillCommandAppService
// @Description:  <no value>
// @return *SaleBillCommandAppService
//
func NewSaleBillCommandAppService() *SaleBillCommandAppService {
	res := &SaleBillCommandAppService{
		saleBillConfirmCommandExecutor: executor.GetSaleBillConfirmCommandExecutor(),
		saleBillCreateCommandExecutor:  executor.GetSaleBillCreateCommandExecutor(),
		saleBillDeleteCommandExecutor:  executor.GetSaleBillDeleteCommandExecutor(),
		saleBillUpdateCommandExecutor:  executor.GetSaleBillUpdateCommandExecutor(),
		saleItemCreateCommandExecutor:  executor.GetSaleItemCreateCommandExecutor(),
		saleItemDeleteCommandExecutor:  executor.GetSaleItemDeleteCommandExecutor(),
		saleItemUpdateCommandExecutor:  executor.GetSaleItemUpdateCommandExecutor(),
		findAggregateByIdExecutor:      executor.GetFindAggregateByIdExecutor(),
	}
	res.Init("dapr-ddd-demo-query-service", "sale-bills", "v1.0")
	return res
}

//
// SaleBillConfirm
// @Description: 下单确认命令
// @receiver s
// @param ctx 上下文
// @param cmd 下单确认命令命令DTO对象
// @return error
//
func (s *SaleBillCommandAppService) SaleBillConfirm(ctx context.Context, appCmd *appcmd.SaleBillConfirmAppCmd) error {
	return s.saleBillConfirmCommandExecutor.Execute(ctx, appCmd)
}

//
// SaleBillCreate
// @Description: 创建销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 创建销售订单命令DTO对象
// @return error
//
func (s *SaleBillCommandAppService) SaleBillCreate(ctx context.Context, appCmd *appcmd.SaleBillCreateAppCmd) error {
	return s.saleBillCreateCommandExecutor.Execute(ctx, appCmd)
}

//
// SaleBillDelete
// @Description: 删除销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 删除销售订单命令DTO对象
// @return error
//
func (s *SaleBillCommandAppService) SaleBillDelete(ctx context.Context, appCmd *appcmd.SaleBillDeleteAppCmd) error {
	return s.saleBillDeleteCommandExecutor.Execute(ctx, appCmd)
}

//
// SaleBillUpdate
// @Description: 更新销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 更新销售订单命令DTO对象
// @return error
//
func (s *SaleBillCommandAppService) SaleBillUpdate(ctx context.Context, appCmd *appcmd.SaleBillUpdateAppCmd) error {
	return s.saleBillUpdateCommandExecutor.Execute(ctx, appCmd)
}

//
// SaleItemCreate
// @Description: 创建扫描文件
// @receiver s
// @param ctx 上下文
// @param cmd 创建扫描文件命令DTO对象
// @return error
//
func (s *SaleBillCommandAppService) SaleItemCreate(ctx context.Context, appCmd *appcmd.SaleItemCreateAppCmd) error {
	return s.saleItemCreateCommandExecutor.Execute(ctx, appCmd)
}

//
// SaleItemDelete
// @Description: 删除扫描单
// @receiver s
// @param ctx 上下文
// @param cmd 删除扫描单命令DTO对象
// @return error
//
func (s *SaleBillCommandAppService) SaleItemDelete(ctx context.Context, appCmd *appcmd.SaleItemDeleteAppCmd) error {
	return s.saleItemDeleteCommandExecutor.Execute(ctx, appCmd)
}

//
// SaleItemUpdate
// @Description: 更新扫描文件
// @receiver s
// @param ctx 上下文
// @param cmd 更新扫描文件命令DTO对象
// @return error
//
func (s *SaleBillCommandAppService) SaleItemUpdate(ctx context.Context, appCmd *appcmd.SaleItemUpdateAppCmd) error {
	return s.saleItemUpdateCommandExecutor.Execute(ctx, appCmd)
}

//
// FindAggregateById
// @Description:
// @receiver s
// @param ctx 上下文
// @param tenantId 租户Id
// @param id 聚合根Id
// @return error
//
func (s *SaleBillCommandAppService) FindAggregateById(ctx context.Context, tenantId string, id string) (*model.SaleBillAggregate, bool, error) {
	return s.findAggregateByIdExecutor.Execute(ctx, tenantId, id)
}

//
// QueryById
// @Description: 按id获取<no value>投影类
// @receiver s queryAppService
// @param ctx 上下文
// @param tenantId  租户id
// @param id <no value> Id
// @return data <no value> 信息
// @return isFound 是否找到
// @return err 错误信息
//
func (s *SaleBillCommandAppService) QueryById(ctx context.Context, tenantId string, id string) (*query_dto.SaleBillFindByIdResponse, bool, error) {
	var resp query_dto.SaleBillFindByIdResponse
	isFound, err := s.BaseQueryAppService.QueryById(ctx, tenantId, id, &resp)
	if err != nil {
		return nil, false, err
	}
	return &resp, isFound, nil
}
