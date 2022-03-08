package user_factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/command/user_commands"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func NewCreateEvent(cmd *user_commands.UserCreateCommand) *user_events.UserCreateEvent {
	return &user_events.UserCreateEvent{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Data.Id,
		TenantId:  cmd.Data.TenantId,
		Code:      cmd.Data.Code,
		UserId:    cmd.Data.UserId,
		UserName:  cmd.Data.UserName,
	}
}

func NewUpdateEvent(cmd *user_commands.UserUpdateCommand) *user_events.UserUpdateEvent {
	return &user_events.UserUpdateEvent{
		CommandId: cmd.CommandId,
		EventId:   cmd.CommandId,
		Id:        cmd.Data.Id,
		TenantId:  cmd.Data.TenantId,
		Code:      cmd.Data.Code,
		UserId:    cmd.Data.UserId,
		UserName:  cmd.Data.UserName,
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
	case user_events.UserEventType_UserCreateEvent.String():
		return &user_events.UserCreateEvent{}
	case user_events.UserEventType_UserUpdateEvent.String():
		return &user_events.UserUpdateEvent{}
	case user_events.UserEventType_UserDeleteEvent.String():
		return &user_events.UserDeleteEvent{}
	}
	return nil
}
