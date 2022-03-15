package factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/command-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
)

type UserFactory struct {
}

func (f UserFactory) NewUserView(event *user_events.UserCreateEventV1) *projection.UserView {
	userView := projection.UserView{
		Id:       event.Id,
		UserName: event.UserName,
		TenantId: event.TenantId,
	}
	return &userView
}

func (f UserFactory) NewCreateEvent(event *user_events.UserCreateEventV1) interface{} {
	userView := projection.UserView{
		Id:       event.Id,
		UserName: event.UserName,
		TenantId: event.TenantId,
	}
	return &userView
}
