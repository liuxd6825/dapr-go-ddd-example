package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

type saleBillViewFactory struct {
}

var SaleBillView = &saleBillViewFactory{}

func (f *saleBillViewFactory) NewBySaleBillUpdateEvent(ctx context.Context, e *event.SaleBillUpdateEvent) (*view.SaleBillView, error) {
	if e == nil {
		return nil, errors.New("NewBySaleBillUpdateEvent(ctx, e) error: e is nil")
	}
	v := &view.SaleBillView{}
	setViewType := utils.SetViewUpdated
	v.Id = e.Data.Id
	v.Remarks = e.Data.Remarks
	v.SaleMoney = e.Data.SaleMoney
	v.SaleTime = e.Data.SaleTime
	v.TenantId = e.Data.TenantId
	v.UserId = e.Data.UserId
	v.UserName = e.Data.UserName
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *saleBillViewFactory) NewBySaleBillConfirmEvent(ctx context.Context, e *event.SaleBillConfirmEvent) (*view.SaleBillView, error) {
	if e == nil {
		return nil, errors.New("NewBySaleBillConfirmEvent(ctx, e) error: e is nil")
	}
	v := &view.SaleBillView{}
	setViewType := utils.SetViewUpdated
	v.Id = e.Data.Id
	v.Remarks = e.Data.Remarks
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *saleBillViewFactory) NewBySaleBillDeleteEvent(ctx context.Context, e *event.SaleBillDeleteEvent) (*view.SaleBillView, error) {
	if e == nil {
		return nil, errors.New("NewBySaleBillDeleteEvent(ctx, e) error: e is nil")
	}
	v := &view.SaleBillView{}
	setViewType := utils.SetViewDeleted
	v.Id = e.Data.Id
	v.Remarks = e.Data.Remarks
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *saleBillViewFactory) NewBySaleBillCreateEvent(ctx context.Context, e *event.SaleBillCreateEvent) (*view.SaleBillView, error) {
	if e == nil {
		return nil, errors.New("NewBySaleBillCreateEvent(ctx, e) error: e is nil")
	}
	v := &view.SaleBillView{}
	setViewType := utils.SetViewCreated
	v.Id = e.Data.Id
	v.Remarks = e.Data.Remarks
	v.SaleMoney = e.Data.SaleMoney
	v.SaleTime = e.Data.SaleTime
	v.TenantId = e.Data.TenantId
	v.UserId = e.Data.UserId
	v.UserName = e.Data.UserName
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}
