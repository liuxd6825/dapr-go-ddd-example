package user_model

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// UserAggregate
// @Description:
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

func NewUserAggregate() *UserAggregate {
	return &UserAggregate{
		Addresses: make(map[string]*AddressEntity),
	}
}

func NewAggregate() ddd.Aggregate {
	return NewUserAggregate()
}

func (a *UserAggregate) AddressCreateCommand(ctx context.Context, cmd *user_commands.AddressCreateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewAddressCreateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *UserAggregate) AddressUpdateCommand(ctx context.Context, cmd *user_commands.AddressUpdateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewAddressUpdateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *UserAggregate) AddressDeleteCommand(ctx context.Context, cmd *user_commands.AddressDeleteCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewAddressDeleteEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *UserAggregate) UserCreateCommand(ctx context.Context, cmd *user_commands.UserCreateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewCreateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *UserAggregate) UserUpdateCommand(ctx context.Context, cmd *user_commands.UserUpdateCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewUpdateEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *UserAggregate) UserDeleteCommand(ctx context.Context, cmd *user_commands.UserDeleteCommand, metadata *map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewDeleteEvent(cmd), ddd.ApplyOptions{}.SetMetadata(metadata))
}

func (a *UserAggregate) OnUserAddressCreateEventV1s0(ctx context.Context, event *user_events.AddressCreateEventV1) (err error) {
	fields := &event.Data
	a.Addresses[fields.Id] = NewAddressEntity(fields)
	return nil
}

func (a *UserAggregate) OnUserAddressUpdateEventV1s0(ctx context.Context, event *user_events.AddressUpdateEventV1) error {
	fields := &event.Data
	a.Addresses[fields.Id] = NewAddressEntity(fields)
	return nil
}

func (a *UserAggregate) OnUserAddressDeleteEventV1s0(ctx context.Context, event *user_events.AddressDeleteEventV1) error {
	delete(a.Addresses, event.AddressId)
	return nil
}

func (a *UserAggregate) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Telephone = event.Data.Telephone
	a.Email = event.Data.Email
	return nil
}

func (a *UserAggregate) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Telephone = event.Data.Telephone
	a.Email = event.Data.Email
	return nil
}

func (a *UserAggregate) OnUserDeleteEventV1s0(ctx context.Context, event *user_events.UserDeleteEventV1) error {
	a.IsDelete = true
	return nil
}

func (a *UserAggregate) GetAggregateRevision() string {
	return "1.0"
}

func (a *UserAggregate) GetAggregateType() string {
	return AggregateType
}

func (a *UserAggregate) GetAggregateId() string {
	return a.Id
}

func (a *UserAggregate) GetTenantId() string {
	return a.TenantId
}