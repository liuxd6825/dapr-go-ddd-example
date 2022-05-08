package user_factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/event/user_events"
	projection2 "github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/projection"
)

func NewAddressViewByUserCreateEventV1(event *user_events.AddressCreateEventV1) *projection2.AddressView {
	addressView := projection2.AddressView{
		TenantId: event.Data.TenantId,
		Id:       event.Data.Id,
		UserId:   event.Data.UserId,
		Province: event.Data.Province,
		City:     event.Data.City,
		Area:     event.Data.Area,
		Location: event.Data.Location,
	}
	return &addressView
}

func NewAddressViewByUserUpdateEventV1(event *user_events.AddressUpdateEventV1) *projection2.AddressView {
	userView := projection2.AddressView{
		TenantId: event.Data.TenantId,
		Id:       event.Data.Id,
		UserId:   event.Data.UserId,
		Province: event.Data.Province,
		City:     event.Data.City,
		Area:     event.Data.Area,
		Location: event.Data.Location,
	}
	return &userView
}

func NewUserViewByUserCreateEventV1(event *user_events.UserCreateEventV1) *projection2.UserView {
	userView := projection2.UserView{
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

func NewUserViewByUserUpdateEventV1(event *user_events.UserUpdateEventV1) *projection2.UserView {
	userView := projection2.UserView{
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
