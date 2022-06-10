package user_factory

import (
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/domain/user/event"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/domain/user/view"
)

func NewAddressViewByUserCreateEventV1(event *event.AddressCreateEventV1) *view.AddressView {
	addressView := view.AddressView{
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

func NewAddressViewByUserUpdateEventV1(event *event.AddressUpdateEventV1) *view.AddressView {
	userView := view.AddressView{
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

func NewUserViewByUserCreateEventV1(event *event.UserCreateEventV1) *view.UserView {
	userView := view.UserView{
		Id:         event.Data.Id,
		UserName:   event.Data.UserName,
		UserCode:   event.Data.UserCode,
		Email:      event.Data.Email,
		Telephone:  event.Data.Telephone,
		Address:    event.Data.Address,
		TenantId:   event.TenantId,
		CreateTime: event.Data.Birthday,
	}
	return &userView
}

func NewUserViewByUserUpdateEventV1(event *event.UserUpdateEventV1) *view.UserView {
	userView := view.UserView{
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
