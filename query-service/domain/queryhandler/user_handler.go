package queryhandler

import (
	"context"
	"errors"
	"github.com/liuxd6825/dapr-go-ddd-example/common/common_user_event"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/queryservice"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/infrastructure/queryservice_impl"
	"github.com/liuxd6825/dapr-go-ddd-sdk/ddd"
)

type userQueryHandler struct {
	service *queryservice.UserQueryService
}

func NewUserQueryHandler() ddd.QueryEventHandler {
	return &userQueryHandler{
		service: queryservice_impl.NewUserQueryService(),
	}
}

func (u *userQueryHandler) doEventHandler(record *ddd.EventRecord, eventType common_user_event.UserEventType, success func(v interface{}) error) error {
	err := ddd.DoEvent(record.EventData, user_events.NewUserDomainEvent(eventType)).OnSuccess(success).OnError(func(err error) {
		println(err)
	}).Error()
	return err
}

func (u *userQueryHandler) OnEvent(ctx context.Context, record *ddd.EventRecord) (err error) {
	switch record.EventType {
	case common_user_event.UserCreateEventType.String():
		err = u.doEventHandler(record, common_user_event.UserCreateEventType, func(v interface{}) error {
			return u.onCreateEvent(ctx, v.(user_events.UserCreateEvent))
		})
		break
	case common_user_event.UserUpdateEventType.String():
		err = u.doEventHandler(record, common_user_event.UserUpdateEventType, func(v interface{}) error {
			return u.onUpdateEvent(ctx, v.(user_events.UserUpdateEvent))
		})
		break
	case common_user_event.UserUpdateEventType.String():
		err = u.doEventHandler(record, common_user_event.UserUpdateEventType, func(v interface{}) error {
			return u.onUpdateEvent(ctx, v.(user_events.UserUpdateEvent))
		})
		break
	default:
		err = errors.New("")
	}
	return err
}

func (u *userQueryHandler) onCreateEvent(ctx context.Context, event user_events.UserCreateEvent) error {
	user := projection.UserView{
		TenantId: event.TenantId,
		Id:       event.Id,
	}
	return u.service.Create(ctx, &user)
}

func (u *userQueryHandler) onUpdateEvent(ctx context.Context, event user_events.UserUpdateEvent) error {
	user := projection.UserView{
		TenantId: event.TenantId,
		Id:       event.Id,
	}
	return u.service.Update(ctx, &user)
}

func (u userQueryHandler) onDeleteEvent(ctx context.Context, event user_events.UserDeleteEvent) error {
	user := projection.UserView{
		TenantId: event.TenantId,
		Id:       event.Id,
	}
	return u.service.Update(ctx, &user)
}
