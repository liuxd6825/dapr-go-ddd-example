package user_events

type UserUpdateEvent struct {
	TenantId  string `json:"tenantId"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
	Id        string `json:"id"`
	Name      string `json:"name"`
}

func NewUserUpdateEvent() *UserUpdateEvent {
	return &UserUpdateEvent{}
}
