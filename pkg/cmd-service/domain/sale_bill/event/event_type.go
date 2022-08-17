package event

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type SaleBillEventType uint32

const (
	SaleBillUpdateEventType SaleBillEventType = iota
	SaleBillConfirmEventType
	SaleBillDeleteEventType
	SaleItemCreateEventType
	SaleItemUpdateEventType
	SaleItemDeleteEventType
	SaleBillCreateEventType
)

//
// String()
// @Description: 转换成字符串
//
func (p SaleBillEventType) String() string {
	switch p {
	case SaleBillUpdateEventType:
		return "dapr-ddd-demo.SaleBillUpdateEvent"
	case SaleBillConfirmEventType:
		return "dapr-ddd-demo.SaleBillConfirmEvent"
	case SaleBillDeleteEventType:
		return "dapr-ddd-demo.SaleBillDeleteEvent"
	case SaleItemCreateEventType:
		return "dapr-ddd-demo.SaleItemCreateEvent"
	case SaleItemUpdateEventType:
		return "dapr-ddd-demo.SaleItemUpdateEvent"
	case SaleItemDeleteEventType:
		return "dapr-ddd-demo.SaleItemDeleteEvent"
	case SaleBillCreateEventType:
		return "dapr-ddd-demo.SaleBillCreateEvent"
	default:
		return "UNKNOWN"
	}
}

//
// GetRegisterEventTypes
// @Description: 获取聚合根注册事件类型
// @return []restapp.RegisterEventType
//
func GetRegisterEventTypes() []restapp.RegisterEventType {
	return []restapp.RegisterEventType{
		{
			EventType: SaleBillConfirmEventType.String(),
			Version:   (&SaleBillConfirmEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillConfirmEvent{} },
		},
		{
			EventType: SaleBillCreateEventType.String(),
			Version:   (&SaleBillCreateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillCreateEvent{} },
		},
		{
			EventType: SaleBillDeleteEventType.String(),
			Version:   (&SaleBillDeleteEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillDeleteEvent{} },
		},
		{
			EventType: SaleBillUpdateEventType.String(),
			Version:   (&SaleBillUpdateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillUpdateEvent{} },
		},
		{
			EventType: SaleItemCreateEventType.String(),
			Version:   (&SaleItemCreateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleItemCreateEvent{} },
		},
		{
			EventType: SaleItemDeleteEventType.String(),
			Version:   (&SaleItemDeleteEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleItemDeleteEvent{} },
		},
		{
			EventType: SaleItemUpdateEventType.String(),
			Version:   (&SaleItemUpdateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleItemUpdateEvent{} },
		},
	}
}
