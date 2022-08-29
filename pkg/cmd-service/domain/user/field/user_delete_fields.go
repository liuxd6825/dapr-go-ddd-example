package field

//
// UserDeleteFields
// @Description: 删除用户
//
type UserDeleteFields struct {
	Id       string `json:"id" validate:"required" `       // 用户Id
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *UserDeleteFields) GetId() string {
	return f.Id
}

func (f *UserDeleteFields) SetId(v string) {
	f.Id = v
}

func (f *UserDeleteFields) GetTenantId() string {
	return f.TenantId
}

func (f *UserDeleteFields) SetTenantId(v string) {
	f.TenantId = v
}
