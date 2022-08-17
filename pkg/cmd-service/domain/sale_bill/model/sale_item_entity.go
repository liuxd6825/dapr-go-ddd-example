package model

//
// SaleItem
// @Description: 销售明细项 实体类型
//
type SaleItem struct {
	Id            string  `json:"id"  validate:"required"`          // 主键
	InventoryId   string  `json:"inventoryId"  validate:"required"` // 存货Id
	InventoryName string  `json:"inventoryName"  validate:"-"`      // 存货名称
	Money         float64 `json:"money"  validate:"-"`              // 文件大小
	Quantity      int64   `json:"quantity"  validate:"-"`           // 文档id
	Remarks       string  `json:"remarks"  validate:"-"`            // 备注
	SaleBillId    string  `json:"saleBillId"  validate:"gt=0"`
	TenantId      string  `json:"tenantId"  validate:"required"` // 租户ID
}

//
// NewSaleItem
// @Description: 新建销售明细项对象
//
func NewSaleItem() *SaleItem {
	return &SaleItem{}
}

func (e *SaleItem) GetId() string {
	return e.Id
}
