package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/field"
	"time"
)

//
// InventoryUpdateEvent
// @Description: 领域事件
//
type InventoryUpdateEvent struct {
	EventId     string                      `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                      `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                   `json:"time" validate:"required"`      // 事件创建时间
	Version     string                      `json:"version" validate:"required"`   // 事件版本
	EventType   string                      `json:"eventType" validate:"required"` // 事件类型
	UpdateMask  []string                    `json:"updateMask" validate:"-"`       // 更新字段
	Data        field.InventoryUpdateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewInventoryUpdateEvent(commandId string) *InventoryUpdateEvent {
	return &InventoryUpdateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, InventoryUpdateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   InventoryUpdateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *InventoryUpdateEvent) GetEventId() string {
	return e.EventId
}

func (e *InventoryUpdateEvent) GetEventType() string {
	return e.EventType
}

func (e *InventoryUpdateEvent) GetEventVersion() string {
	return e.Version
}

func (e *InventoryUpdateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *InventoryUpdateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *InventoryUpdateEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *InventoryUpdateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *InventoryUpdateEvent) GetData() interface{} {
	return e.Data
}
