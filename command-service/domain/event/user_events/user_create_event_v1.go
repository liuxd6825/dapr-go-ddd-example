package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/common/event_type"

type UserCreateEventV1 struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	CommandId string `json:"commandId"`
	EventId   string `json:"eventId"`
	Code      string `json:"code"`
	UserName  string `json:"userName"`
}

func NewUserCreateEventV1() *UserCreateEventV1 {
	return &UserCreateEventV1{}
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
	return e.Id
}
