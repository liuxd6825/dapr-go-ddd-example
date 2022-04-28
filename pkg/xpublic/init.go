package xpublic

import (
	user_events2 "github.com/liuxd6825/dapr-go-ddd-example/pkg/xpublic/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func get() {
	_ = ddd.RegisterEventType(user_events2.UserCreateEventType.String(), "1.0", func() interface{} {
		return &user_events2.UserCreateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events2.UserUpdateEventType.String(), "1.0", func() interface{} {
		return &user_events2.UserUpdateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events2.UserDeleteEventType.String(), "1.0", func() interface{} {
		return &user_events2.UserDeleteEventV1{}
	})
	_ = ddd.RegisterEventType(user_events2.AddressCreateEventType.String(), "1.0", func() interface{} {
		return &user_events2.AddressCreateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events2.AddressUpdateEventType.String(), "1.0", func() interface{} {
		return &user_events2.AddressUpdateEventV1{}
	})
	_ = ddd.RegisterEventType(user_events2.AddressDeleteEventType.String(), "1.0", func() interface{} {
		return &user_events2.AddressDeleteEventV1{}
	})
}

func GetRegisterEventTypes() *[]restapp.RegisterEventType {
	return &[]restapp.RegisterEventType{
		{
			EventType: user_events2.UserCreateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events2.UserCreateEventV1{} },
		},
		{
			EventType: user_events2.UserUpdateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events2.UserUpdateEventV1{} },
		},
		{
			EventType: user_events2.UserDeleteEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events2.UserDeleteEventV1{} },
		},
		{
			EventType: user_events2.AddressCreateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events2.AddressCreateEventV1{} },
		},
		{
			EventType: user_events2.AddressUpdateEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events2.AddressUpdateEventV1{} },
		},
		{
			EventType: user_events2.AddressDeleteEventType.String(), Revision: "1.0", NewFunc: func() interface{} { return &user_events2.AddressDeleteEventV1{} },
		},
	}
}
