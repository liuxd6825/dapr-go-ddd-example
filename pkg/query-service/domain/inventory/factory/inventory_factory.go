package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/inventory/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/inventory/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

type inventoryViewFactory struct {
}

var InventoryView = &inventoryViewFactory{}

func (f *inventoryViewFactory) NewByInventoryUpdateEvent(ctx context.Context, e *event.InventoryUpdateEvent) (*view.InventoryView, error) {
	if e == nil {
		return nil, errors.New("NewByInventoryUpdateEvent(ctx, e) error: e is nil")
	}
	v := &view.InventoryView{}
	setViewType := utils.SetViewUpdated
	v.Brand = e.Data.Brand
	v.Id = e.Data.Id
	v.Keywords = e.Data.Keywords
	v.Name = e.Data.Name
	v.Remarks = e.Data.Remarks
	v.Spec = e.Data.Spec
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *inventoryViewFactory) NewByInventoryCreateEvent(ctx context.Context, e *event.InventoryCreateEvent) (*view.InventoryView, error) {
	if e == nil {
		return nil, errors.New("NewByInventoryCreateEvent(ctx, e) error: e is nil")
	}
	v := &view.InventoryView{}
	setViewType := utils.SetViewCreated
	v.Brand = e.Data.Brand
	v.Id = e.Data.Id
	v.Keywords = e.Data.Keywords
	v.Name = e.Data.Name
	v.Remarks = e.Data.Remarks
	v.Spec = e.Data.Spec
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}
