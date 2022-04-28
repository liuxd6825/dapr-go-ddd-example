package user_commands

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserCreateCommand struct {
	ddd.BaseCommand
	Data user_fields.UserFields `json:"data"`
}

func (c *UserCreateCommand) NewDomainEvent() ddd.DomainEvent {
	return &user_events.UserCreateEventV1{
		EventId:   c.CommandId,
		CommandId: c.CommandId,
		Data:      c.Data,
	}
}

func (c *UserCreateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.Id)
}

func (c *UserCreateCommand) GetCommandId() string {
	return c.CommandId
}

func (c *UserCreateCommand) GetTenantId() string {
	return c.Data.TenantId
}

func (c *UserCreateCommand) Validate() error {
	return nil
}
