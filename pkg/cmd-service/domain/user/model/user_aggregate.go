package model

import (
	"context"

	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/command"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/factory"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// UserAggregate
// @Description:  <no value> 聚合类型
//
type UserAggregate struct {
	Email     string `json:"email" validate:"-"`           // 电子邮箱
	Id        string `json:"id" validate:"required"`       // 用户Id
	IsDeleted bool   `json:"isDeleted" validate:"-"`       // 已删除
	Name      string `json:"name" validate:"-"`            // 用户名称
	Remarks   string `json:"remarks" validate:"-"`         // 备注
	TenantId  string `json:"tenantId" validate:"required"` // 租户ID
}

const AggregateType = "dapr-ddd-demo.UserAggregate"

// MaskMapper时不复制的属性
var aggMapperRemove []string

func init() {
}

//
// NewUserAggregate
// @Description: 新建<no value> 聚合对象
// @return *UserAggregate
//
func NewUserAggregate() *UserAggregate {
	return &UserAggregate{}
}

//
// NewAggregate
// @Description: 新建当前包中的聚合对象，当前包中只能有一个聚合类型
// @return ddd.Aggregate
//
func NewAggregate() ddd.Aggregate {
	return NewUserAggregate()
}

//
// UserCreateCommand
// @Description: 执行 UserCreateCommand 创建用户 命令
// @receiver a
// @param ctx 上下文
// @param cmd UserCreateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *UserAggregate) UserCreateCommand(ctx context.Context, cmd *command.UserCreateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewUserCreateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.CreateEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// UserDeleteCommand
// @Description: 执行 UserDeleteCommand 删除用户 命令
// @receiver a
// @param ctx 上下文
// @param cmd UserDeleteCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *UserAggregate) UserDeleteCommand(ctx context.Context, cmd *command.UserDeleteCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewUserDeleteEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.DeleteEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// UserUpdateCommand
// @Description: 执行 UserUpdateCommand 更新用户 命令
// @receiver a
// @param ctx 上下文
// @param cmd UserUpdateCommand 命令
// @param metadata 元数据
// @return error 错误
//
func (a *UserAggregate) UserUpdateCommand(ctx context.Context, cmd *command.UserUpdateCommand, metadata *map[string]string) (any, error) {
	e, err := factory.NewUserUpdateEvent(ctx, cmd, metadata)
	if err != nil {
		return nil, err
	}
	return ddd.ApplyEvent(ctx, a, e, ddd.NewApplyEventOptions(metadata))
}

//
// OnUserCreateEvent
// @Description: UserCreateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *UserAggregate) OnUserCreateEvent(ctx context.Context, e *event.UserCreateEvent) error {
	return utils.Mapper(e.Data, a)
}

//
// OnUserDeleteEvent
// @Description: UserDeleteEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *UserAggregate) OnUserDeleteEvent(ctx context.Context, e *event.UserDeleteEvent) error {
	a.IsDeleted = true
	return nil
}

//
// OnUserUpdateEvent
// @Description: UserUpdateEvent 领域事件 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *UserAggregate) OnUserUpdateEvent(ctx context.Context, e *event.UserUpdateEvent) error {
	return utils.MaskMapperRemove(e.Data, a, e.UpdateMask, aggMapperRemove)
}

//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *UserAggregate) GetAggregateVersion() string {
	return "v1.0"
}

//
// GetAggregateType
// @Description: 获取 聚合的类型
// @receiver a
// @return string 聚合的类型
//
func (a *UserAggregate) GetAggregateType() string {
	return AggregateType
}

//
// GetAggregateId
// @Description: 获取 聚合id
// @receiver a
// @return string 聚合id
//
func (a *UserAggregate) GetAggregateId() string {
	return a.Id
}

//
// GetTenantId
// @Description: 租户id
// @receiver a
// @return string 租户id
//
func (a *UserAggregate) GetTenantId() string {
	return a.TenantId
}
