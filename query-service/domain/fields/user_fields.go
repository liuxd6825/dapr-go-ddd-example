package fields

type UserFields struct {
	Id        string `json:"id" validate:"gt=0" bson:"id"`
	TenantId  string `json:"tenantId" validate:"gt=0" bson:"tenantId"`
	UserCode  string `json:"userCode" validate:"gt=0" bson:"userCode"`
	UserName  string `json:"userName" validate:"gt=0" bson:"userName"`
	Email     string `json:"email" validate:"gt=0" bson:"email"`
	Telephone string `json:"telephone" validate:"gt=0" bson:"telephone"`
	Address   string `json:"address" validate:"gt=0" bson:"address"`
}
