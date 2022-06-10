package command

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type AddressUpdateCommand struct {
	CommandId   string               `json:"commandId"`
	UpdateMask  []string             `json:"updateMask"`
	IsValidOnly bool                 `json:"isValidOnly"`
	Data        fields.AddressFields `json:"data"`
}

func (c *AddressUpdateCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

func (c *AddressUpdateCommand) GetUpdateMask() []string {
	return c.UpdateMask
}

func (c *AddressUpdateCommand) Validate() error {
	return ddd.ValidateUpdateCommand(c, nil).GetError()
}

func (c *AddressUpdateCommand) NewDomainEvent() ddd.DomainEvent {
	return &event.AddressUpdateEventV1{
		TenantId:  c.Data.TenantId,
		CommandId: c.CommandId,
		EventId:   c.CommandId,
		Data:      c.Data,
	}
}

func (c *AddressUpdateCommand) GetCommandId() string {
	return c.CommandId
}

func (c *AddressUpdateCommand) GetTenantId() string {
	return c.Data.TenantId
}

func (c *AddressUpdateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.UserId, c.Data.Id)
}
