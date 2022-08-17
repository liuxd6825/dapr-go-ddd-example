package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"time"
)

//
// SaleBillDeleteEvent
// @Description: 领域事件
//
type SaleBillDeleteEvent struct {
	EventId     string                     `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                     `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                  `json:"time" validate:"required"`      // 事件创建时间
	Version     string                     `json:"version" validate:"required"`   // 事件版本
	EventType   string                     `json:"eventType" validate:"required"` // 事件类型
	Data        field.SaleBillDeleteFields `json:"data" validate:"required"`      // 业务字段项
}

func NewSaleBillDeleteEvent(commandId string) *SaleBillDeleteEvent {
	return &SaleBillDeleteEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, SaleBillDeleteEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   SaleBillDeleteEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *SaleBillDeleteEvent) GetEventId() string {
	return e.EventId
}

func (e *SaleBillDeleteEvent) GetEventType() string {
	return e.EventType
}

func (e *SaleBillDeleteEvent) GetEventVersion() string {
	return e.Version
}

func (e *SaleBillDeleteEvent) GetCommandId() string {
	return e.CommandId
}

func (e *SaleBillDeleteEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *SaleBillDeleteEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *SaleBillDeleteEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *SaleBillDeleteEvent) GetData() interface{} {
	return e.Data
}
