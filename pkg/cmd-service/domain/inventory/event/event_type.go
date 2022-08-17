package event

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type InventoryEventType uint32

const (
	InventoryUpdateEventType InventoryEventType = iota
	InventoryCreateEventType
)

//
// String()
// @Description: 转换成字符串
//
func (p InventoryEventType) String() string {
	switch p {
	case InventoryUpdateEventType:
		return "dapr-ddd-demo.InventoryUpdateEvent"
	case InventoryCreateEventType:
		return "dapr-ddd-demo.InventoryCreateEvent"
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
			Version:   (&InventoryCreateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &InventoryCreateEvent{} },
		},
		{
			EventType: InventoryUpdateEventType.String(),
			Version:   (&InventoryUpdateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &InventoryUpdateEvent{} },
		},
	}
}
