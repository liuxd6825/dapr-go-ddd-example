package fields

type UserFields struct {
	TenantId string `json:"tenantId" validate:"gt=0"`
	Id       string `json:"id" validate:"gt=0"`
	Code     string `json:"code" validate:"gt=0"`
	UserId   string `json:"userId" validate:"gt=0"`
	UserName string `json:"userName" validate:"gt=0"`
}
