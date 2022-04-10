package user_factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func NewAddressCreateEvent(cmd *user_commands.AddressCreateCommand) *user_events.AddressCreateEventV1 {
	return &user_events.AddressCreateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewAddressUpdateEvent(cmd *user_commands.AddressUpdateCommand) *user_events.AddressUpdateEventV1 {
	return &user_events.AddressUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewAddressDeleteEvent(cmd *user_commands.AddressDeleteCommand) *user_events.AddressDeleteEventV1 {
	return &user_events.AddressDeleteEventV1{
		TenantId:  cmd.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		UserId:    cmd.UserId,
		AddressId: cmd.AddressId,
	}
}

func NewCreateEvent(cmd *user_commands.UserCreateCommand) *user_events.UserCreateEventV1 {
	return &user_events.UserCreateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewUpdateEvent(cmd *user_commands.UserUpdateCommand) *user_events.UserUpdateEventV1 {
	return &user_events.UserUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewDeleteEvent(cmd *user_commands.UserDeleteCommand) *user_events.UserDeleteEventV1 {
	return &user_events.UserDeleteEventV1{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Id,
		TenantId:  cmd.TenantId,
	}
}

func NewEvent(eventType string) ddd.DomainEvent {
	switch eventType {
	case user_events.UserCreateEventType.String():
		return &user_events.UserCreateEventV1{}
	case user_events.UserUpdateEventType.String():
		return &user_events.UserUpdateEventV1{}
	case user_events.UserDeleteEventType.String():
		return &user_events.UserDeleteEventV1{}
	}
	return nil
}
