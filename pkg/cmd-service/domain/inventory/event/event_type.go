package event

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type InventoryEventType uint32

const (
	InventoryCreateEventType InventoryEventType = iota
	InventoryUpdateEventType
)

//
// String()
// @Description: 转换成字符串
//
func (p InventoryEventType) String() string {
	switch p {
	case InventoryCreateEventType:
		return "dapr-ddd-demo.InventoryCreateEvent"
	case InventoryUpdateEventType:
		return "dapr-ddd-demo.InventoryUpdateEvent"
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
			EventType: InventoryCreateEventType.String(),
			Version:   NewInventoryCreateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &InventoryCreateEvent{} },
		},
		{
			EventType: InventoryUpdateEventType.String(),
			Version:   NewInventoryUpdateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &InventoryUpdateEvent{} },
		},
	}
}
