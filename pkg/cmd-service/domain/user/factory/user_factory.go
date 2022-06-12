package factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/command"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func NewAddressCreateEvent(cmd *command.AddressCreateCommand) *event.AddressCreateEventV1 {
	return &event.AddressCreateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewAddressUpdateEvent(cmd *command.AddressUpdateCommand) *event.AddressUpdateEventV1 {
	return &event.AddressUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewAddressDeleteEvent(cmd *command.AddressDeleteCommand) *event.AddressDeleteEventV1 {
	return &event.AddressDeleteEventV1{
		TenantId:  cmd.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		UserId:    cmd.UserId,
		AddressId: cmd.AddressId,
	}
}

func NewCreateEvent(cmd *command.UserCreateCommand) *event.UserCreateEventV1 {
	return &event.UserCreateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewUpdateEvent(cmd *command.UserUpdateCommand) *event.UserUpdateEventV1 {
	return &event.UserUpdateEventV1{
		TenantId:  cmd.Data.TenantId,
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Data:      cmd.Data,
	}
}

func NewDeleteEvent(cmd *command.UserDeleteCommand) *event.UserDeleteEventV1 {
	return &event.UserDeleteEventV1{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Data.Id,
		TenantId:  cmd.Data.TenantId,
	}
}

func NewEvent(eventType string) ddd.DomainEvent {
	switch eventType {
	case event.UserCreateEventType.String():
		return &event.UserCreateEventV1{}
	case event.UserUpdateEventType.String():
		return &event.UserUpdateEventV1{}
	case event.UserDeleteEventType.String():
		return &event.UserDeleteEventV1{}
	}
	return nil
}
