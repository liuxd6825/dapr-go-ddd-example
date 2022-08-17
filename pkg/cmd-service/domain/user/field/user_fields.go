package field

//
// UserFields
// @Description: 创建扫描文件
//
type UserFields struct {
	Email    string `json:"email" validate:"-" `           // 电子邮箱
	Id       string `json:"id" validate:"required" `       // 用户Id
	Name     string `json:"name" validate:"-" `            // 用户名称
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *UserFields) GetId() string {
	return f.Id
}

func (f *UserFields) GetTenantId() string {
	return f.TenantId
}
