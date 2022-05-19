package user_events

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

type UserEventType uint32

const (
	UserCreateEventType UserEventType = iota
	UserUpdateEventType
	UserDeleteEventType
	AddressCreateEventType
	AddressUpdateEventType
	AddressDeleteEventType
)

func (p UserEventType) String() string {
	switch p {
	case UserCreateEventType:
		return "ddd-example.UserCreateEvent"
	case UserUpdateEventType:
		return "ddd-example.UserUpdateEvent"
	case UserDeleteEventType:
		return "ddd-example.UserDeleteEvent"
	case AddressCreateEventType:
		return "ddd-example.UserAddressCreateEvent"
	case AddressUpdateEventType:
		return "ddd-example.UserAddressUpdateEvent"
	case AddressDeleteEventType:
		return "ddd-example.UserAddressDeleteEvent"
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
			Revision:  (&UserCreateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &UserCreateEventV1{} },
		},
		{
			EventType: UserUpdateEventType.String(),
			Revision:  (&UserUpdateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &UserUpdateEventV1{} },
		},
		{
			EventType: UserDeleteEventType.String(),
			Revision:  (&UserDeleteEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &UserDeleteEventV1{} },
		},
		{
			EventType: AddressCreateEventType.String(),
			Revision:  (&AddressCreateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &AddressCreateEventV1{} },
		},
		{
			EventType: AddressUpdateEventType.String(),
			Revision:  (&AddressUpdateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &AddressUpdateEventV1{} },
		},
		{
			EventType: AddressDeleteEventType.String(),
			Revision:  (&AddressDeleteEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &AddressDeleteEventV1{} },
		},
	}
}
