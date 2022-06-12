package command

import (
	"github.com/google/uuid"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/fields"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

//
// UserCreateCommand
// @Description:
//
type UserCreateCommand struct {
	CommandId   string            `json:"commandId"  validate:"gt=0"`
	IsValidOnly bool              `json:"isValidOnly"`
	Data        fields.UserFields `json:"data"`
}

func (c *UserCreateCommand) NewDomainEvent() ddd.DomainEvent {
	eventId := uuid.New().String()
	return &event.UserCreateEventV1{
		EventId:   eventId,
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

func (c *UserCreateCommand) GetIsValidOnly() bool {
	return c.IsValidOnly
}

func (c *UserCreateCommand) Validate() error {
	ve := ddd.ValidateCreateCommand(c, nil)
	if c.Data.Id == "" {
		ve.AppendField("data.id", "不能为空")
	}
	if c.Data.UserCode == "" {
		ve.AppendField("data.userCode", "不能为空")
	}
	if c.Data.UserName == "" {
		ve.AppendField("data.userName", "不能为空")
	}
	return ve.GetError()
}
