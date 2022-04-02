package factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/event/user_events"
	"github.com/liuxd6825/dapr-go-ddd-example/query-service/domain/projection"
)

type UserFactory struct {
}

func NewUserViewByUserCreateEventV1(event *user_events.UserCreateEventV1) *projection.UserView {
	userView := projection.UserView{
		Id:        event.Data.Id,
		UserName:  event.Data.UserName,
		UserCode:  event.Data.UserCode,
		Email:     event.Data.Email,
		Telephone: event.Data.Telephone,
		Address:   event.Data.Address,
		TenantId:  event.TenantId,
	}
	return &userView
}

func NewUserViewByUserUpdateEventV1(event *user_events.UserUpdateEventV1) *projection.UserView {
	userView := projection.UserView{
		Id:        event.Data.Id,
		UserName:  event.Data.UserName,
		UserCode:  event.Data.UserCode,
		Email:     event.Data.Email,
		Telephone: event.Data.Telephone,
		Address:   event.Data.Address,
		TenantId:  event.TenantId,
	}
	return &userView
}
