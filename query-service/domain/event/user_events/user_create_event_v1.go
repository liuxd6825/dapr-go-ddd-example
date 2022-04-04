package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/fields"

type UserCreateEventV1 struct {
	TenantId  string            `json:"tenantId"`
	CommandId string            `json:"commandId"`
	EventId   string            `json:"eventId"`
	Data      fields.UserFields `json:"data"`
}

func NewUserCreateEventV1() *UserCreateEventV1 {
	return &UserCreateEventV1{}
}

func (e *UserCreateEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *UserCreateEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *UserCreateEventV1) GetEventId() string {
	return e.EventId
}
