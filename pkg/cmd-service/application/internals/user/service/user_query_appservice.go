package service

import (
	"context"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/cmd-service/application/internals"
	"github.com/liuxd6825/dapr-go-ddd-example/pkg/query-service/userinterface/rest/user/dto"
)

type UserQueryAppService interface {
	internals.QueryAppService[*dto.UserFindByIdResponse]
	FindById(ctx context.Context, tenantId, id string) (data *dto.UserFindByIdResponse, isFound bool, err error)
}

type userQueryAppService struct {
	internals.BaseQueryAppService[*dto.UserFindByIdResponse]
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
	res.SetAppId("query-service")
	res.SetResourceName("users")
	res.SetApiVersion("v1.0")
	return res
}

//
// FindById
// @Description: 按id获取用户信息
// @receiver s queryAppService
// @param ctx 上下文
// @param tenantId  租户id
// @param id 用户Id
// @return data 用户信息
// @return isFound 是否找到
// @return err 错误信息
//
func (s *userQueryAppService) FindById(ctx context.Context, tenantId, id string) (data *dto.UserFindByIdResponse, isFound bool, err error) {
	data = &dto.UserFindByIdResponse{}
	isFound, err = s.GetData(ctx, tenantId, id, data)
	return
}
