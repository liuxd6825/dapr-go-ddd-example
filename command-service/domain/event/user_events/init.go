package user_events

import (
	"github.com/liuxd6825/dapr-go-ddd-example/common/event_type"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

func init() {
	_ = ddd.RegisterEventType(event_type.UserCreateEventType.String(), "1.0", func() interface{} {
		return &UserUpdateEventV1{}
	})
	_ = ddd.RegisterEventType(event_type.UserUpdateEventType.String(), "1.0", func() interface{} {
		return &UserUpdateEventV1{}
	})
	_ = ddd.RegisterEventType(event_type.UserDeleteEventType.String(), "1.0", func() interface{} {
		return &UserDeleteEventV1{}
	})
}
