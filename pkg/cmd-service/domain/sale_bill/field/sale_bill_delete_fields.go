package field

//
// SaleBillDeleteFields
// @Description: 删除销售订单
//
type SaleBillDeleteFields struct {
	Id       string `json:"id" validate:"required" `       // 主键
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleBillDeleteFields) GetId() string {
	return f.Id
}

func (f *SaleBillDeleteFields) SetId(v string) {
	f.Id = v
}

func (f *SaleBillDeleteFields) GetTenantId() string {
	return f.TenantId
}

func (f *SaleBillDeleteFields) SetTenantId(v string) {
	f.TenantId = v
}
