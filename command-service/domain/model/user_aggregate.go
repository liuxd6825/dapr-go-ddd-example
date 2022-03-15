package model

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/factory/user_factory"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_context"
)

type UserAggregate struct {
	Id       string `json:"id"`
	TenantId string `json:"tenantId"`
	Code     string `json:"code"`
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	IsDelete bool   `json:"isDelete"`
}

func (a *UserAggregate) CreateDomainEvent(ctx context.Context, eventRecord *ddd.EventRecord) ddd.DomainEvent {
	return user_factory.NewEvent(eventRecord.EventType)
}

func (a *UserAggregate) OnCommand(ctx context.Context, domainCmd ddd.DomainCommand) error {
	any := domainCmd.(interface{})
	metadata := ddd_context.GetMetadataContext(ctx)
	switch any.(type) {
	case *user_commands.UserCreateCommand:
		cmd := any.(*user_commands.UserCreateCommand)
		return ddd.Apply(ctx, a, user_factory.NewCreateEvent(cmd), ddd.ApplyMetadata(metadata))
	case *user_commands.UserUpdateCommand:
		cmd := any.(*user_commands.UserUpdateCommand)
		return ddd.Apply(ctx, a, user_factory.NewUpdateEvent(cmd), ddd.ApplyMetadata(metadata))
	case *user_commands.UserDeleteCommand:
		cmd := any.(*user_commands.UserDeleteCommand)
		return ddd.Apply(ctx, a, user_factory.NewDeleteEvent(cmd), ddd.ApplyMetadata(metadata))
	}
	return nil
}

func (a *UserAggregate) OnSourceEvent(ctx context.Context, domainEvent ddd.DomainEvent) error {
	any := domainEvent.(interface{})
	switch any.(type) {
	case *user_events.UserCreateEventV1:
		e, _ := any.(*user_events.UserCreateEventV1)
		return a.OnUserCreateEventV1_0(ctx, e)
	case *user_events.UserUpdateEvent:
		e, _ := any.(*user_events.UserUpdateEvent)
		return a.OnUserUpdateEventV1_0(ctx, e)
	case user_events.UserDeleteEvent:
		e, _ := any.(*user_events.UserDeleteEvent)
		return a.OnUserDeleteEventV1_0(ctx, e)
	}
	return nil
}

func NewUserAggregate() *UserAggregate {
	return &UserAggregate{}
}

func (a *UserAggregate) OnUserCreateEventV1_0(ctx context.Context, event *user_events.UserCreateEventV1) error {
	a.UserName = event.UserName
	a.Id = event.Id
	a.TenantId = event.TenantId
	a.Code = event.Code
	return nil
}

func (a *UserAggregate) OnUserUpdateEventV1_0(ctx context.Context, event *user_events.UserUpdateEvent) error {
	a.UserId = event.UserId
	a.UserName = event.UserName
	a.Id = event.Id
	a.TenantId = event.TenantId
	a.Code = event.Code
	return nil
}

func (a *UserAggregate) OnUserDeleteEventV1_0(ctx context.Context, event *user_events.UserDeleteEvent) error {
	a.Id = event.Id
	a.TenantId = event.TenantId
	a.IsDelete = true
	return nil
}

func (a *UserAggregate) GetAggregateRevision() string {
	return "1.0"
}

func (a *UserAggregate) GetAggregateType() string {
	return "Test.UserAggregate"
}

func (a *UserAggregate) GetAggregateId() string {
	return a.Id
}

func (a *UserAggregate) GetTenantId() string {
	return a.TenantId
}

func (a *UserAggregate) SetTenantId(tenantId string) {
	a.TenantId = tenantId
}
