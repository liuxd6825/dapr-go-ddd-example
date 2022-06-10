package dto

// FindAggregateById

type FindAggregateByIdRequest struct {
	Id       string `json:"id" validate:"gt=0"`
	TenantId string `json:"tenantId" validate:"gt=0"`
}

type FindAggregateByIdResponse map[string]any

// Create

type UserCreateRequest struct {
	CommandId   string                `json:"commandId"`
	IsValidOnly bool                  `json:"isValidOnly"`
	Data        UserCreateRequestData `json:"data"`
}

type UserCreateRequestData struct {
	UserRequestData
}

type UserCreateResponse struct {
}

// Update

//
// UserUpdateRequest
// @Description:
//
type UserUpdateRequest struct {
	CommandId   string                `json:"commandId"`
	IsValidOnly bool                  `json:"isValidOnly"`
	UpdateMask  []string              `json:"updateMask"`
	Data        UserUpdateRequestData `json:"data"`
}

type UserUpdateRequestData struct {
	UserRequestData
}

type UserUpdateResponse struct {
}

// DeleteById

type UserDeleteByIdRequest struct {
	CommandId   string                    `json:"commandId"`
	IsValidOnly bool                      `json:"isValidOnly"`
	Data        UserDeleteByIdRequestData `json:"data"`
}

type UserDeleteByIdRequestData struct {
	Id       string `json:"id" validate:"gt=0"`
	TenantId string `json:"tenantId" validate:"gt=0"`
}

type UserDeleteByIdResponse struct {
}

// DeleteByIds

type UserDeleteByIdsRequest struct {
	CommandId string                     `json:"commandId"`
	Data      UserDeleteByIdsRequestData `json:"data"`
}

type UserDeleteByIdsRequestData struct {
	Ids      []string `json:"ids" validate:"gt=0"`
	TenantId string   `json:"tenantId" validate:"gt=0"`
}

type UserDeleteByIdsResponse struct {
}

// create update request data

type UserRequestData struct {
	Id        string `json:"id" validate:"gt=0" minLength:"16" maxLength:"16" example:"random string"`
	TenantId  string `json:"tenantId" validate:"gt=0" minLength:"16" maxLength:"16" example:"random string"`
	UserCode  string `json:"userCode" validate:"gt=0"  minLength:"16" maxLength:"16" example:"random string"`
	UserName  string `json:"userName" validate:"gt=0"  minLength:"16" maxLength:"16" example:"random string"`
	Email     string `json:"email" validate:"gt=0"    example:"xxx@163.com"`
	Telephone string `json:"telephone" validate:"gt=0"  length:"11" example:"18867766829"`
	Address   string `json:"address" validate:"gt=0"`
}
