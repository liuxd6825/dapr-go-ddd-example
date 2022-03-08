package user_events

type UserUpdateEvent struct {
	TenantId  string `json:"tenantId"`
	Id        string `json:"id"`
	EventId   string `json:"eventId"`
	CommandId string `json:"commandId"`
	Code      string `json:"code"`
	UserId    string `json:"userId"`
	UserName  string `json:"userName"`
}

func (e *UserUpdateEvent) GetEventId() string {
	return e.EventId
}

func (e *UserUpdateEvent) GetEventType() string {
	return UserEventType_UserUpdateEvent.String()
}

func (e *UserUpdateEvent) GetEventRevision() string {
	return "1.0"
}

func (e *UserUpdateEvent) GetCommandId() string {
	return e.CommandId
}

func (e *UserUpdateEvent) GetTenantId() string {
	return e.TenantId
}

func (e *UserUpdateEvent) GetAggregateId() string {
	return e.Id
}