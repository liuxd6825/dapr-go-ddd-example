package event

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/fields"
)

type UserUpdateEventV1 struct {
	TenantId   string            `json:"tenantId"`
	CommandId  string            `json:"commandId"`
	EventId    string            `json:"eventId"`
	UpdateMask []string          `json:"updateMask"`
	Data       fields.UserFields `json:"data"`
}

func NewUserUpdateEventV1() *UserUpdateEventV1 {
	return &UserUpdateEventV1{}
}

func (e *UserUpdateEventV1) GetEventId() string {
	return e.EventId
}

func (e *UserUpdateEventV1) GetEventType() string {
	return UserUpdateEventType.String()
}

func (e *UserUpdateEventV1) GetEventVersion() string {
	return "1.0"
}

func (e *UserUpdateEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *UserUpdateEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *UserUpdateEventV1) GetAggregateId() string {
	return e.Data.Id
}
