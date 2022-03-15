package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/common/event_type"

type UserDeleteEvent struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
}

func NewUserDeleteEvent() *UserDeleteEvent {
	return &UserDeleteEvent{}
}

func (e *UserDeleteEvent) GetEventId() string {
	return e.EventId
}
func (e *UserDeleteEvent) GetEventType() string {
	return event_type.UserDeleteEventType.String()
}
func (e *UserDeleteEvent) GetEventRevision() string {
	return "1.0"
}

func (e *UserDeleteEvent) GetCommandId() string {
	return e.CommandId
}

func (e *UserDeleteEvent) GetTenantId() string {
	return e.TenantId
}

func (e *UserDeleteEvent) GetAggregateId() string {
	return e.Id
}
