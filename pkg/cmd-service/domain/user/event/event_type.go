package event

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserEventType uint32

const (
	UserCreateEventType UserEventType = iota
	UserUpdateEventType
	UserDeleteEventType
)

//
// String()
// @Description: 转换成字符串
//
func (p UserEventType) String() string {
	switch p {
	case UserCreateEventType:
		return "dapr-ddd-demo.UserCreateEvent"
	case UserUpdateEventType:
		return "dapr-ddd-demo.UserUpdateEvent"
	case UserDeleteEventType:
		return "dapr-ddd-demo.UserDeleteEvent"
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
			Version:   NewUserCreateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &UserCreateEvent{} },
		},
		{
			EventType: UserDeleteEventType.String(),
			Version:   NewUserDeleteEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &UserDeleteEvent{} },
		},
		{
			EventType: UserUpdateEventType.String(),
			Version:   NewUserUpdateEvent("").GetEventVersion(),
			NewFunc:   func() interface{} { return &UserUpdateEvent{} },
		},
	}
}
