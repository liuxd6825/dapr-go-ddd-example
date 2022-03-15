package event_type

type UserEventType int32

const (
	UserCreateEventType UserEventType = 0
	UserUpdateEventType UserEventType = 1
	UserDeleteEventType UserEventType = 2
)

func (p UserEventType) String() string {
	switch p {
	case UserCreateEventType:
		return "ddd-example.UserCreateEvent"
	case UserUpdateEventType:
		return "ddd-example.UserUpdateEvent"
	case UserDeleteEventType:
		return "ddd-example.UserDeleteEvent"
	default:
		return "UNKNOWN"
	}
}
