package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/user/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/user/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

type userViewFactory struct {
}

var UserView = &userViewFactory{}

func (f *userViewFactory) NewByUserCreateEvent(ctx context.Context, e *event.UserCreateEvent) (*view.UserView, error) {
	if e == nil {
		return nil, errors.New("NewByUserCreateEvent(ctx, e) error: e is nil")
	}
	v := &view.UserView{}
	setViewType := utils.SetViewCreated
	v.Email = e.Data.Email
	v.Id = e.Data.Id
	v.Name = e.Data.Name
	v.Remarks = e.Data.Remarks
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *userViewFactory) NewByUserUpdateEvent(ctx context.Context, e *event.UserUpdateEvent) (*view.UserView, error) {
	if e == nil {
		return nil, errors.New("NewByUserUpdateEvent(ctx, e) error: e is nil")
	}
	v := &view.UserView{}
	setViewType := utils.SetViewUpdated
	v.Email = e.Data.Email
	v.Id = e.Data.Id
	v.Name = e.Data.Name
	v.Remarks = e.Data.Remarks
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *userViewFactory) NewByUserDeleteEvent(ctx context.Context, e *event.UserDeleteEvent) (*view.UserView, error) {
	if e == nil {
		return nil, errors.New("NewByUserDeleteEvent(ctx, e) error: e is nil")
	}
	v := &view.UserView{}
	setViewType := utils.SetViewDeleted
	v.Id = e.Data.Id
	v.Remarks = e.Data.Remarks
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}
