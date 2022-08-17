package event

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserEventType uint32

const (
	UserDeleteEventType UserEventType = iota
	UserCreateEventType
	UserUpdateEventType
)

//
// String()
// @Description: 转换成字符串
//
func (p UserEventType) String() string {
	switch p {
	case UserDeleteEventType:
		return "dapr-ddd-demo.UserDeleteEvent"
	case UserCreateEventType:
		return "dapr-ddd-demo.UserCreateEvent"
	case UserUpdateEventType:
		return "dapr-ddd-demo.UserUpdateEvent"
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
			EventType: UserCreateEventType.String(),
			Version:   (&UserCreateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &UserCreateEvent{} },
		},
		{
			EventType: UserDeleteEventType.String(),
			Version:   (&UserDeleteEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &UserDeleteEvent{} },
		},
		{
			EventType: UserUpdateEventType.String(),
			Version:   (&UserUpdateEvent{}).GetEventVersion(),
			NewFunc:   func() interface{} { return &UserUpdateEvent{} },
		},
	}
}
