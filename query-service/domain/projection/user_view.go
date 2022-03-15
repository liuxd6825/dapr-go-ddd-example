package projection

type UserView struct {
	Id       string `json:"id"`
	TenantId string `json:"tenantId"`
	Code     string `json:"code"`
	UserName string `json:"userName"`
}

func (u *UserView) GetTenantId() string {
	return u.TenantId
}
func (u *UserView) GetId() string {
	return u.Id
}
