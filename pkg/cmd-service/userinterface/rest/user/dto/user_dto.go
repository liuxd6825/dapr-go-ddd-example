package dto

// UserDeleteCommand

//
// UserDeleteCommandRequest
// @Description: 删除用户
//
type UserDeleteCommandRequest struct {
	CommandId   string                       `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                         `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	Data        UserDeleteCommandRequestData `json:"data"  validate:"required"`
}

//
// UserDeleteCommandRequestData
// @Description: 删除用户
//
type UserDeleteCommandRequestData struct {
	Id       string `json:"id,omitempty" validate:"required"`       // 用户Id
	Remarks  string `json:"remarks,omitempty" validate:"-"`         // 备注
	TenantId string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}

//
// UserDeleteCommandResponse
// @Description: 删除用户
type UserDeleteCommandResponse struct {
}

// UserCreateCommand

//
// UserCreateCommandRequest
// @Description: 创建用户
//
type UserCreateCommandRequest struct {
	CommandId   string                       `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                         `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	Data        UserCreateCommandRequestData `json:"data"  validate:"required"`
}

//
// UserCreateCommandRequestData
// @Description: 创建用户
//
type UserCreateCommandRequestData struct {
	Email    string `json:"email,omitempty" validate:"-"`           // 电子邮箱
	Id       string `json:"id,omitempty" validate:"required"`       // 用户Id
	Name     string `json:"name,omitempty" validate:"-"`            // 用户名称
	Remarks  string `json:"remarks,omitempty" validate:"-"`         // 备注
	TenantId string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}

//
// UserCreateCommandResponse
// @Description: 创建用户
type UserCreateCommandResponse struct {
}

// UserUpdateCommand

//
// UserUpdateCommandRequest
// @Description: 更新用户
//
type UserUpdateCommandRequest struct {
	CommandId   string                       `json:"commandId"  validate:"required"` // 命令ID
	IsValidOnly bool                         `json:"isValidOnly"  validate:"-"`      // 是否仅验证，不执行
	UpdateMask  []string                     `json:"updateMask"  validate:"-"`       // 要更新的字段项，空值：更新所有字段
	Data        UserUpdateCommandRequestData `json:"data"  validate:"required"`
}

//
// UserUpdateCommandRequestData
// @Description: 更新用户
//
type UserUpdateCommandRequestData struct {
	Email    string `json:"email,omitempty" validate:"-"`           // 电子邮箱
	Id       string `json:"id,omitempty" validate:"required"`       // 用户Id
	Name     string `json:"name,omitempty" validate:"-"`            // 用户名称
	Remarks  string `json:"remarks,omitempty" validate:"-"`         // 备注
	TenantId string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}

//
// UserUpdateCommandResponse
// @Description: 更新用户
type UserUpdateCommandResponse struct {
}

//
// UserDto
// @Description: 用户  请求或响应业务数据
//
type UserDto struct {
	Email     string `json:"email,omitempty" validate:"-"`           // 电子邮箱
	Id        string `json:"id,omitempty" validate:"required"`       // 用户Id
	IsDeleted bool   `json:"isDeleted,omitempty" validate:"-"`       // 已删除
	Name      string `json:"name,omitempty" validate:"-"`            // 用户名称
	Remarks   string `json:"remarks,omitempty" validate:"-"`         // 备注
	TenantId  string `json:"tenantId,omitempty" validate:"required"` // 租户ID
}
