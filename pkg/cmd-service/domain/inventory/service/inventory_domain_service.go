package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

// InventoryCommandDomainService
// @Description:  <no value> 命令领域服务
type InventoryCommandDomainService struct {
	ddd_service.CommandHelper[*model.InventoryAggregate]
}

// NewInventoryCommandDomainService
// @Description: 创建领域服务
// @return *InventoryCommandDomainService
func NewInventoryCommandDomainService() *InventoryCommandDomainService {
	return singleutils.Create[*InventoryCommandDomainService]("InventoryCommandDomainService", func() *InventoryCommandDomainService {
		return &InventoryCommandDomainService{}
	})
}

// InventoryCreate
// @Description: 创建存货档案
// @receiver s
// @param ctx 上下文
// @param cmd 创建存货档案
// @return *model.InventoryCommandDomainService
// @return error
func (s *InventoryCommandDomainService) InventoryCreate(ctx context.Context, cmd *command.InventoryCreateCommand, opts ...ddd.DoCommandOption) (*model.InventoryAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// InventoryUpdate
// @Description: 更新存货档案
// @receiver s
// @param ctx 上下文
// @param cmd 更新存货档案
// @return *model.InventoryCommandDomainService
// @return error
func (s *InventoryCommandDomainService) InventoryUpdate(ctx context.Context, cmd *command.InventoryUpdateCommand, opts ...ddd.DoCommandOption) (*model.InventoryAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// GetAggregateById
// @Description: 获取聚合对象
// @receiver s
// @param ctx 上下文
// @param tenantId 租户id
// @param id 主键id
// @return *inventory_model.InventoryCommandDomainService  聚合对象
// @return bool 是否找到聚合根对象
// @return error 错误对象
func (s *InventoryCommandDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.InventoryAggregate, bool, error) {
	agg, _ := s.NewAggregate()
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, agg)
	return agg, ok, err
}
