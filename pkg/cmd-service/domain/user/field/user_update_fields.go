package field

//
// UserUpdateFields
// @Description: 更新用户
//
type UserUpdateFields struct {
	Email    string `json:"email" validate:"-" `           // 电子邮箱
	Id       string `json:"id" validate:"required" `       // 用户Id
	Name     string `json:"name" validate:"-" `            // 用户名称
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *UserUpdateFields) GetId() string {
	return f.Id
}

func (f *UserUpdateFields) SetId(v string) {
	f.Id = v
}

func (f *UserUpdateFields) GetTenantId() string {
	return f.TenantId
}

func (f *UserUpdateFields) SetTenantId(v string) {
	f.TenantId = v
}
