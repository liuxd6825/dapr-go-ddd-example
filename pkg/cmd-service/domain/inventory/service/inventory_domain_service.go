package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/model"
	base_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/base/domain/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"sync"
)

//
// InventoryCommandDomainService
// @Description:  <no value> 命令领域服务
//
type InventoryCommandDomainService struct {
	base_service.BaseCommandDomainService
}

// 单例应用服务
var inventoryCommandDomainService *InventoryCommandDomainService

// 并发安全
var onceInventory sync.Once

//
// GetInventoryCommandDomainService
// @Description: 获取单例领域服务
// @return service.InventoryQueryDomainService
//
func GetInventoryCommandDomainService() *InventoryCommandDomainService {
	onceInventory.Do(func() {
		inventoryCommandDomainService = newInventoryCommandDomainService()
	})
	return inventoryCommandDomainService
}

//
// NewInventoryCommandDomainService
// @Description: 创建领域服务
// @return *InventoryCommandDomainService
//
func newInventoryCommandDomainService() *InventoryCommandDomainService {
	return &InventoryCommandDomainService{}
}

//
// InventoryCreate
// @Description: 创建存货档案
// @receiver s
// @param ctx 上下文
// @param cmd 创建存货档案
// @return *model.InventoryCommandDomainService
// @return error
//
func (s *InventoryCommandDomainService) InventoryCreate(ctx context.Context, cmd *command.InventoryCreateCommand, opts ...ddd.DoCommandOption) (*model.InventoryAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// InventoryUpdate
// @Description: 更新存货档案
// @receiver s
// @param ctx 上下文
// @param cmd 更新存货档案
// @return *model.InventoryCommandDomainService
// @return error
//
func (s *InventoryCommandDomainService) InventoryUpdate(ctx context.Context, cmd *command.InventoryUpdateCommand, opts ...ddd.DoCommandOption) (*model.InventoryAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
//  doCommand
//  @Description:
//  @receiver s
//  @param ctx
//  @param cmd
//  @return *model.InventoryAggregate
//  @return error
//
func (s *InventoryCommandDomainService) doCommand(ctx context.Context, cmd ddd.Command, validateFunc func() error, opts ...ddd.DoCommandOption) (*model.InventoryAggregate, error) {
	option := ddd.NewDoCommandOptionMerges(opts...)

	// 进行业务检查
	if validateFunc != nil {
		if err := validateFunc(); err != nil {
			return nil, err
		}
	} else if err := cmd.Validate(); err != nil {
		return nil, err
	}

	// 如果只是业务检查，则不执行领域命令，
	validOnly := option.GetIsValidOnly()
	if (validOnly == nil && cmd.GetIsValidOnly()) || (validOnly != nil && *validOnly == true) {
		return nil, nil
	}

	// 新建聚合根对象
	agg := s.NewAggregate()

	// 如果领域命令执行时出错，则返回错误
	if err := ddd.ApplyCommand(ctx, agg, cmd); err != nil {
		return nil, err
	}

	return agg, nil
}

//
// GetAggregateById
// @Description: 获取聚合对象
// @receiver s
// @param ctx 上下文
// @param tenantId 租户id
// @param id 主键id
// @return *inventory_model.InventoryCommandDomainService  聚合对象
// @return bool 是否找到聚合根对象
// @return error 错误对象
//
func (s *InventoryCommandDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.InventoryAggregate, bool, error) {
	agg := s.NewAggregate()
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, agg)
	return agg, ok, err
}

//
// NewAggregate
// @Description: 新建聚合对象
// @receiver s
// @return *inventory_model.InventoryCommandDomainService 聚合对象
//
func (s *InventoryCommandDomainService) NewAggregate() *model.InventoryAggregate {
	return model.NewInventoryAggregate()
}
