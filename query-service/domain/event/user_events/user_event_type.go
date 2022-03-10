package user_events

import "github.com/liuxd6825/dapr-go-ddd-example/common/common_user_event"

func NewUserDomainEvent(eventType common_user_event.UserEventType) interface{} {
	switch eventType {
	case common_user_event.UserCreateEventType:
		return &UserCreateEvent{}
	case common_user_event.UserUpdateEventType:
		return &UserUpdateEvent{}
	case common_user_event.UserDeleteEventType:
		return &UserDeleteEvent{}
	default:
		return nil
	}
}
