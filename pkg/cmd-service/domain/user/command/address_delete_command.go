package command

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type AddressDeleteCommand struct {
	CommandId   string `json:"commandId"`
	IsValidOnly bool   `json:"isValidOnly"`
	UserId      string `json:"userId"`
	AddressId   string `json:"addressId"`
	TenantId    string `json:"tenantId"`
}

func (a *AddressDeleteCommand) GetIsValidOnly() bool {
	return a.IsValidOnly
}

func (a *AddressDeleteCommand) GetCommandId() string {
	return a.CommandId
}

func (a *AddressDeleteCommand) Validate() error {
	return nil
}

func (a *AddressDeleteCommand) NewDomainEvent() ddd.DomainEvent {
	return &event.AddressDeleteEventV1{
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
