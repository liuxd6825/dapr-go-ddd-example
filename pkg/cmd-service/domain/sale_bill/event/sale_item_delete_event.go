package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"time"
)

//
// SaleItemDeleteEvent
// @Description: 领域事件
//
type SaleItemDeleteEvent struct {
	EventId     string                     `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                     `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                  `json:"time" validate:"required"`      // 事件创建时间
	Version     string                     `json:"version" validate:"required"`   // 事件版本
	EventType   string                     `json:"eventType" validate:"required"` // 事件类型
	Data        field.SaleItemDeleteFields `json:"data" validate:"required"`      // 业务字段项
}

func NewSaleItemDeleteEvent(commandId string) *SaleItemDeleteEvent {
	return &SaleItemDeleteEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, SaleItemDeleteEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   SaleItemDeleteEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *SaleItemDeleteEvent) GetEventId() string {
	return e.EventId
}

func (e *SaleItemDeleteEvent) GetEventType() string {
	return e.EventType
}

func (e *SaleItemDeleteEvent) GetEventVersion() string {
	return e.Version
}

func (e *SaleItemDeleteEvent) GetCommandId() string {
	return e.CommandId
}

func (e *SaleItemDeleteEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *SaleItemDeleteEvent) GetAggregateId() string {
	return e.Data.SaleBillId
}

func (e *SaleItemDeleteEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *SaleItemDeleteEvent) GetData() interface{} {
	return e.Data
}
