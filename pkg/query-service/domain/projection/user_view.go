package projection

type UserView struct {
	Id        string `json:"id" bson:"_id" validate:"gt=0"`
	TenantId  string `json:"tenantId" bson:"tenantId" validate:"gt=0"`
	UserCode  string `json:"userCode" validate:"gt=0"  validate:"gt=0"`
	UserName  string `json:"userName" bson:"userName" validate:"gt=0"`
	Email     string `json:"email" bson:"email" validate:"gt=0"`
	Telephone string `json:"telephone" bson:"telephone" validate:"gt=0"`
	Address   string `json:"address" bson:"address" validate:"gt=0"`
}

func (u *UserView) GetTenantId() string {
	return u.TenantId
}
func (u *UserView) GetId() string {
	return u.Id
}
