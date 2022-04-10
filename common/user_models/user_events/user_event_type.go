package user_events

type UserEventType int32

const (
	UserCreateEventType    UserEventType = 0
	UserUpdateEventType    UserEventType = 1
	UserDeleteEventType    UserEventType = 2
	AddressCreateEventType UserEventType = 3
	AddressUpdateEventType UserEventType = 4
	AddressDeleteEventType UserEventType = 5
)

func (p UserEventType) String() string {
	switch p {
	case UserCreateEventType:
		return "ddd-example.UserCreateEvent"
	case UserUpdateEventType:
		return "ddd-example.UserUpdateEvent"
	case UserDeleteEventType:
		return "ddd-example.UserDeleteEvent"
	case AddressCreateEventType:
		return "ddd-example.UserAddressCreateEvent"
	case AddressUpdateEventType:
		return "ddd-example.UserAddressUpdateEvent"
	case AddressDeleteEventType:
		return "ddd-example.UserAddressDeleteEvent"
	default:
		return "UNKNOWN"
	}
}
