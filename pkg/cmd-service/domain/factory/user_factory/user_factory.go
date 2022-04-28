package user_factory

import (
	user_commands2 "github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/command/user_commands"
	user_events2 "github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func NewAddressCreateEvent(cmd *user_commands2.AddressCreateCommand) *user_events2.AddressCreateEventV1 {
	return &user_events2.AddressCreateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewAddressUpdateEvent(cmd *user_commands2.AddressUpdateCommand) *user_events2.AddressUpdateEventV1 {
	return &user_events2.AddressUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewAddressDeleteEvent(cmd *user_commands2.AddressDeleteCommand) *user_events2.AddressDeleteEventV1 {
	return &user_events2.AddressDeleteEventV1{
		TenantId:  cmd.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		UserId:    cmd.UserId,
		AddressId: cmd.AddressId,
	}
}

func NewCreateEvent(cmd *user_commands2.UserCreateCommand) *user_events2.UserCreateEventV1 {
	return &user_events2.UserCreateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewUpdateEvent(cmd *user_commands2.UserUpdateCommand) *user_events2.UserUpdateEventV1 {
	return &user_events2.UserUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewDeleteEvent(cmd *user_commands2.UserDeleteCommand) *user_events2.UserDeleteEventV1 {
	return &user_events2.UserDeleteEventV1{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Id,
		TenantId:  cmd.TenantId,
	}
}

func NewEvent(eventType string) ddd.DomainEvent {
	switch eventType {
	case user_events2.UserCreateEventType.String():
		return &user_events2.UserCreateEventV1{}
	case user_events2.UserUpdateEventType.String():
		return &user_events2.UserUpdateEventV1{}
	case user_events2.UserDeleteEventType.String():
		return &user_events2.UserDeleteEventV1{}
	}
	return nil
}
