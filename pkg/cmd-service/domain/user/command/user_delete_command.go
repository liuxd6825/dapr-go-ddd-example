package command

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type UserDeleteCommand struct {
	CommandId   string `json:"commandId"  validate:"gt=0"`
	IsValidOnly bool   `json:"isValidOnly"`
	Data        fields.UserDeleteByIdField
}

func (c *UserDeleteCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

func (c *UserDeleteCommand) NewDomainEvent() ddd.DomainEvent {
	return &event.UserDeleteEventV1{
		EventId:   c.CommandId,
		CommandId: c.CommandId,
		TenantId:  c.Data.TenantId,
		Id:        c.Data.TenantId,
	}
}

func (c *UserDeleteCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.Id)
}

func (c *UserDeleteCommand) GetCommandId() string {
	return c.CommandId
}

func (c *UserDeleteCommand) GetTenantId() string {
	return c.Data.TenantId
}

func (c *UserDeleteCommand) Validate() error {
	errs := ddd.ValidateDeleteCommand(c, nil)
	return errs.GetError()
}
