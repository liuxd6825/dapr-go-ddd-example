package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/field"
	"time"
)

//
// UserUpdateEvent
// @Description: 领域事件
//
type UserUpdateEvent struct {
	EventId     string                 `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                 `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time              `json:"time" validate:"required"`      // 事件创建时间
	Version     string                 `json:"version" validate:"required"`   // 事件版本
	EventType   string                 `json:"eventType" validate:"required"` // 事件类型
	UpdateMask  []string               `json:"updateMask" validate:"-"`       // 更新字段
	Data        field.UserUpdateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewUserUpdateEvent(commandId string) *UserUpdateEvent {
	return &UserUpdateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, UserUpdateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   UserUpdateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *UserUpdateEvent) GetEventId() string {
	return e.EventId
}

func (e *UserUpdateEvent) GetEventType() string {
	return e.EventType
}

func (e *UserUpdateEvent) GetEventVersion() string {
	return e.Version
}

func (e *UserUpdateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *UserUpdateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *UserUpdateEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *UserUpdateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *UserUpdateEvent) GetData() interface{} {
	return e.Data
}
