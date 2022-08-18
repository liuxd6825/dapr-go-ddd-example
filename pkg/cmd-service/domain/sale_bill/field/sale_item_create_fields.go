package field

//
// SaleItemCreateFields
// @Description: 添加明细
//
type SaleItemCreateFields struct {
	Items      []*SaleItemCreateItem `json:"items" validate:"-" `
	SaleBillId string                `json:"saleBillId" validate:"required" `
	TenantId   string                `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleItemCreateFields) GetIds() []string {
	var ids []string
	ids = []string{}
	for _, item := range f.Items {
		ids = append(ids, item.GetId())
	}
	return ids
}
