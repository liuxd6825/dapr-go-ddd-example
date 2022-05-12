package user_events

import (
	"github.com/liuxd6825/dapr-go-ddd-sdk/restapp"
)

func GetRegisterEventTypes() *[]restapp.RegisterEventType {
	return &[]restapp.RegisterEventType{
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
