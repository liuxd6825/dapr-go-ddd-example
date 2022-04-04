package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/fields"

type UserUpdateEventV1 struct {
	TenantId  string            `json:"tenantId"`
	CommandId string            `json:"commandId"`
	EventId   string            `json:"eventId"`
	Data      fields.UserFields `json:"data"`
}

func NewUserUpdateEvent() *UserUpdateEventV1 {
	return &UserUpdateEventV1{}
}

func (e *UserUpdateEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *UserUpdateEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *UserUpdateEventV1) GetEventId() string {
	return e.EventId
}
