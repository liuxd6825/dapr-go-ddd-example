package user_events

type UserCreateEventV2 struct {
	TenantId      string           `json:"tenantId"`
	CommandId     string           `json:"commandId"`
	EventId       string           `json:"eventId"`
	EventRevision string           `json:"eventRevision"`
	Data          UserCreateDataV2 `json:"data"`
}

type UserCreateDataV2 struct {
	Id       string `json:"id"`
	Code     string `json:"code"`
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
}

func NewUserCreateEventV2() *UserCreateEventV2 {
	return &UserCreateEventV2{}
}

func (e *UserCreateEventV2) GetEventRevision() string {
	return "2.0"
}
