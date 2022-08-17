package dto

// InventoryCreateCommand

//
// InventoryCreateCommandRequest
// @Description: 创建存货档案
//
type InventoryCreateCommandRequest struct {
	CommandId   string                            `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                              `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	Data        InventoryCreateCommandRequestData `json:"data"  validate:"required"`
}

//
// InventoryCreateCommandRequestData
// @Description: 创建存货档案
//
type InventoryCreateCommandRequestData struct {
	Brand    string `json:"brand,omitempty" validate:"-"`           // 品牌
	Id       string `json:"id,omitempty" validate:"required"`       // Id
	Keywords string `json:"keywords,omitempty" validate:"-"`        // 搜索关键字
	Name     string `json:"name,omitempty" validate:"-"`            // 名称
	Remarks  string `json:"remarks,omitempty" validate:"-"`         // 备注
	Spec     string `json:"spec,omitempty" validate:"-"`            // 规格
	TenantId string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}

//
// InventoryCreateCommandResponse
// @Description: 创建存货档案
type InventoryCreateCommandResponse struct {
}

// InventoryUpdateCommand

//
// InventoryUpdateCommandRequest
// @Description: 更新存货档案
//
type InventoryUpdateCommandRequest struct {
	CommandId   string                            `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                              `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	UpdateMask  []string                          `json:"updateMask"  validate:"-"`       // 要更新的字段项，空值：更新所有字段
	Data        InventoryUpdateCommandRequestData `json:"data"  validate:"required"`
}

//
// InventoryUpdateCommandRequestData
// @Description: 更新存货档案
//
type InventoryUpdateCommandRequestData struct {
	Brand    string `json:"brand,omitempty" validate:"-"`           // 品牌
	Id       string `json:"id,omitempty" validate:"required"`       // Id
	Keywords string `json:"keywords,omitempty" validate:"-"`        // 搜索关键字
	Name     string `json:"name,omitempty" validate:"-"`            // 名称
	Remarks  string `json:"remarks,omitempty" validate:"-"`         // 备注
	Spec     string `json:"spec,omitempty" validate:"-"`            // 规格
	TenantId string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}

//
// InventoryUpdateCommandResponse
// @Description: 更新存货档案
type InventoryUpdateCommandResponse struct {
}

//
// InventoryDto
// @Description: 存货档案  请求或响应业务数据
//
type InventoryDto struct {
	Brand     string `json:"brand,omitempty" validate:"-"`           // 品牌
	Id        string `json:"id,omitempty" validate:"required"`       // Id
	IsDeleted bool   `json:"isDeleted,omitempty" validate:"-"`       // 已删除
	Keywords  string `json:"keywords,omitempty" validate:"-"`        // 搜索关键字
	Name      string `json:"name,omitempty" validate:"-"`            // 名称
	Remarks   string `json:"remarks,omitempty" validate:"-"`         // 备注
	Spec      string `json:"spec,omitempty" validate:"-"`            // 规格
	TenantId  string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}
