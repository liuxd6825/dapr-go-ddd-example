package event

import (
	"fmt"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/field"
	"time"
)

//
// SaleBillConfirmEvent
// @Description: 领域事件
//
type SaleBillConfirmEvent struct {
	EventId     string                      `json:"eventId" validate:"required"`   // 领域事件ID
	CommandId   string                      `json:"commandId" validate:"required"` // 关联命令ID
	CreatedTime time.Time                   `json:"time" validate:"required"`      // 事件创建时间
	Version     string                      `json:"version" validate:"required"`   // 事件版本
	EventType   string                      `json:"eventType" validate:"required"` // 事件类型
	UpdateMask  []string                    `json:"updateMask" validate:"-"`       // 更新字段
	Data        field.SaleBillConfirmFields `json:"data" validate:"required"`      // 业务字段项
}

func NewSaleBillConfirmEvent(commandId string) *SaleBillConfirmEvent {
	return &SaleBillConfirmEvent{
		EventId:     fmt.Sprintf("%s(%s)", commandId, SaleBillConfirmEventType.String()),
		CommandId:   commandId,
		Version:     "v1.0",
		EventType:   SaleBillConfirmEventType.String(),
		CreatedTime: time.Now(),
	}
}

func (e *SaleBillConfirmEvent) GetEventId() string {
	return e.EventId
}

func (e *SaleBillConfirmEvent) GetEventType() string {
	return e.EventType
}

func (e *SaleBillConfirmEvent) GetEventVersion() string {
	return e.Version
}

func (e *SaleBillConfirmEvent) GetCommandId() string {
	return e.CommandId
}

func (e *SaleBillConfirmEvent) GetTenantId() string {
	return e.Data.TenantId
}

func (e *SaleBillConfirmEvent) GetAggregateId() string {
	return e.Data.Id
}

func (e *SaleBillConfirmEvent) GetCreatedTime() time.Time {
	return e.CreatedTime
}

func (e *SaleBillConfirmEvent) GetData() interface{} {
	return e.Data
}
