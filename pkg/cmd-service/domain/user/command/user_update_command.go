package command

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd/ddd_errors"
)

type UserUpdateCommand struct {
	CommandId   string            `json:"commandId"  validate:"gt=0"`
	IsValidOnly bool              `json:"isValidOnly" validate:"gt=0"`
	UpdateMask  []string          `json:"updateMask"`
	Data        fields.UserFields `json:"data"`
}

func (c *UserUpdateCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

func (c *UserUpdateCommand) NewDomainEvent() ddd.DomainEvent {
	return &event.UserUpdateEventV1{
		EventId:    c.CommandId,
		UpdateMask: c.UpdateMask,
		Data:       c.Data,
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

func (c *UserUpdateCommand) GetUpdateMask() []string {
	return c.UpdateMask
}

func (c *UserUpdateCommand) Validate() error {
	errs := ddd_errors.NewVerifyError()
	ddd.ValidateUpdateCommand(c, errs)
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
