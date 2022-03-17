package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/common/event_type"

type UserDeleteEventV1 struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
}

func NewUserDeleteEvent() *UserDeleteEventV1 {
	return &UserDeleteEventV1{}
}

func (e *UserDeleteEventV1) GetEventId() string {
	return e.EventId
}

func (e *UserDeleteEventV1) GetEventType() string {
	return event_type.UserDeleteEventType.String()
}

func (e *UserDeleteEventV1) GetEventRevision() string {
	return "1.0"
}

func (e *UserDeleteEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *UserDeleteEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *UserDeleteEventV1) GetAggregateId() string {
	return e.Id
}
