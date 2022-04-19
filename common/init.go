package common

import (
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func get() {
	_ = ddd.RegisterEventType(user_events.UserCreateEventType.String(), "1.0", func() interface{} {
		return &user_events.UserCreateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events.UserUpdateEventType.String(), "1.0", func() interface{} {
		return &user_events.UserUpdateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events.UserDeleteEventType.String(), "1.0", func() interface{} {
		return &user_events.UserDeleteEventV1{}
	})
	_ = ddd.RegisterEventType(user_events.AddressCreateEventType.String(), "1.0", func() interface{} {
		return &user_events.AddressCreateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events.AddressUpdateEventType.String(), "1.0", func() interface{} {
		return &user_events.AddressUpdateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events.AddressDeleteEventType.String(), "1.0", func() interface{} {
		return &user_events.AddressDeleteEventV1{}
	})
}

func GetRegisterEventTypes() *[]restapp.RegisterEventType {
	return &[]restapp.RegisterEventType{
		{
			EventType: user_events.UserCreateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events.UserCreateEventV1{} },
		},
		{
			EventType: user_events.UserUpdateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events.UserUpdateEventV1{} },
		},
		{
			EventType: user_events.UserDeleteEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events.UserDeleteEventV1{} },
		},
		{
			EventType: user_events.AddressCreateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events.AddressCreateEventV1{} },
		},
		{
			EventType: user_events.AddressUpdateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events.AddressUpdateEventV1{} },
		},
		{
			EventType: user_events.AddressDeleteEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events.AddressDeleteEventV1{} },
		},
	}
}
