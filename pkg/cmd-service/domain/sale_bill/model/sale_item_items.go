package model

import (
	"context"
	"encoding/json"
	"github.com/liuxd6825/dapr-go-ddd-sdk/types"
)

type SaleItemItems struct {
	types.Items[*SaleItem]
}

func NewSaleItemItems() *SaleItemItems {
	res := &SaleItemItems{}
	res.Init(func() interface{} {
		return &SaleItem{}
	})
	return res
}

func (i *SaleItemItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.Items.MapData())
}

func (i *SaleItemItems) UnmarshalJSON(b []byte) error {
	data := i.Items.MapData()
	return json.Unmarshal(b, &data)
}

func (i *SaleItemItems) AddItem(ctx context.Context, row *SaleItem) error {
	return i.Items.AddItem(ctx, row)
}

func (i *SaleItemItems) DeleteItem(ctx context.Context, row *SaleItem) error {
	return i.Items.Delete(ctx, row)
}

func (i *SaleItemItems) UpdateItem(ctx context.Context, row *SaleItem) error {
	return i.Items.UpdateItem(ctx, row)
}

func (i *SaleItemItems) MapData() map[string]*SaleItem {
	return i.Items.MapData()
}

func (i *SaleItemItems) Length() int {
	m := i.MapData()
	return len(m)
}
