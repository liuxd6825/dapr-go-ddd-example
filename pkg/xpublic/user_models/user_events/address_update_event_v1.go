package user_events

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
)

type AddressUpdateEventV1 struct {
	TenantId  string                    `json:"tenantId"`
	CommandId string                    `json:"commandId"`
	EventId   string                    `json:"eventId"`
	Data      user_fields.AddressFields `json:"data"`
}

func NewAddressUpdateEventV1() *AddressUpdateEventV1 {
	return &AddressUpdateEventV1{}
}

func (e *AddressUpdateEventV1) GetEventId() string {
	return e.EventId
}

func (e *AddressUpdateEventV1) GetEventType() string {
	return AddressUpdateEventType.String()
}

func (e *AddressUpdateEventV1) GetEventRevision() string {
	return "1.0"
}

func (e *AddressUpdateEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *AddressUpdateEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *AddressUpdateEventV1) GetAggregateId() string {
	return e.Data.UserId
}
