package user_events

type UserDeleteEventV1 struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
}

func NewUserDeleteEvent() *UserDeleteEventV1 {
	return &UserDeleteEventV1{}
}
