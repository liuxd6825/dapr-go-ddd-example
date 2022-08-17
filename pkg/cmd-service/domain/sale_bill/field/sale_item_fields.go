package field

//
// SaleItemFields
// @Description: 扫描文件
//
type SaleItemFields struct {
	Id            string  `json:"id" validate:"required" `          // 主键
	InventoryId   string  `json:"inventoryId" validate:"required" ` // 存货Id
	InventoryName string  `json:"inventoryName" validate:"-" `      // 存货名称
	Money         float64 `json:"money" validate:"-" `              // 文件大小
	Quantity      int64   `json:"quantity" validate:"-" `           // 文档id
	Remarks       string  `json:"remarks" validate:"-" `            // 备注
	SaleBillId    string  `json:"saleBillId" validate:"required" `
	TenantId      string  `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleItemFields) GetId() string {
	return f.Id
}

func (f *SaleItemFields) GetTenantId() string {
	return f.TenantId
}
