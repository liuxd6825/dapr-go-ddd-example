package field

//
// SaleItemUpdateFields
// @Description: 更新明细
//
type SaleItemUpdateFields struct {
	Items      []*SaleItemUpdateItem `json:"items" validate:"-" `
	SaleBillId string                `json:"saleBillId" validate:"required" `
	TenantId   string                `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleItemUpdateFields) GetIds() []string {
	var ids []string
	ids = []string{}
	for _, item := range f.Items {
		ids = append(ids, item.GetId())
	}
	return ids
}
