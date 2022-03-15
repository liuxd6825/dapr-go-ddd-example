package user_events

type UserUpdateEvent struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
}

func NewUserUpdateEvent() *UserUpdateEvent {
	return &UserUpdateEvent{}
}
