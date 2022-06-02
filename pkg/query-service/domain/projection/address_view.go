package projection

type AddressView struct {
	TenantId string `json:"tenantId"  bson:"tenantId"  validate:"gt=0"`
	Id       string `json:"id"  bson:"_id"  validate:"gt=0"`
	UserId   string `json:"userId" bson:"userId" validate:"gt=0"`
	Province string `json:"province" bson:"province" validate:"gt=0"`
	City     string `json:"city" bson:"city" validate:"gt=0"`
	Area     string `json:"area" bson:"area" validate:"gt=0"`
	Location string `json:"location" bson:"location" validate:"gt=0"`
}

func (u *AddressView) GetTenantId() string {
	return u.TenantId
}
func (u *AddressView) GetId() string {
	return u.Id
}
