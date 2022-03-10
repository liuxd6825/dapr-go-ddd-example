package user_events

type UserCreateEvent struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	CommandId string `json:"commandId"`
	EventId   string `json:"eventId"`
	Code      string `json:"code"`
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
}

func NewUserCreateEvent() *UserCreateEvent {
	return &UserCreateEvent{}
}
