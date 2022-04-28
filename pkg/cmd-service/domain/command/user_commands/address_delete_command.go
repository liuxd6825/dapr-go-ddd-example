package user_commands

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type AddressDeleteCommand struct {
	ddd.BaseCommand
	UserId    string `json:"userId"`
	AddressId string `json:"addressId"`
	TenantId  string `json:"tenantId"`
}

func (a *AddressDeleteCommand) NewDomainEvent() ddd.DomainEvent {
	return &user_events.AddressDeleteEventV1{
		TenantId:  a.TenantId,
		CommandId: a.CommandId,
		EventId:   a.CommandId,
		UserId:    a.UserId,
		AddressId: a.AddressId,
	}
}

func (a *AddressDeleteCommand) GetTenantId() string {
	return a.TenantId
}

func (a *AddressDeleteCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(a.UserId, a.AddressId)
}
