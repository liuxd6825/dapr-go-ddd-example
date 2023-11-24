package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/model"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/utils/singleutils"
)

// SaleBillCommandDomainService
// @Description:  <no value> 命令领域服务
type SaleBillCommandDomainService struct {
	ddd_service.CommandHelper[*model.SaleBillAggregate]
}

// NewSaleBillCommandDomainService
// @Description: 创建领域服务
// @return *SaleBillCommandDomainService
func NewSaleBillCommandDomainService() *SaleBillCommandDomainService {
	return singleutils.Create[*SaleBillCommandDomainService]("cmd-service.domain.user.SaleBillCommandDomainService", func() *SaleBillCommandDomainService {
		return &SaleBillCommandDomainService{}
	})
}

// SaleBillConfirm
// @Description: 下单确认命令
// @receiver s
// @param ctx 上下文
// @param cmd 下单确认命令
// @return *model.SaleBillCommandDomainService
// @return error
func (s *SaleBillCommandDomainService) SaleBillConfirm(ctx context.Context, cmd *command.SaleBillConfirmCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// SaleBillCreate
// @Description: 创建销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 创建销售订单
// @return *model.SaleBillCommandDomainService
// @return error
func (s *SaleBillCommandDomainService) SaleBillCreate(ctx context.Context, cmd *command.SaleBillCreateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// SaleBillDelete
// @Description: 删除销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 删除销售订单
// @return *model.SaleBillCommandDomainService
// @return error
func (s *SaleBillCommandDomainService) SaleBillDelete(ctx context.Context, cmd *command.SaleBillDeleteCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// SaleBillUpdate
// @Description: 更新销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 更新销售订单
// @return *model.SaleBillCommandDomainService
// @return error
func (s *SaleBillCommandDomainService) SaleBillUpdate(ctx context.Context, cmd *command.SaleBillUpdateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// SaleItemCreate
// @Description: 添加明细
// @receiver s
// @param ctx 上下文
// @param cmd 添加明细
// @return *model.SaleBillCommandDomainService
// @return error
func (s *SaleBillCommandDomainService) SaleItemCreate(ctx context.Context, cmd *command.SaleItemCreateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// SaleItemDelete
// @Description: 删除销售明细项
// @receiver s
// @param ctx 上下文
// @param cmd 删除销售明细项
// @return *model.SaleBillCommandDomainService
// @return error
func (s *SaleBillCommandDomainService) SaleItemDelete(ctx context.Context, cmd *command.SaleItemDeleteCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.DoCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

// SaleItemUpdate
// @Description: 更新明细
// @receiver s
// @param ctx 上下文
// @param cmd 更新明细
// @return *model.SaleBillCommandDomainService
// @return error
func (s *SaleBillCommandDomainService) SaleItemUpdate(ctx context.Context, cmd *command.SaleItemUpdateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
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
// @return *sale_bill_model.SaleBillCommandDomainService  聚合对象
// @return bool 是否找到聚合根对象
// @return error 错误对象
func (s *SaleBillCommandDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.SaleBillAggregate, bool, error) {
	agg := s.NewAggregate()
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, agg)
	return agg, ok, err
}

// NewAggregate
// @Description: 新建聚合对象
// @receiver s
// @return *sale_bill_model.SaleBillCommandDomainService 聚合对象
func (s *SaleBillCommandDomainService) NewAggregate() *model.SaleBillAggregate {
	return model.NewSaleBillAggregate()
}
