package view

import (
	base "gitee.com/liuxu6825/dapr-ddd-demo/pkg/query-service/infrastructure/base/domain/view"
	"time"
)

//
// SaleBillView
// @Description: 销售订单
//
type SaleBillView struct {
	base.BaseView `bson:",inline"`
	SaleMoney     float64   `json:"saleMoney,omitempty"  bson:"saleMoney"` // 销售金额
	SaleTime      time.Time `json:"saleTime,omitempty"  bson:"saleTime"`   // 文件大小
	UserId        string    `json:"userId,omitempty"  bson:"userId"`       // 用户Id
	UserName      string    `json:"userName,omitempty"  bson:"userName"`   // 用户名称
}

//
// NewSaleBillView
// @Description: 销售订单
//
func NewSaleBillView() *SaleBillView {
	return &SaleBillView{}
}
