package user_commands

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_errors"
)

type UserUpdateCommand struct {
	ddd.BaseCommand
	Data user_fields.UserFields `json:"data"`
}

func (c *UserUpdateCommand) NewDomainEvent() ddd.DomainEvent {
	return &user_events.UserUpdateEventV1{
		EventId: c.CommandId,
		Data:    c.Data,
	}
}

func (c *UserUpdateCommand) GetAggregateId() ddd.AggregateId {
	return ddd.NewAggregateId(c.Data.Id)
}

func (c *UserUpdateCommand) GetCommandId() string {
	return c.CommandId
}

func (c *UserUpdateCommand) GetTenantId() string {
	return c.Data.TenantId
}

func (c *UserUpdateCommand) Validate() error {
	errs := ddd_errors.NewVerifyError()
	c.BaseCommand.ValidateError(errs)
	if c.Data.TenantId == "" {
		errs.AppendField("data.tenantId", "不能为空")
	}
	if c.Data.Id == "" {
		errs.AppendField("data.id", "不能为空")
	}
	if c.Data.UserCode == "" {
		errs.AppendField("data.userCode", "不能为空")
	}
	if c.Data.UserName == "" {
		errs.AppendField("data.userName", "不能为空")
	}
	return errs.GetError()
}
