package command

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type AddressCreateCommand struct {
	CommandId   string               `json:"commandId"`
	IsValidOnly bool                 `json:"isValidOnly"`
	Data        fields.AddressFields `json:"data"`
}

func (a *AddressCreateCommand) GetIsValidOnly() bool {
	return a.IsValidOnly
}

func (a *AddressCreateCommand) GetCommandId() string {
	return a.CommandId
}

func (a *AddressCreateCommand) Validate() error {
	return nil
}

func (a *AddressCreateCommand) NewDomainEvent() ddd.DomainEvent {
	return &event.AddressCreateEventV1{
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
