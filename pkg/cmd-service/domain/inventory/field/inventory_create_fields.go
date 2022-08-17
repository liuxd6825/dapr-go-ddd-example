package field

//
// InventoryCreateFields
// @Description: 创建存货档案
//
type InventoryCreateFields struct {
	Brand    string `json:"brand" validate:"-" `           // 品牌
	Id       string `json:"id" validate:"required" `       // Id
	Keywords string `json:"keywords" validate:"-" `        // 搜索关键字
	Name     string `json:"name" validate:"-" `            // 名称
	Remarks  string `json:"remarks" validate:"-" `         // 备注
	Spec     string `json:"spec" validate:"-" `            // 规格
	TenantId string `json:"tenantId" validate:"required" ` // 租户ID
}

func (f *InventoryCreateFields) GetId() string {
	return f.Id
}

func (f *InventoryCreateFields) GetTenantId() string {
	return f.TenantId
}
