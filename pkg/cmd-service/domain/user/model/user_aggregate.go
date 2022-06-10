package model

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// UserAggregate
// @Description: 聚合类型
//
type UserAggregate struct {
	Id        string                    `json:"id" validate:"gt=0"`
	TenantId  string                    `json:"tenantId" validate:"gt=0"`
	UserCode  string                    `json:"code" validate:"gt=0"`
	UserName  string                    `json:"userName" validate:"gt=0"`
	Email     string                    `json:"email" validate:"gt=0"`
	Telephone string                    `json:"telephone" validate:"gt=0"`
	IsDelete  bool                      `json:"isDelete"`
	Addresses map[string]*AddressEntity `json:"addresses"`
}

const AggregateType = "ddd-example.UserAggregate"

//
// NewUserAggregate
// @Description: 新建User聚合对象
// @return *UserAggregate
//
func NewUserAggregate() *UserAggregate {
	return &UserAggregate{
		Addresses: make(map[string]*AddressEntity),
	}
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
// @Description: 命令执行
// @receiver a
// @param ctx 上下文
// @param cmd 命令
// @param metadata 元数据
// @return error 错误
//
func (a *UserAggregate) UserCreateCommand(ctx context.Context, cmd *command.UserCreateCommand, metadata *map[string]string) error {
	return ddd.CreateEvent(ctx, a, factory.NewCreateEvent(cmd), ddd.NewApplyEventOptions(metadata))
}

func (a *UserAggregate) UserUpdateCommand(ctx context.Context, cmd *command.UserUpdateCommand, metadata *map[string]string) error {
	return ddd.ApplyEvent(ctx, a, factory.NewUpdateEvent(cmd), ddd.NewApplyEventOptions(metadata))
}

func (a *UserAggregate) UserDeleteCommand(ctx context.Context, cmd *command.UserDeleteCommand, metadata *map[string]string) error {
	return ddd.DeleteEvent(ctx, a, factory.NewDeleteEvent(cmd), ddd.NewApplyEventOptions(metadata))
}

func (a *UserAggregate) AddressCreateCommand(ctx context.Context, cmd *command.AddressCreateCommand, metadata *map[string]string) error {
	return ddd.ApplyEvent(ctx, a, factory.NewAddressCreateEvent(cmd), ddd.NewApplyEventOptions(metadata))
}

func (a *UserAggregate) AddressUpdateCommand(ctx context.Context, cmd *command.AddressUpdateCommand, metadata *map[string]string) error {
	return ddd.ApplyEvent(ctx, a, factory.NewAddressUpdateEvent(cmd), ddd.NewApplyEventOptions(metadata))
}

func (a *UserAggregate) AddressDeleteCommand(ctx context.Context, cmd *command.AddressDeleteCommand, metadata *map[string]string) error {
	return ddd.ApplyEvent(ctx, a, factory.NewAddressDeleteEvent(cmd), ddd.NewApplyEventOptions(metadata))
}

//
// OnUserAddressCreateEventV1s0
// @Description: AddressCreateEventV1 事件溯源处理器
// @receiver a
// @param ctx 上下文件
// @param event 领域事件
// @return err 错误
//
func (a *UserAggregate) OnUserAddressCreateEventV1s0(ctx context.Context, event *event.AddressCreateEventV1) (err error) {
	fields := &event.Data
	a.Addresses[fields.Id] = NewAddressEntity(fields)
	return nil
}

func (a *UserAggregate) OnUserAddressUpdateEventV1s0(ctx context.Context, event *event.AddressUpdateEventV1) error {
	fields := &event.Data
	a.Addresses[fields.Id] = NewAddressEntity(fields)
	return nil
}

func (a *UserAggregate) OnUserAddressDeleteEventV1s0(ctx context.Context, event *event.AddressDeleteEventV1) error {
	delete(a.Addresses, event.AddressId)
	return nil
}

func (a *UserAggregate) OnUserCreateEventV1s0(ctx context.Context, event *event.UserCreateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Telephone = event.Data.Telephone
	a.Email = event.Data.Email
	return nil
}

func (a *UserAggregate) OnUserUpdateEventV1s0(ctx context.Context, event *event.UserUpdateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Telephone = event.Data.Telephone
	a.Email = event.Data.Email
	return nil
}

func (a *UserAggregate) OnUserDeleteEventV1s0(ctx context.Context, event *event.UserDeleteEventV1) error {
	a.IsDelete = true
	return nil
}

//
// GetAggregateVersion
// @Description: 聚合的版本号
// @receiver a
// @return string 版本号
//
func (a *UserAggregate) GetAggregateVersion() string {
	return "1.0"
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
