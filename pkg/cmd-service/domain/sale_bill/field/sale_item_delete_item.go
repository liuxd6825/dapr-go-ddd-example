package field

//
// SaleItemDeleteItem
// @Description: 删除明细项
//
type SaleItemDeleteItem struct {
	Id         string `json:"id" validate:"required" ` // 明细Id
	Remarks    string `json:"remarks" validate:"-" `   // 备注
	SaleBillId string `json:"saleBillId" validate:"gt=0" `
	TenantId   string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleItemDeleteItem) GetId() string {
	return f.Id
}

func (f *SaleItemDeleteItem) SetId(v string) {
	f.Id = v
}

func (f *SaleItemDeleteItem) GetTenantId() string {
	return f.TenantId
}

func (f *SaleItemDeleteItem) SetTenantId(v string) {
	f.TenantId = v
}
