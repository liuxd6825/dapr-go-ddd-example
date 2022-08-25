package field

//
// SaleItemDeleteItem
// @Description: 删除明细项
//
type SaleItemDeleteItem struct {
	Id       string `json:"id" validate:"required" `       // 明细Id
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleItemDeleteItem) GetId() string {
	return f.Id
}

func (f *SaleItemDeleteItem) GetTenantId() string {
	return f.TenantId
}
