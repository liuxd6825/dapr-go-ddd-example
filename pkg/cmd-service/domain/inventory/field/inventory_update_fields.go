package field

//
// InventoryUpdateFields
// @Description: 更新存货档案
//
type InventoryUpdateFields struct {
	Brand    string `json:"brand" validate:"-" `           // 品牌
	Id       string `json:"id" validate:"required" `       // Id
	Keywords string `json:"keywords" validate:"-" `        // 搜索关键字
	Name     string `json:"name" validate:"-" `            // 名称
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	Spec     string `json:"spec" validate:"-" `            // 规格
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *InventoryUpdateFields) GetId() string {
	return f.Id
}

func (f *InventoryUpdateFields) SetId(v string) {
	f.Id = v
}

func (f *InventoryUpdateFields) GetTenantId() string {
	return f.TenantId
}

func (f *InventoryUpdateFields) SetTenantId(v string) {
	f.TenantId = v
}
