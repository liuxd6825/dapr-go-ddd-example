package common

import (
	"github.com/liuxd6825/dapr-go-ddd-example/common/user_models/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func init() {
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
