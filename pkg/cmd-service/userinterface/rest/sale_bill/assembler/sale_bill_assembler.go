package assembler

import (
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/application/internals/sale_bill/appcmd"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/infrastructure/utils"
	"gitee.com/liuxu6825/dapr-ddd-demo/pkg/cmd-service/userinterface/rest/sale_bill/dto"
	"github.com/kataras/iris/v12"
)

type SaleBillAssembler struct {
}

//
// AssSaleBillCreateCommandDto
// @Description: 创建销售订单
// @receiver a
// @param ictx
// @return *appcmd.SaleBillCreateAppCmd 创建销售订单 应用层DTO对象
// @return error 错误
//
func (a *SaleBillAssembler) AssSaleBillCreateAppCmd(ictx iris.Context) (*appcmd.SaleBillCreateAppCmd, error) {
	var request dto.SaleBillCreateCommandRequest
	var appCmd appcmd.SaleBillCreateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssSaleBillUpdateCommandDto
// @Description: 更新销售订单
// @receiver a
// @param ictx
// @return *appcmd.SaleBillUpdateAppCmd 更新销售订单 应用层DTO对象
// @return error 错误
//
func (a *SaleBillAssembler) AssSaleBillUpdateAppCmd(ictx iris.Context) (*appcmd.SaleBillUpdateAppCmd, error) {
	var request dto.SaleBillUpdateCommandRequest
	var appCmd appcmd.SaleBillUpdateAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssSaleBillConfirmCommandDto
// @Description: 下单确认命令
// @receiver a
// @param ictx
// @return *appcmd.SaleBillConfirmAppCmd 下单确认命令 应用层DTO对象
// @return error 错误
//
func (a *SaleBillAssembler) AssSaleBillConfirmAppCmd(ictx iris.Context) (*appcmd.SaleBillConfirmAppCmd, error) {
	var request dto.SaleBillConfirmCommandRequest
	var appCmd appcmd.SaleBillConfirmAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}

//
// AssSaleBillDeleteCommandDto
// @Description: 删除销售订单
// @receiver a
// @param ictx
// @return *appcmd.SaleBillDeleteAppCmd 删除销售订单 应用层DTO对象
// @return error 错误
//
func (a *SaleBillAssembler) AssSaleBillDeleteAppCmd(ictx iris.Context) (*appcmd.SaleBillDeleteAppCmd, error) {
	var request dto.SaleBillDeleteCommandRequest
	var appCmd appcmd.SaleBillDeleteAppCmd
	if err := utils.AssemblerRequestBody(ictx, &request, &appCmd); err != nil {
		return nil, err
	}
	return &appCmd, nil
}
