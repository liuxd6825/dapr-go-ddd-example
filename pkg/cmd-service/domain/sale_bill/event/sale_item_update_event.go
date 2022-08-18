package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"time"
)

//
// SaleItemUpdateEvent
// @Description: 领域事件
//
type SaleItemUpdateEvent struct {
	EventId     string                     `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                     `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                  `json:"time" validate:"required"`      // 事件创建时间
	Version     string                     `json:"version" validate:"required"`   // 事件版本
	EventType   string                     `json:"eventType" validate:"required"` // 事件类型
	UpdateMask  []string                   `json:"updateMask" validate:"-"`       // 更新字段
	Data        field.SaleItemUpdateFields `json:"data" validate:"required"`      // 业务字段项
}

func NewSaleItemUpdateEvent(commandId string) *SaleItemUpdateEvent {
	return &SaleItemUpdateEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, SaleItemUpdateEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   SaleItemUpdateEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *SaleItemUpdateEvent) GetEventId() string {
	return e.EventId
}

func (e *SaleItemUpdateEvent) GetEventType() string {
	return e.EventType
}

func (e *SaleItemUpdateEvent) GetEventVersion() string {
	return e.Version
}

func (e *SaleItemUpdateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *SaleItemUpdateEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *SaleItemUpdateEvent) GetAggregateId() string {
	return e.Data.SaleBillId
}

func (e *SaleItemUpdateEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *SaleItemUpdateEvent) GetData() interface{} {
	return e.Data
}
