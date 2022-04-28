package user_commands

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type AddressCreateCommand struct {
	ddd.BaseCommand
	Data user_fields.AddressFields `json:"data"`
}

func (a *AddressCreateCommand) NewDomainEvent() ddd.DomainEvent {
	return &user_events.AddressCreateEventV1{
		TenantId:  a.Data.TenantId,
		CommandId: a.CommandId,
		EventId:   a.CommandId,
		Data:      a.Data,
	}
}

func (a *AddressCreateCommand) GetTenantId() string {
	return a.Data.TenantId
}

func (a *AddressCreateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(a.Data.UserId)
}
