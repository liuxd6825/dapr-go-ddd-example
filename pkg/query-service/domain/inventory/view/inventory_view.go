package view

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/domain/view"
)

//
// InventoryView
// @Description: 存货档案
//
type InventoryView struct {
	base.BaseView `bson:",inline"`
	Brand         string `json:"brand,omitempty"  bson:"brand"`       // 品牌
	Keywords      string `json:"keywords,omitempty"  bson:"keywords"` // 搜索关键字
	Name          string `json:"name,omitempty"  bson:"name"`         // 名称
	Spec          string `json:"spec,omitempty"  bson:"spec"`         // 规格
}

//
// NewInventoryView
// @Description: 存货档案
//
func NewInventoryView() *InventoryView {
	return &InventoryView{}
}
