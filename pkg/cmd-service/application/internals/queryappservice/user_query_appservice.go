package queryappservice

import "context"

type UserView struct {
	Id        string `json:"id" `
	TenantId  string `json:"tenantId" `
	UserCode  string `json:"userCode"`
	UserName  string `json:"userName"`
	Email     string `json:"email"`
	Telephone string `json:"telephone"`
	Address   string `json:"address"`
}

type UserQueryAppService interface {
	QueryAppService[*UserView]
	GetById(ctx context.Context, tenantId, id string) (data *UserView, isFound bool, err error)
}

type userQueryAppService struct {
	BaseQueryAppService[*UserView]
}

var _userQueryAppService UserQueryAppService

func init() {
	_userQueryAppService = newUserQueryAppService()
}

func GetUserQueryAppService() UserQueryAppService {
	return _userQueryAppService
}

func newUserQueryAppService() UserQueryAppService {
	res := &userQueryAppService{}
	res.appId = "query-service"
	res.resourceName = "users"
	res.apiVersion = "v1.0"
	return res
}

func (s *userQueryAppService) GetById(ctx context.Context, tenantId, id string) (data *UserView, isFound bool, err error) {
	data = &UserView{}
	isFound, err = s.GetData(ctx, tenantId, id, data)
	return
}
