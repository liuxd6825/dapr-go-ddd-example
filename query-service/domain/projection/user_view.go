package projection

type UserView struct {
	TenantId string
	Id       string
}

func (u *UserView) GetTenantId() string {
	return u.TenantId
}
func (u *UserView) GetId() string {
	return u.Id
}
