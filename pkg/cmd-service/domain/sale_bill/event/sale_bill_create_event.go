package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"time"
)

//
// SaleBillCreateEvent
// @Description: 领域事件
//
type SaleBillCreateEvent struct {
	EventId     string                     `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                     `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                  `json:"time" validate:"required"`      // 事件创建时间
	Version     string                     `json:"version" validate:"required"`   // 事件版本
	EventType   string                     `json:"eventType" validate:"required"` // 事件类型
	Data        field.SaleBillCreateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewSaleBillCreateEvent(commandId string) *SaleBillCreateEvent {
	return &SaleBillCreateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, SaleBillCreateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   SaleBillCreateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *SaleBillCreateEvent) GetEventId() string {
	return e.EventId
}

func (e *SaleBillCreateEvent) GetEventType() string {
	return e.EventType
}

func (e *SaleBillCreateEvent) GetEventVersion() string {
	return e.Version
}

func (e *SaleBillCreateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *SaleBillCreateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *SaleBillCreateEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *SaleBillCreateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *SaleBillCreateEvent) GetData() interface{} {
	return e.Data
}
