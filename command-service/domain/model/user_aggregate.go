package model

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserAggregate struct {
	Id        string `json:"id" validate:"gt=0"`
	TenantId  string `json:"tenantId" validate:"gt=0"`
	UserCode  string `json:"code" validate:"gt=0"`
	UserName  string `json:"userName" validate:"gt=0"`
	Email     string `json:"email" validate:"gt=0"`
	Telephone string `json:"telephone" validate:"gt=0"`
	Address   string `json:"address" validate:"gt=0"`
	IsDelete  bool   `json:"isDelete"`
}

func NewUserAggregate() *UserAggregate {
	return &UserAggregate{}
}

func (a *UserAggregate) UserCreateCommand(ctx context.Context, cmd *user_commands.UserCreateCommand, metadata map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewCreateEvent(cmd), ddd.ApplyMetadata(metadata))
}

func (a *UserAggregate) UserUpdateCommand(ctx context.Context, cmd *user_commands.UserUpdateCommand, metadata map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewUpdateEvent(cmd), ddd.ApplyMetadata(metadata))
}

func (a *UserAggregate) UserDeleteCommand(ctx context.Context, cmd *user_commands.UserDeleteCommand, metadata map[string]string) error {
	return ddd.Apply(ctx, a, user_factory.NewDeleteEvent(cmd), ddd.ApplyMetadata(metadata))
}

func (a *UserAggregate) OnUserCreateEventV1s0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Address = event.Data.Address
	a.Telephone = event.Data.Telephone
	return nil
}

func (a *UserAggregate) OnUserUpdateEventV1s0(ctx context.Context, event *user_events.UserUpdateEventV1) error {
	a.Id = event.Data.Id
	a.TenantId = event.TenantId
	a.UserName = event.Data.UserName
	a.UserCode = event.Data.UserCode
	a.Address = event.Data.Address
	a.Telephone = event.Data.Telephone
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
	return "ddd-example.UserAggregate"
}

func (a *UserAggregate) GetAggregateId() string {
	return a.Id
}

func (a *UserAggregate) GetTenantId() string {
	return a.TenantId
}
