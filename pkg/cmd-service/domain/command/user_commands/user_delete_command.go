package user_commands

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserDeleteCommand struct {
	TenantId  string `json:"tenantId"`
	CommandId string `json:"commandId"`
	Id        string `json:"id"`
}

func (c *UserDeleteCommand) NewDomainEvent() ddd.DomainEvent {
	return &user_events.UserDeleteEventV1{
		TenantId: c.TenantId,
		EventId:  c.CommandId,
		Id:       c.Id,
	}
}

func (c *UserDeleteCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Id)
}

func (c *UserDeleteCommand) GetCommandId() string {
	return c.CommandId
}

func (c *UserDeleteCommand) GetTenantId() string {
	return c.TenantId
}
