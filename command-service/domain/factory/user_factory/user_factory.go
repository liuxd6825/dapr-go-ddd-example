package user_factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/common/event_type"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func NewCreateEvent(cmd *user_commands.UserCreateCommand) *user_events.UserCreateEventV1 {
	return &user_events.UserCreateEventV1{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Data.Id,
		TenantId:  cmd.Data.TenantId,
		Code:      cmd.Data.Code,
		UserName:  cmd.Data.Name,
	}
}

func NewUpdateEvent(cmd *user_commands.UserUpdateCommand) *user_events.UserUpdateEvent {
	return &user_events.UserUpdateEvent{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Data.Id,
		TenantId:  cmd.Data.TenantId,
		Code:      cmd.Data.Code,
		UserName:  cmd.Data.Name,
	}
}

func NewDeleteEvent(cmd *user_commands.UserDeleteCommand) *user_events.UserDeleteEvent {
	return &user_events.UserDeleteEvent{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Id,
		TenantId:  cmd.TenantId,
	}
}

func NewEvent(eventType string) ddd.DomainEvent {
	switch eventType {
	case event_type.UserCreateEventType.String():
		return &user_events.UserCreateEventV1{}
	case event_type.UserUpdateEventType.String():
		return &user_events.UserUpdateEvent{}
	case event_type.UserDeleteEventType.String():
		return &user_events.UserDeleteEvent{}
	}
	return nil
}
