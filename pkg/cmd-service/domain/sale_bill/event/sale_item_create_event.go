package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"time"
)

//
// SaleItemCreateEvent
// @Description: 领域事件
//
type SaleItemCreateEvent struct {
	EventId     string                     `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                     `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                  `json:"time" validate:"required"`      // 事件创建时间
	Version     string                     `json:"version" validate:"required"`   // 事件版本
	EventType   string                     `json:"eventType" validate:"required"` // 事件类型
	Data        field.SaleItemCreateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewSaleItemCreateEvent(commandId string) *SaleItemCreateEvent {
	return &SaleItemCreateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, SaleItemCreateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   SaleItemCreateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *SaleItemCreateEvent) GetEventId() string {
	return e.EventId
}

func (e *SaleItemCreateEvent) GetEventType() string {
	return e.EventType
}

func (e *SaleItemCreateEvent) GetEventVersion() string {
	return e.Version
}

func (e *SaleItemCreateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *SaleItemCreateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *SaleItemCreateEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *SaleItemCreateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *SaleItemCreateEvent) GetData() interface{} {
	return e.Data
}
