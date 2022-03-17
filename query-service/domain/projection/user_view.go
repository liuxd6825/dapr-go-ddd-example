package projection

type UserView struct {
	Id        string `json:"id" validate:"gt=0"`
	TenantId  string `json:"tenantId" validate:"gt=0"`
	UserCode  string `json:"code" validate:"gt=0"`
	UserName  string `json:"userName" validate:"gt=0"`
	Email     string `json:"email" validate:"gt=0"`
	Telephone string `json:"telephone" validate:"gt=0"`
	Address   string `json:"address" validate:"gt=0"`
}

func (u *UserView) GetTenantId() string {
	return u.TenantId
}
func (u *UserView) GetId() string {
	return u.Id
}
