package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/common/common_user_event"

type UserDeleteEvent struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
}

func (e *UserDeleteEvent) GetEventId() string {
	return e.EventId
}
func (e *UserDeleteEvent) GetEventType() string {
	return common_user_event.UserDeleteEventType.String()
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
