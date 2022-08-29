package field

import (
	"time"
)

//
// SaleBillUpdateFields
// @Description: 销售订单
//
type SaleBillUpdateFields struct {
	Id        string    `json:"id" validate:"required" `       // 主键
	Remarks   string    `json:"remarks" validate:"-" `         // 备注
	SaleMoney float64   `json:"saleMoney" validate:"-" `       // 销售金额
	SaleTime  time.Time `json:"saleTime" validate:"-" `        // 文件大小
	TenantId  string    `json:"tenantId" validate:"required" ` // 租户ID
	UserId    string    `json:"userId" validate:"required" `   // 用户Id
	UserName  string    `json:"userName" validate:"-" `        // 用户名称
}

func (f *SaleBillUpdateFields) GetId() string {
	return f.Id
}

func (f *SaleBillUpdateFields) SetId(v string) {
	f.Id = v
}

func (f *SaleBillUpdateFields) GetTenantId() string {
	return f.TenantId
}

func (f *SaleBillUpdateFields) SetTenantId(v string) {
	f.TenantId = v
}
