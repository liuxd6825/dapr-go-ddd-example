package view

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/domain/view"
)

//
// SaleItemView
// @Description: 销售明细项
//
type SaleItemView struct {
	base.BaseView `bson:",inline"`
	InventoryId   string  `json:"inventoryId,omitempty"  bson:"inventory_id"`     // 存货Id
	InventoryName string  `json:"inventoryName,omitempty"  bson:"inventory_name"` // 存货名称
	Money         float64 `json:"money,omitempty"  bson:"money"`                  // 文件大小
	Quantity      int64   `json:"quantity,omitempty"  bson:"quantity"`            // 数量
	SaleBillId    string  `json:"saleBillId,omitempty"  bson:"sale_bill_id" validate:"gt=0"`
}

//
// NewSaleItemView
// @Description: 创建 销售明细项 视图对象
//
func NewSaleItemView() *SaleItemView {
	return &SaleItemView{}
}
