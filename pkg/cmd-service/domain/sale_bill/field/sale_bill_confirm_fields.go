package field

//
// SaleBillConfirmFields
// @Description: 下单确认
//
type SaleBillConfirmFields struct {
	Id       string `json:"id" validate:"required" `       // 主键
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *SaleBillConfirmFields) GetId() string {
	return f.Id
}

func (f *SaleBillConfirmFields) SetId(v string) {
	f.Id = v
}

func (f *SaleBillConfirmFields) GetTenantId() string {
	return f.TenantId
}

func (f *SaleBillConfirmFields) SetTenantId(v string) {
	f.TenantId = v
}
