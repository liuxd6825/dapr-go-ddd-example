package event

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type SaleBillEventType uint32

const (
	SaleBillConfirmEventType SaleBillEventType = iota
	SaleBillDeleteEventType
	SaleItemDeleteEventType
	SaleItemCreateEventType
	SaleItemUpdateEventType
	SaleBillCreateEventType
	SaleBillUpdateEventType
)

//
// String()
// @Description: 转换成字符串
//
func (p SaleBillEventType) String() string {
	switch p {
	case SaleBillConfirmEventType:
		return "dapr-ddd-demo.SaleBillConfirmEvent"
	case SaleBillDeleteEventType:
		return "dapr-ddd-demo.SaleBillDeleteEvent"
	case SaleItemDeleteEventType:
		return "dapr-ddd-demo.SaleItemDeleteEvent"
	case SaleItemCreateEventType:
		return "dapr-ddd-demo.SaleItemCreateEvent"
	case SaleItemUpdateEventType:
		return "dapr-ddd-demo.SaleItemUpdateEvent"
	case SaleBillCreateEventType:
		return "dapr-ddd-demo.SaleBillCreateEvent"
	case SaleBillUpdateEventType:
		return "dapr-ddd-demo.SaleBillUpdateEvent"
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
			Version:   NewSaleBillConfirmEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillConfirmEvent{} },
		},
		{
			EventType: SaleBillCreateEventType.String(),
			Version:   NewSaleBillCreateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillCreateEvent{} },
		},
		{
			EventType: SaleBillDeleteEventType.String(),
			Version:   NewSaleBillDeleteEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillDeleteEvent{} },
		},
		{
			EventType: SaleBillUpdateEventType.String(),
			Version:   NewSaleBillUpdateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleBillUpdateEvent{} },
		},
		{
			EventType: SaleItemCreateEventType.String(),
			Version:   NewSaleItemCreateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleItemCreateEvent{} },
		},
		{
			EventType: SaleItemDeleteEventType.String(),
			Version:   NewSaleItemDeleteEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleItemDeleteEvent{} },
		},
		{
			EventType: SaleItemUpdateEventType.String(),
			Version:   NewSaleItemUpdateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &SaleItemUpdateEvent{} },
		},
	}
}
