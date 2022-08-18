package field

//
// SaleItemDeleteFields
// @Description: 删除销售明细项
//
type SaleItemDeleteFields struct {
	Items      []*SaleItemDeleteItem `json:"items" validate:"-" `
	SaleBillId string                `json:"saleBillId" validate:"required" ` // 销售单Id
	TenantId   string                `json:"tenantId" validate:"required" `   // 租户ID
}

func (f *SaleItemDeleteFields) GetIds() []string {
	var ids []string
	ids = []string{}
	for _, item := range f.Items {
		ids = append(ids, item.GetId())
	}
	return ids
}
