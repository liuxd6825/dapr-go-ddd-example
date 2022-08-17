package field

//
// UserCreateFields
// @Description: 创建用户
//
type UserCreateFields struct {
	Email    string `json:"email" validate:"-" `           // 电子邮箱
	Id       string `json:"id" validate:"required" `       // 用户Id
	Name     string `json:"name" validate:"-" `            // 用户名称
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *UserCreateFields) GetId() string {
	return f.Id
}

func (f *UserCreateFields) GetTenantId() string {
	return f.TenantId
}
