package factory

import (
	"context"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/domain/sale_bill/event"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/domain/sale_bill/view"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/utils"
)

type saleItemViewFactory struct {
}

var SaleItemView = &saleItemViewFactory{}

func (f *saleItemViewFactory) NewBySaleItemDeleteEvent(ctx context.Context, e *event.SaleItemDeleteEvent) ([]*view.SaleItemView, error) {
	if e == nil || len(e.Data.Items) == 0 {
		return []*view.SaleItemView{}, nil
	}

	var vList []*view.SaleItemView
	setViewType := utils.SetViewDeleted
	for _, item := range e.Data.Items {
		v := &view.SaleItemView{}
		v.Id = item.Id
		v.TenantId = e.GetTenantId()
		v.SaleBillId = e.Data.SaleBillId
		/*
		   if err := utils.Mapper(item, v); err != nil {
		       return nil, err
		   }*/
		if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
			return nil, err
		}
		vList = append(vList, v)
	}
	return vList, nil
}

func (f *saleItemViewFactory) NewBySaleItemCreateEvent(ctx context.Context, e *event.SaleItemCreateEvent) ([]*view.SaleItemView, error) {
	if e == nil || len(e.Data.Items) == 0 {
		return []*view.SaleItemView{}, nil
	}

	var vList []*view.SaleItemView
	setViewType := utils.SetViewCreated
	for _, item := range e.Data.Items {
		v := &view.SaleItemView{}
		v.Id = item.Id
		v.InventoryId = item.InventoryId
		v.InventoryName = item.InventoryName
		v.Money = item.Money
		v.Quantity = item.Quantity
		v.Remarks = item.Remarks
		v.TenantId = e.GetTenantId()
		/*
		   if err := utils.Mapper(item, v); err != nil {
		       return nil, err
		   }*/
		if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
			return nil, err
		}
		vList = append(vList, v)
	}
	return vList, nil
}

func (f *saleItemViewFactory) NewBySaleItemUpdateEvent(ctx context.Context, e *event.SaleItemUpdateEvent) ([]*view.SaleItemView, error) {
	if e == nil || len(e.Data.Items) == 0 {
		return []*view.SaleItemView{}, nil
	}

	var vList []*view.SaleItemView
	setViewType := utils.SetViewUpdated
	for _, item := range e.Data.Items {
		v := &view.SaleItemView{}
		v.Id = item.Id
		v.InventoryId = item.InventoryId
		v.InventoryName = item.InventoryName
		v.Money = item.Money
		v.Quantity = item.Quantity
		v.Remarks = item.Remarks
		v.TenantId = e.GetTenantId()
		v.SaleBillId = e.Data.SaleBillId
		/*
		   if err := utils.Mapper(item, v); err != nil {
		       return nil, err
		   }*/
		if err := utils.SetViewDefaultFields(ctx, v, e.GetCreatedTime(), setViewType); err != nil {
			return nil, err
		}
		vList = append(vList, v)
	}
	return vList, nil
}
