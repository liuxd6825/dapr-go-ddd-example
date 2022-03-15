package user_events

type UserDeleteEvent struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
}

func NewUserDeleteEvent() *UserDeleteEvent {
	return &UserDeleteEvent{}
}
