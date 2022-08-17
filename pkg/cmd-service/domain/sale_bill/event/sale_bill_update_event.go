package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"time"
)

//
// SaleBillUpdateEvent
// @Description: 领域事件
//
type SaleBillUpdateEvent struct {
	EventId     string                     `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                     `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                  `json:"time" validate:"required"`      // 事件创建时间
	Version     string                     `json:"version" validate:"required"`   // 事件版本
	EventType   string                     `json:"eventType" validate:"required"` // 事件类型
	UpdateMask  []string                   `json:"updateMask" validate:"-"`       // 更新字段
	Data        field.SaleBillUpdateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewSaleBillUpdateEvent(commandId string) *SaleBillUpdateEvent {
	return &SaleBillUpdateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, SaleBillUpdateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   SaleBillUpdateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *SaleBillUpdateEvent) GetEventId() string {
	return e.EventId
}

func (e *SaleBillUpdateEvent) GetEventType() string {
	return e.EventType
}

func (e *SaleBillUpdateEvent) GetEventVersion() string {
	return e.Version
}

func (e *SaleBillUpdateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *SaleBillUpdateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *SaleBillUpdateEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *SaleBillUpdateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *SaleBillUpdateEvent) GetData() interface{} {
	return e.Data
}
