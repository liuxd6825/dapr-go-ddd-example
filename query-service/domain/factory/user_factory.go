package factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
)

type UserFactory struct {
}

func NewUserView_UserCreateEventV1(event *user_events.UserCreateEventV1) *projection.UserView {
	userView := projection.UserView{
		Id:       event.Id,
		UserName: event.UserName,
		TenantId: event.TenantId,
	}
	return &userView
}

func NewUserView_UserCreateEventV2(event *user_events.UserCreateEventV2) *projection.UserView {
	userView := projection.UserView{
		Id:       event.Data.Id,
		UserName: event.Data.UserName,
		TenantId: event.TenantId,
	}
	return &userView
}

func NewUserView_UserUpdateEventV1(event *user_events.UserUpdateEvent) *projection.UserView {
	userView := projection.UserView{
		Id:       event.Id,
		UserName: event.Name,
		TenantId: event.TenantId,
	}
	return &userView
}
