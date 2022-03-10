package common_user_event

type UserEventType int32

const (
	UserCreateEventType UserEventType = 0
	UserUpdateEventType UserEventType = 1
	UserDeleteEventType UserEventType = 2
)

func (p UserEventType) String() string {
	switch p {
	case UserCreateEventType:
		return "ddd-example.user-create-event"
	case UserUpdateEventType:
		return "ddd-example.user-update-event"
	case UserDeleteEventType:
		return "ddd-example.user-delete-event"
	default:
		return "UNKNOWN"
	}
}
