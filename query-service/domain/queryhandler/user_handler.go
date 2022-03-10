package queryhandler

import (
	"errors"
	"github.com/liuxd6825/dapr-go-ddd-example/common/common_user_event"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type userQueryHandler struct {
}

func (u *userQueryHandler) doEventHandler(record *ddd.EventRecord, eventType common_user_event.UserEventType, success func(v interface{}) error) error {
	err := ddd.DoEvent(record.EventData, user_events.NewUserDomainEvent(eventType)).OnSuccess(success).OnError(func(err error) {
		println(err)
	}).Error()
	return err
}

func (u *userQueryHandler) OnEvent(record *ddd.EventRecord) (err error) {
	switch record.EventType {
	case common_user_event.UserCreateEventType.String():
		err = u.doEventHandler(record, common_user_event.UserCreateEventType, func(v interface{}) error {
			return u.onCreateEvent(v.(user_events.UserCreateEvent))
		})
		break
	case common_user_event.UserUpdateEventType.String():
		err = u.doEventHandler(record, common_user_event.UserUpdateEventType, func(v interface{}) error {
			return u.onUpdateEvent(v.(user_events.UserUpdateEvent))
		})
		break
	case common_user_event.UserUpdateEventType.String():
		err = u.doEventHandler(record, common_user_event.UserUpdateEventType, func(v interface{}) error {
			return u.onUpdateEvent(v.(user_events.UserUpdateEvent))
		})
		break
	default:
		err = errors.New("")
	}
	return err
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &userQueryHandler{}
}

func (u *userQueryHandler) onCreateEvent(event user_events.UserCreateEvent) error {
	return nil
}

func (u *userQueryHandler) onUpdateEvent(event user_events.UserUpdateEvent) error {
	return nil
}

func (u userQueryHandler) onDeleteEvent(event user_events.UserDeleteEvent) error {
	return nil
}
