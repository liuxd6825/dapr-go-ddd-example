package user_events

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/fields"
	"github.com/liuxd6825/dapr-go-ddd-example/common/event_type"
)

type UserCreateEventV1 struct {
	TenantId  string            `json:"tenantId"`
	CommandId string            `json:"commandId"`
	EventId   string            `json:"eventId"`
	Data      fields.UserFields `json:"data"`
}

func (e *UserCreateEventV1) GetEventId() string {
	return e.EventId
}

func (e *UserCreateEventV1) GetEventType() string {
	return event_type.UserCreateEventType.String()
}

func (e *UserCreateEventV1) GetEventRevision() string {
	return "1.0"
}

func (e *UserCreateEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *UserCreateEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *UserCreateEventV1) GetAggregateId() string {
	return e.Data.Id
}
