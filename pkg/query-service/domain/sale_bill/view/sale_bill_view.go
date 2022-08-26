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
	SaleMoney     float64        `json:"saleMoney,omitempty"  bson:"sale_money"` // 销售金额
	SaleTime      time.Time      `json:"saleTime,omitempty"  bson:"sale_time"`   // 文件大小
	Statue        SaleBillStatue `json:"statue,omitempty"  bson:"statue"`        // 单据状态
	UserId        string         `json:"userId,omitempty"  bson:"user_id"`       // 用户Id
	UserName      string         `json:"userName,omitempty"  bson:"user_name"`   // 用户名称
}

//
// NewSaleBillView
// @Description: 销售订单
//
func NewSaleBillView() *SaleBillView {
	return &SaleBillView{}
}
