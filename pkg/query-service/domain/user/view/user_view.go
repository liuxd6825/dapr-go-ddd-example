package view

import "time"

type UserView struct {
	Id         string     `json:"id" bson:"_id" `
	TenantId   string     `json:"tenantId" bson:"tenantId" `
	UserCode   string     `json:"userCode" bson:"userCode" `
	UserName   string     `json:"userName" bson:"userName" `
	Email      string     `json:"email" bson:"email" `
	Telephone  string     `json:"telephone" bson:"telephone" `
	Address    string     `json:"address" bson:"address" `
	Birthday   *time.Time `json:"birthday" bson:"birthday"`
	CreateTime *time.Time `json:"createTime" bson:"createTime"`
}

func (u *UserView) GetTenantId() string {
	return u.TenantId
}

func (u *UserView) GetId() string {
	return u.Id
}
