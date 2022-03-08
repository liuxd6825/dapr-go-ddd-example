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

func (e *UserCreateEvent) GetEventId() string {
	return e.EventId
}

func (e *UserCreateEvent) GetEventType() string {
	return UserEventType_UserCreateEvent.String()
}

func (e *UserCreateEvent) GetEventRevision() string {
	return "1.0"
}

func (e *UserCreateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *UserCreateEvent) GetTenantId() string {
	return e.TenantId
}

func (e *UserCreateEvent) GetAggregateId() string {
	return e.Id
}
