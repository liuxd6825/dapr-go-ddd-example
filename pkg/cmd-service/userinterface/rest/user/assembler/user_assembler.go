package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/user/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/user/dto"
	"github.com/kataras/iris/v12"
)

type UserAssembler struct {
}

//
// AssUserDeleteAppCmd
// @Description: 删除用户
// @receiver a
// @param ictx
// @return *appcmd.UserDeleteAppCmd 删除用户 应用层DTO对象
// @return error 错误
//
func (a *UserAssembler) AssUserDeleteAppCmd(ictx iris.Context) (*appcmd.UserDeleteAppCmd, error) {
	var request dto.UserDeleteCommandRequest
	var appCmd appcmd.UserDeleteAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssUserCreateAppCmd
// @Description: 创建用户
// @receiver a
// @param ictx
// @return *appcmd.UserCreateAppCmd 创建用户 应用层DTO对象
// @return error 错误
//
func (a *UserAssembler) AssUserCreateAppCmd(ictx iris.Context) (*appcmd.UserCreateAppCmd, error) {
	var request dto.UserCreateCommandRequest
	var appCmd appcmd.UserCreateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssUserUpdateAppCmd
// @Description: 更新用户
// @receiver a
// @param ictx
// @return *appcmd.UserUpdateAppCmd 更新用户 应用层DTO对象
// @return error 错误
//
func (a *UserAssembler) AssUserUpdateAppCmd(ictx iris.Context) (*appcmd.UserUpdateAppCmd, error) {
	var request dto.UserUpdateCommandRequest
	var appCmd appcmd.UserUpdateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}
