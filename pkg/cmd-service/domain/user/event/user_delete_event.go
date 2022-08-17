package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/field"
	"time"
)

//
// UserDeleteEvent
// @Description: 领域事件
//
type UserDeleteEvent struct {
	EventId     string                 `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                 `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time              `json:"time" validate:"required"`      // 事件创建时间
	Version     string                 `json:"version" validate:"required"`   // 事件版本
	EventType   string                 `json:"eventType" validate:"required"` // 事件类型
	Data        field.UserDeleteFields `json:"data" validate:"required"`      // 业务字段项
}

func NewUserDeleteEvent(commandId string) *UserDeleteEvent {
	return &UserDeleteEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, UserDeleteEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   UserDeleteEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *UserDeleteEvent) GetEventId() string {
	return e.EventId
}

func (e *UserDeleteEvent) GetEventType() string {
	return e.EventType
}

func (e *UserDeleteEvent) GetEventVersion() string {
	return e.Version
}

func (e *UserDeleteEvent) GetCommandId() string {
	return e.CommandId
}

func (e *UserDeleteEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *UserDeleteEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *UserDeleteEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *UserDeleteEvent) GetData() interface{} {
	return e.Data
}
