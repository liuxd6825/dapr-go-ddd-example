package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/common/event_type"

type UserCreateEventV2 struct {
	TenantId      string           `json:"tenantId"`
	CommandId     string           `json:"commandId"`
	EventId       string           `json:"eventId"`
	EventRevision string           `json:"eventRevision"`
	Data          UserCreateDataV2 `json:"data"`
}

type UserCreateDataV2 struct {
	Id       string `json:"id"`
	Code     string `json:"code"`
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func NewUserCreateEventV2() *UserCreateEventV2 {
	return &UserCreateEventV2{}
}

func (e *UserCreateEventV2) GetEventId() string {
	return e.EventId
}

func (e *UserCreateEventV2) GetEventType() string {
	return event_type.UserCreateEventType.String()
}

func (e *UserCreateEventV2) GetEventRevision() string {
	return "1.0"
}

func (e *UserCreateEventV2) GetCommandId() string {
	return e.CommandId
}

func (e *UserCreateEventV2) GetTenantId() string {
	return e.TenantId
}

func (e *UserCreateEventV2) GetAggregateId() string {
	return e.Data.Id
}
