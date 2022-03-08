package user_events

import "github.com/liuxd6825/dapr-go-ddd-sdk/ddd"

type UserEventType int32

const (
	UserEventType_UserCreateEvent UserEventType = 0
	UserEventType_UserUpdateEvent UserEventType = 1
	UserEventType_UserDeleteEvent UserEventType = 2
)

func (p UserEventType) String() string {
	switch p {
	case UserEventType_UserCreateEvent:
		return "user-create-event"
	case UserEventType_UserUpdateEvent:
		return "user-update-event"
	case UserEventType_UserDeleteEvent:
		return "user-delete-event"
	default:
		return "UNKNOWN"
	}
}

func (p UserEventType) NewDomainEvent() ddd.DomainEvent {
	switch p {
	case UserEventType_UserCreateEvent:
		return &UserCreateEvent{}
	case UserEventType_UserUpdateEvent:
		return &UserUpdateEvent{}
	case UserEventType_UserDeleteEvent:
		return &UserDeleteEvent{}
	default:
		return nil
	}
}
