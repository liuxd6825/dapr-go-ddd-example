package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/inventory/executor"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/model"
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/base/application/service"
	query_dto "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/userinterface/rest/inventory/dto"
)

type InventoryCommandAppService struct {
	base.BaseQueryAppService
	inventoryCreateCommandExecutor executor.InventoryCreateCommandExecutor
	inventoryUpdateCommandExecutor executor.InventoryUpdateCommandExecutor
	findAggregateByIdExecutor      executor.FindAggregateByIdExecutor
}

//
// NewInventoryCommandAppService
// @Description:  <no value>
// @return *InventoryCommandAppService
//
func NewInventoryCommandAppService() *InventoryCommandAppService {
	res := &InventoryCommandAppService{
		inventoryCreateCommandExecutor: executor.GetInventoryCreateCommandExecutor(),
		inventoryUpdateCommandExecutor: executor.GetInventoryUpdateCommandExecutor(),
		findAggregateByIdExecutor:      executor.GetFindAggregateByIdExecutor(),
	}
	res.Init("dapr-ddd-demo-query-service", "inventories", "v1.0")
	return res
}

//
// InventoryCreate
// @Description: 创建存货档案
// @receiver s
// @param ctx 上下文
// @param cmd 创建存货档案命令DTO对象
// @return error
//
func (s *InventoryCommandAppService) InventoryCreate(ctx context.Context, appCmd *appcmd.InventoryCreateAppCmd) error {
	return s.inventoryCreateCommandExecutor.Execute(ctx, appCmd)
}

//
// InventoryUpdate
// @Description: 更新存货档案
// @receiver s
// @param ctx 上下文
// @param cmd 更新存货档案命令DTO对象
// @return error
//
func (s *InventoryCommandAppService) InventoryUpdate(ctx context.Context, appCmd *appcmd.InventoryUpdateAppCmd) error {
	return s.inventoryUpdateCommandExecutor.Execute(ctx, appCmd)
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
func (s *InventoryCommandAppService) FindAggregateById(ctx context.Context, tenantId string, id string) (*model.InventoryAggregate, bool, error) {
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
func (s *InventoryCommandAppService) QueryById(ctx context.Context, tenantId string, id string) (*query_dto.InventoryFindByIdResponse, bool, error) {
	var resp query_dto.InventoryFindByIdResponse
	isFound, err := s.BaseQueryAppService.QueryById(ctx, tenantId, id, &resp)
	if err != nil {
		return nil, false, err
	}
	return &resp, isFound, nil
}
