package event

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func GetRegisterEventTypes() *[]restapp.RegisterEventType {
	return &[]restapp.RegisterEventType{
		{
			EventType: user_events.UserCreateEventType.String(),
			Revision:  (&user_events.UserCreateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &user_events.UserCreateEventV1{} },
		},
		{
			EventType: user_events.UserUpdateEventType.String(),
			Revision:  (&user_events.UserUpdateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &user_events.UserUpdateEventV1{} },
		},
		{
			EventType: user_events.UserDeleteEventType.String(),
			Revision:  (&user_events.UserDeleteEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &user_events.UserDeleteEventV1{} },
		},
		{
			EventType: user_events.AddressCreateEventType.String(),
			Revision:  (&user_events.AddressCreateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &user_events.AddressCreateEventV1{} },
		},
		{
			EventType: user_events.AddressUpdateEventType.String(),
			Revision:  (&user_events.AddressUpdateEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &user_events.AddressUpdateEventV1{} },
		},
		{
			EventType: user_events.AddressDeleteEventType.String(),
			Revision:  (&user_events.AddressDeleteEventV1{}).GetEventRevision(),
			NewFunc:   func() interface{} { return &user_events.AddressDeleteEventV1{} },
		},
	}
}
