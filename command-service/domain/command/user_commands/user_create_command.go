package user_commands

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserCreateCommand struct {
	ddd.BaseDomainCommand
	Data fields.UserFields `json:"data"`
}

func (c *UserCreateCommand) NewDomainEvent() ddd.DomainEvent {
	return &user_events.UserCreateEvent{
		EventId:   c.CommandId,
		CommandId: c.CommandId,
		Id:        c.Data.Id,
		TenantId:  c.Data.TenantId,
		Code:      c.Data.Code,
		UserId:    c.Data.UserId,
		UserName:  c.Data.UserName,
	}
}

func (c *UserCreateCommand) GetAggregateId() string {
	return c.Data.Id
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
