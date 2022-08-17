package model

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/factory"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"time"
)

//
// SaleBillAggregate
// @Description:  <no value> 聚合类型
//
type SaleBillAggregate struct {
	Id        string         `json:"id" validate:"required"`            // 主键
	IsDeleted bool           `json:"isDeleted" validate:"-"`            // 已删除
	Remarks   string         `json:"remarks" validate:"-"`              // 备注
	SaleItems *SaleItemItems `json:"saleItems" copier:"-" validate:"-"` //
	SaleMoney float64        `json:"saleMoney" validate:"-"`            // 销售金额
	SaleTime  time.Time      `json:"saleTime" validate:"-"`             // 文件大小
	TenantId  string         `json:"tenantId" validate:"required"`      // 租户ID
	UserId    string         `json:"userId" validate:"required"`        // 用户Id
	UserName  string         `json:"userName" validate:"-"`             // 用户名称
}

const AggregateType = "dapr-ddd-demo.SaleBillAggregate"

// MaskMapper时不复制的属性
var aggMapperRemove []string

func init() {
	aggMapperRemove = append(aggMapperRemove, "SaleItems")
}

//
// NewSaleBillAggregate
// @Description: 新建<no value> 聚合对象
// @return *SaleBillAggregate
//
func NewSaleBillAggregate() *SaleBillAggregate {
	return &SaleBillAggregate{
		SaleItems: NewSaleItemItems(),
	}
}

//
// NewAggregate
// @Description: 新建当前包中的聚合对象，当前包中只能有一个聚合类型
// @return ddd.Aggregate
//
func NewAggregate() ddd.Aggregate {
	return NewSaleBillAggregate()
}

//
// SaleBillConfirmCommand
// @Description: 执行 SaleBillConfirmCommand 下单确认命令 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillConfirmCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillConfirmCommand(ctx context.Context, cmd *command.SaleBillConfirmCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewSaleBillConfirmEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleBillCreateCommand
// @Description: 执行 SaleBillCreateCommand 创建销售订单 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillCreateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillCreateCommand(ctx context.Context, cmd *command.SaleBillCreateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewSaleBillCreateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.CreateEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleBillDeleteCommand
// @Description: 执行 SaleBillDeleteCommand 删除销售订单 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillDeleteCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillDeleteCommand(ctx context.Context, cmd *command.SaleBillDeleteCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewSaleBillDeleteEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.DeleteEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleBillUpdateCommand
// @Description: 执行 SaleBillUpdateCommand 更新销售订单 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleBillUpdateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleBillUpdateCommand(ctx context.Context, cmd *command.SaleBillUpdateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewSaleBillUpdateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleItemCreateCommand
// @Description: 执行 SaleItemCreateCommand 创建扫描文件 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleItemCreateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleItemCreateCommand(ctx context.Context, cmd *command.SaleItemCreateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewSaleItemCreateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleItemDeleteCommand
// @Description: 执行 SaleItemDeleteCommand 删除扫描单 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleItemDeleteCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleItemDeleteCommand(ctx context.Context, cmd *command.SaleItemDeleteCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewSaleItemDeleteEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// SaleItemUpdateCommand
// @Description: 执行 SaleItemUpdateCommand 更新扫描文件 命令
// @receiver a
// @param ctx 上下文
// @param cmd SaleItemUpdateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *SaleBillAggregate) SaleItemUpdateCommand(ctx context.Context, cmd *command.SaleItemUpdateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewSaleItemUpdateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// OnSaleBillConfirmEvent
// @Description: SaleBillConfirmEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillConfirmEvent(ctx context.Context, e *event.SaleBillConfirmEvent) error {
	panic("SaleBillAggregate.OnSaleBillConfirmEvent to=SaleBill error")
	return nil
}

//
// OnSaleBillCreateEvent
// @Description: SaleBillCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillCreateEvent(ctx context.Context, e *event.SaleBillCreateEvent) error {
	return utils.Mapper(e.Data, a)
}

//
// OnSaleBillDeleteEvent
// @Description: SaleBillDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillDeleteEvent(ctx context.Context, e *event.SaleBillDeleteEvent) error {
	a.IsDeleted = true
	return nil
}

//
// OnSaleBillUpdateEvent
// @Description: SaleBillUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleBillUpdateEvent(ctx context.Context, e *event.SaleBillUpdateEvent) error {
	return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
}

//
// OnSaleItemCreateEvent
// @Description: SaleItemCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemCreateEvent(ctx context.Context, e *event.SaleItemCreateEvent) error {
	_, err := a.SaleItems.AddMapper(ctx, e.Data.Id, &e.Data)
	return err
}

//
// OnSaleItemDeleteEvent
// @Description: SaleItemDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemDeleteEvent(ctx context.Context, e *event.SaleItemDeleteEvent) error {
	return a.SaleItems.DeleteById(ctx, e.Data.Id)
}

//
// OnSaleItemUpdateEvent
// @Description: SaleItemUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *SaleBillAggregate) OnSaleItemUpdateEvent(ctx context.Context, e *event.SaleItemUpdateEvent) error {
	_, _, err := a.SaleItems.UpdateMapper(ctx, e.Data.Id, &e.Data, e.UpdateMask)
	return err
}

//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *SaleBillAggregate) GetAggregateVersion() string {
	return "v1.0"
}

//
// GetAggregateType
// @Description: 获取 聚合的类型
// @receiver a
// @return string 聚合的类型
//
func (a *SaleBillAggregate) GetAggregateType() string {
	return AggregateType
}

//
// GetAggregateId
// @Description: 获取 聚合id
// @receiver a
// @return string 聚合id
//
func (a *SaleBillAggregate) GetAggregateId() string {
	return a.Id
}

//
// GetTenantId
// @Description: 租户id
// @receiver a
// @return string 租户id
//
func (a *SaleBillAggregate) GetTenantId() string {
	return a.TenantId
}
