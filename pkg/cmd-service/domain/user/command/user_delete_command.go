package command

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserDeleteCommand struct {
	CommandId   string `json:"commandId"  validate:"gt=0"`
	IsValidOnly bool   `json:"isValidOnly"`
	TenantId    string `json:"tenantId"`
	Id          string `json:"id"`
}

func (c *UserDeleteCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

func (c *UserDeleteCommand) NewDomainEvent() ddd.DomainEvent {
	return &event.UserDeleteEventV1{
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

func (c *UserDeleteCommand) Validate() error {
	errs := ddd.ValidateDeleteCommand(c, nil)
	return errs.GetError()
}
