package event

//
// BaseDomainEvent
// @Description: 事件基类
//
type BaseDomainEvent struct {
	EventId     string    `json:"eventId"     validate:"gt=0"` // 领域事件ID
	CommandId   string    `json:"commandId"   validate:"gt=0"` // 关联命令ID
	IsValidOnly bool      `json:"isValidOnly" validate:"gt=0"` // 是否仅验证，不执行
	Time        time.Time `json:"time"        validate:"gt=0"` // 事件创建时间
}

func (e *BaseDomainEvent) GetEventId() string {
	return e.EventId
}

func (e *BaseDomainEvent) GetCommandId() string {
	return e.CommandId
}

func (e *BaseDomainEvent) GetIsValidOnly() bool {
	return e.IsValidOnly
}

func (e *BaseDomainEvent) GetTime() time.Time {
	return e.Time
}

//
// UpdateDomainEvent
// @Description: 更新事件基类
//
type UpdateDomainEvent struct {
	UpdateMask []string `json:"updateMask"` // 要更新的字段项，空值：更新所有字段
}

func (e *UpdateDomainEvent) GetUpdateMask() []string {
	return e.UpdateMask
}
