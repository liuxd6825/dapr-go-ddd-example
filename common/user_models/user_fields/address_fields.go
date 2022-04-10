package user_fields

type AddressFields struct {
	TenantId string `json:"tenantId" validate:"gt=0"`
	Id       string `json:"id" validate:"gt=0"`
	UserId   string `json:"userId" validate:"gt=0"`
	Province string `json:"province"  validate:"gt=0"`
	City     string `json:"city" validate:"gt=0"`
	Area     string `json:"area" validate:"gt=0"`
	Location string `json:"location" validate:"gt=0"`
}
