package user_events

type AddressDeleteEventV1 struct {
	TenantId  string `json:"tenantId"`
	CommandId string `json:"commandId"`
	EventId   string `json:"eventId"`
	UserId    string `json:"userId"`
	AddressId string `json:"addressId"`
}

func NewAddressDeleteEventV1() *AddressDeleteEventV1 {
	return &AddressDeleteEventV1{}
}

func (e *AddressDeleteEventV1) GetEventId() string {
	return e.EventId
}

func (e *AddressDeleteEventV1) GetEventType() string {
	return AddressDeleteEventType.String()
}

func (e *AddressDeleteEventV1) GetEventRevision() string {
	return "1.0"
}

func (e *AddressDeleteEventV1) GetCommandId() string {
	return e.CommandId
}

func (e *AddressDeleteEventV1) GetTenantId() string {
	return e.TenantId
}

func (e *AddressDeleteEventV1) GetAggregateId() string {
	return e.UserId
}
