package field

//
// SaleItemDeleteFields
// @Description: 删除扫描单
//
type SaleItemDeleteFields struct {
	Id         string `json:"id" validate:"required" `         // 销售明细Id
	Remarks    string `json:"remarks" validate:"-" `           // 备注
	SaleBillId string `json:"saleBillId" validate:"required" ` // 销售单Id
	TenantId   string `json:"tenantId" validate:"required" `   // 租户ID
}

func (f *SaleItemDeleteFields) GetId() string {
	return f.Id
}

func (f *SaleItemDeleteFields) GetTenantId() string {
	return f.TenantId
}
