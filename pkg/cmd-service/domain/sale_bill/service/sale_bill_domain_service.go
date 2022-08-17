package service

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/model"
	base_service "gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/base/domain/service"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"sync"
)

//
// SaleBillCommandDomainService
// @Description:  <no value> 命令领域服务
//
type SaleBillCommandDomainService struct {
	base_service.BaseCommandDomainService
}

// 单例应用服务
var saleBillCommandDomainService *SaleBillCommandDomainService

// 并发安全
var onceSaleBill sync.Once

//
// GetSaleBillCommandDomainService
// @Description: 获取单例领域服务
// @return service.SaleBillQueryDomainService
//
func GetSaleBillCommandDomainService() *SaleBillCommandDomainService {
	onceSaleBill.Do(func() {
		saleBillCommandDomainService = newSaleBillCommandDomainService()
	})
	return saleBillCommandDomainService
}

//
// NewSaleBillCommandDomainService
// @Description: 创建领域服务
// @return *SaleBillCommandDomainService
//
func newSaleBillCommandDomainService() *SaleBillCommandDomainService {
	return &SaleBillCommandDomainService{}
}

//
// SaleBillConfirm
// @Description: 下单确认命令
// @receiver s
// @param ctx 上下文
// @param cmd 下单确认命令
// @return *model.SaleBillCommandDomainService
// @return error
//
func (s *SaleBillCommandDomainService) SaleBillConfirm(ctx context.Context, cmd *command.SaleBillConfirmCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// SaleBillCreate
// @Description: 创建销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 创建销售订单
// @return *model.SaleBillCommandDomainService
// @return error
//
func (s *SaleBillCommandDomainService) SaleBillCreate(ctx context.Context, cmd *command.SaleBillCreateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// SaleBillDelete
// @Description: 删除销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 删除销售订单
// @return *model.SaleBillCommandDomainService
// @return error
//
func (s *SaleBillCommandDomainService) SaleBillDelete(ctx context.Context, cmd *command.SaleBillDeleteCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// SaleBillUpdate
// @Description: 更新销售订单
// @receiver s
// @param ctx 上下文
// @param cmd 更新销售订单
// @return *model.SaleBillCommandDomainService
// @return error
//
func (s *SaleBillCommandDomainService) SaleBillUpdate(ctx context.Context, cmd *command.SaleBillUpdateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// SaleItemCreate
// @Description: 创建扫描文件
// @receiver s
// @param ctx 上下文
// @param cmd 创建扫描文件
// @return *model.SaleBillCommandDomainService
// @return error
//
func (s *SaleBillCommandDomainService) SaleItemCreate(ctx context.Context, cmd *command.SaleItemCreateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// SaleItemDelete
// @Description: 删除扫描单
// @receiver s
// @param ctx 上下文
// @param cmd 删除扫描单
// @return *model.SaleBillCommandDomainService
// @return error
//
func (s *SaleBillCommandDomainService) SaleItemDelete(ctx context.Context, cmd *command.SaleItemDeleteCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
	return s.doCommand(ctx, cmd, func() error {
		return cmd.Validate()
	}, opts...)
}

//
// SaleItemUpdate
// @Description: 更新扫描文件
// @receiver s
// @param ctx 上下文
// @param cmd 更新扫描文件
// @return *model.SaleBillCommandDomainService
// @return error
//
func (s *SaleBillCommandDomainService) SaleItemUpdate(ctx context.Context, cmd *command.SaleItemUpdateCommand, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
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
//  @return *model.SaleBillAggregate
//  @return error
//
func (s *SaleBillCommandDomainService) doCommand(ctx context.Context, cmd ddd.Command, validateFunc func() error, opts ...ddd.DoCommandOption) (*model.SaleBillAggregate, error) {
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

	// 执行领域命令
	var err error
	agg := s.NewAggregate()
	if _, ok := cmd.(*command.SaleBillCreateCommand); ok {
		err = ddd.CreateAggregate(ctx, agg, cmd)
	} else {
		err = ddd.CommandAggregate(ctx, agg, cmd)
	}

	// 如果领域命令执行时出错，则返回错误
	if err != nil {
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
// @return *sale_bill_model.SaleBillCommandDomainService  聚合对象
// @return bool 是否找到聚合根对象
// @return error 错误对象
//
func (s *SaleBillCommandDomainService) GetAggregateById(ctx context.Context, tenantId string, id string) (*model.SaleBillAggregate, bool, error) {
	agg := s.NewAggregate()
	_, ok, err := ddd.LoadAggregate(ctx, tenantId, id, agg)
	return agg, ok, err
}

//
// NewAggregate
// @Description: 新建聚合对象
// @receiver s
// @return *sale_bill_model.SaleBillCommandDomainService 聚合对象
//
func (s *SaleBillCommandDomainService) NewAggregate() *model.SaleBillAggregate {
	return model.NewSaleBillAggregate()
}
