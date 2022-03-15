package user_events

type UserCreateEventV1 struct {
	Id        string `json:"id"`
	TenantId  string `json:"tenantId"`
	CommandId string `json:"commandId"`

	EventId       string `json:"eventId"`
	EventType     string `json:"eventType"`
	EventRevision string `json:"eventRevision"`

	Code     string `json:"code"`
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func NewUserCreateEventV1() *UserCreateEventV1 {
	return &UserCreateEventV1{}
}

func (e *UserCreateEventV1) GetEventRevision() string {
	return e.EventRevision
}
