package field

//
// SaleItemCreateItem
// @Description: 创建明细项
//
type SaleItemCreateItem struct {
	Id            string  `json:"id" validate:"required" `          // 主键
	InventoryId   string  `json:"inventoryId" validate:"required" ` // 存货Id
	InventoryName string  `json:"inventoryName" validate:"-" `      // 存货名称
	Money         float64 `json:"money" validate:"-" `              // 销售金额
	Quantity      int64   `json:"quantity" validate:"-" `           // 销售数量
	Remarks       string  `json:"remarks" validate:"-" `            // 备注
	SaleBillId    string  `json:"saleBillId" validate:"gt=0" `
	TenantId      string  `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleItemCreateItem) GetId() string {
	return f.Id
}

func (f *SaleItemCreateItem) SetId(v string) {
	f.Id = v
}

func (f *SaleItemCreateItem) GetTenantId() string {
	return f.TenantId
}

func (f *SaleItemCreateItem) SetTenantId(v string) {
	f.TenantId = v
}
