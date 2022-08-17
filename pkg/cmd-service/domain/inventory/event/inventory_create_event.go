package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/field"
	"time"
)

//
// InventoryCreateEvent
// @Description: 领域事件
//
type InventoryCreateEvent struct {
	EventId     string                      `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                      `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                   `json:"time" validate:"required"`      // 事件创建时间
	Version     string                      `json:"version" validate:"required"`   // 事件版本
	EventType   string                      `json:"eventType" validate:"required"` // 事件类型
	Data        field.InventoryCreateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewInventoryCreateEvent(commandId string) *InventoryCreateEvent {
	return &InventoryCreateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, InventoryCreateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   InventoryCreateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *InventoryCreateEvent) GetEventId() string {
	return e.EventId
}

func (e *InventoryCreateEvent) GetEventType() string {
	return e.EventType
}

func (e *InventoryCreateEvent) GetEventVersion() string {
	return e.Version
}

func (e *InventoryCreateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *InventoryCreateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *InventoryCreateEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *InventoryCreateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *InventoryCreateEvent) GetData() interface{} {
	return e.Data
}
