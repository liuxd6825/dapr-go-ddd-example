package user_events

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
)

type AddressCreateEventV1 struct {
	TenantId  string                    `json:"tenantId"`
	CommandId string                    `json:"commandId"`
	EventId   string                    `json:"eventId"`
	Data      user_fields.AddressFields `json:"data"`
}

func NewAddressCreateEventV1() *AddressCreateEventV1 {
	return &AddressCreateEventV1{}
}

func (e *AddressCreateEventV1) GetEventId() string {
	return e.EventId
}

func (e *AddressCreateEventV1) GetEventType() string {
	return AddressCreateEventType.String()
}

func (e *AddressCreateEventV1) GetEventRevision() string {
	return "1.0"
}

func (e *AddressCreateEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *AddressCreateEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *AddressCreateEventV1) GetAggregateId() string {
	return e.Data.UserId
}
