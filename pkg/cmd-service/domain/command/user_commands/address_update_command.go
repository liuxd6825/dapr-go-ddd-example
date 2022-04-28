package user_commands

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type AddressUpdateCommand struct {
	ddd.BaseCommand
	Data user_fields.AddressFields `json:"data"`
}

func (a *AddressUpdateCommand) NewDomainEvent() ddd.DomainEvent {
	return &user_events.AddressUpdateEventV1{
		TenantId:  a.Data.TenantId,
		CommandId: a.CommandId,
		EventId:   a.CommandId,
		Data:      a.Data,
	}
}

func (a *AddressUpdateCommand) GetTenantId() string {
	return a.Data.TenantId
}

func (a *AddressUpdateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(a.Data.UserId, a.Data.Id)
}
