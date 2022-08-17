package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/field"
	"time"
)

//
// UserCreateEvent
// @Description: 领域事件
//
type UserCreateEvent struct {
	EventId     string                 `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                 `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time              `json:"time" validate:"required"`      // 事件创建时间
	Version     string                 `json:"version" validate:"required"`   // 事件版本
	EventType   string                 `json:"eventType" validate:"required"` // 事件类型
	Data        field.UserCreateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewUserCreateEvent(commandId string) *UserCreateEvent {
	return &UserCreateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, UserCreateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   UserCreateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *UserCreateEvent) GetEventId() string {
	return e.EventId
}

func (e *UserCreateEvent) GetEventType() string {
	return e.EventType
}

func (e *UserCreateEvent) GetEventVersion() string {
	return e.Version
}

func (e *UserCreateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *UserCreateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *UserCreateEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *UserCreateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *UserCreateEvent) GetData() interface{} {
	return e.Data
}
