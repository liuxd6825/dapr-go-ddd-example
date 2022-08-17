package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
	"github.com/liuxd6825/dapr-go-ddd-sdk/errors"
)

type saleItemViewFactory struct {
}

var SaleItemView = &saleItemViewFactory{}

func (f *saleItemViewFactory) NewBySaleItemCreateEvent(ctx context.Context, e *event.SaleItemCreateEvent) (*view.SaleItemView, error) {
	if e == nil {
		return nil, errors.New("NewBySaleItemCreateEvent(ctx, e) error: e is nil")
	}
	v := &view.SaleItemView{}
	setViewType := utils.SetViewCreated
	v.Id = e.Data.Id
	v.InventoryId = e.Data.InventoryId
	v.InventoryName = e.Data.InventoryName
	v.Money = e.Data.Money
	v.Quantity = e.Data.Quantity
	v.Remarks = e.Data.Remarks
	v.SaleBillId = e.Data.SaleBillId
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *saleItemViewFactory) NewBySaleItemUpdateEvent(ctx context.Context, e *event.SaleItemUpdateEvent) (*view.SaleItemView, error) {
	if e == nil {
		return nil, errors.New("NewBySaleItemUpdateEvent(ctx, e) error: e is nil")
	}
	v := &view.SaleItemView{}
	setViewType := utils.SetViewUpdated
	v.Id = e.Data.Id
	v.InventoryId = e.Data.InventoryId
	v.InventoryName = e.Data.InventoryName
	v.Money = e.Data.Money
	v.Quantity = e.Data.Quantity
	v.Remarks = e.Data.Remarks
	v.SaleBillId = e.Data.SaleBillId
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}

func (f *saleItemViewFactory) NewBySaleItemDeleteEvent(ctx context.Context, e *event.SaleItemDeleteEvent) (*view.SaleItemView, error) {
	if e == nil {
		return nil, errors.New("NewBySaleItemDeleteEvent(ctx, e) error: e is nil")
	}
	v := &view.SaleItemView{}
	setViewType := utils.SetViewDeleted
	v.Id = e.Data.Id
	v.Remarks = e.Data.Remarks
	v.SaleBillId = e.Data.SaleBillId
	v.TenantId = e.Data.TenantId
	if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
		return nil, err
	}
	return v, nil
}
